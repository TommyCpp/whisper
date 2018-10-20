package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github.com/tommycpp/Whisper/config"
	"github.com/tommycpp/Whisper/model"
	"github.com/tommycpp/Whisper/sqlconnection"
	"log"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
)

var server = model.Server{
	UserHandlerMap:      make(map[string]*model.WsHandler),
	QueryRedirectTarget: make(chan model.HandlerQuery),
	CreateHandler:       make(chan *model.WsHandler),
	CloseHandler:        make(chan *model.WsHandler),
}

var configuration = config.NewConfiguration()
var db *sqlconnection.SqlConnection
var userIndex uint32

func main() {
	start(&server)
}

func start(server *model.Server) {
	if err := config.ReadConfig("./config/config.json", configuration); err != nil {
		log.Fatal("Cannot read config file")
		log.Fatal(err)
		os.Exit(1)
	}
	db = sqlconnection.GetSqlConnection(configuration) //get database connection
	//get next user id
	rows, err := db.Query("SELECT MAX(Id) FROM user")
	defer rows.Close()
	if err != nil {
		log.Fatal("Cannot get last id")
		log.Fatal(err)
	} else {
		for rows.Next() {
			if err := rows.Scan(&userIndex); err != nil {
				if strings.Contains(err.Error(), `Scan error on column index 0`) {
					//if the database is empty
					userIndex = 0
				} else {
					log.Fatal(err)
				}
			}
		}
	}
	fmt.Println("Start processing....")
	go server.Handle()
	http.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			http.ServeFile(writer, request, "./frontend/register.html")
		case "POST":
			registerHandler(writer, request)
		}
	})
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			http.ServeFile(writer, request, "./frontend/login.html")
		case "POST":
			loginHandler(writer, request)
		}
	})
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8086", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		http.NotFound(res, req)
		return
	}
	fmt.Println("Open an WebSocket channel")
	wsHandler := model.NewWsHandler(*conn, *model.NewUser(conn), configuration)
	server.CreateHandler <- wsHandler
}

func registerHandler(res http.ResponseWriter, req *http.Request) {
	var account model.Account
	err := json.NewDecoder(req.Body).Decode(&account)
	account.Id = int(atomic.AddUint32(&userIndex, 1)) // get next id
	if err != nil {
		http.Error(res, "Cannot create user", http.StatusBadRequest)
	}
	fmt.Println("Creating " + account.Username)
	storeResult, err := account.StoreIntoDB(db)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			// when there is a duplicated ID or duplicated username
			http.Error(res, "Username has already been taken", http.StatusBadRequest)
		} else {
			http.Error(res, "Error when inserting into DB", http.StatusBadRequest)
		}
	} else {
		_, err := storeResult.LastInsertId()
		if err != nil {
			http.Error(res, "Error when connection to DB", http.StatusInternalServerError)
		} else {
			if err := json.NewEncoder(res).Encode(account); err != nil {
				http.Error(res, "Error when encode response", http.StatusInternalServerError)
			}
		}
	}
}

func loginHandler(res http.ResponseWriter, req *http.Request) {
	var account model.Account
	err := json.NewDecoder(req.Body).Decode(&account) // read User
	if err != nil {
		http.Error(res, "Cannot authentication", http.StatusUnauthorized)
		return
	}
	if ifValid, err := account.CheckIfValid(db); err == nil && ifValid {
		json.NewEncoder(res).Encode(struct {
			Token []byte `json:"token"`
		}{generateToken(account.Username)})
		fmt.Println("User " + account.Username + " has logged in")
	}

	return
}

func generateToken(username string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(username))
	return []byte(hex.EncodeToString(hasher.Sum(nil)))
}
