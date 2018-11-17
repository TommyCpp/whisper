package main

import (
	"crypto/md5"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/websocket"
	"github.com/tommycpp/Whisper/config"
	"github.com/tommycpp/Whisper/model"
	"github.com/tommycpp/Whisper/sqlconnection"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
)

var server = *model.NewServer()

var configuration = config.NewConfiguration()
var db *sqlconnection.SqlConnection
var userIndex uint32
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func main() {
	start(&server)
}

func start(server *model.Server) {
	if err := config.ReadConfig("./config/config.json", configuration); err != nil {
		log.Fatal("Cannot read config file")
		log.Fatal(err)
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
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		cookie := make(map[string]string)
		rawCookie, err := request.Cookie("cookie")
		if err != nil {
			log.Println("User does not login")
			http.Redirect(writer, request, "/login", 302)
			return
		}
		if err := cookieHandler.Decode("cookie", rawCookie.Value, &cookie); err != nil {
			log.Println("User does not login")
			http.Redirect(writer, request, "/login", 302)
			return
		}
		http.ServeFile(writer, request, "./static/client.html")
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			http.ServeFile(writer, request, "./static/register.html")
		case "POST":
			registerHandler(writer, request)
		}
	})
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			http.ServeFile(writer, request, "./static/login.html")
		case "POST":
			loginHandler(writer, request)
		}
	})
	http.HandleFunc("/message", handler)
	http.HandleFunc("/config", func(writer http.ResponseWriter, request *http.Request) {
		if request.Header.Get("Whisper-Config") != "" {
			//if it is a config request
			id := request.Header.Get("Handler-Id")
			if id != "" {
				handlerConfig, middlewareName, err := GetHandlerConfig(request)
				if err != nil {
					log.Print("Error when process the config request")
					log.Print(err)
				} else {
					switch middlewareName {
					case "RSA":
						{
							idAndConfig := &model.IdAndHandlerConfig{
								Id:     id,
								Config: handlerConfig,
							}
							server.ConfigHandler <- idAndConfig
							publicKey := handlerConfig.MiddleWare.(*model.RSAEncryptionMiddleware).Cipher.(*model.RSACipher).KeyPair.PublicKey
							derPkix := x509.MarshalPKCS1PublicKey(publicKey)
							block := &pem.Block{
								Type:  "PUBLIC KEY",
								Bytes: derPkix,
							}
							err = pem.Encode(writer, block)
							if err != nil {
								log.Println("Error when encode pem")
							}
							break
						}
					case "E2E":
						{
							idAndConfig := &model.IdAndHandlerConfig{
								Id:     id,
								Config: handlerConfig,
							}
							server.ConfigHandler <- idAndConfig
							break
						}
					}

				}
			}
			//send
		}
	})
	http.ListenAndServe("localhost:8086", nil)
}

func GetHandlerConfig(request *http.Request) (*model.HandlerConfig, string, error) {
	var handlerConfig = new(model.HandlerConfig)
	var handlerConfigString = new(struct {
		Op             string                      `json:"op"`
		MiddlewareName string                      `json:"middleware_name"`
		Settings       map[string]*json.RawMessage `json:"setting"`
	})
	if err := json.NewDecoder(request.Body).Decode(handlerConfigString); err == nil {
		handlerConfig.Op = handlerConfigString.Op
		switch handlerConfigString.MiddlewareName {
		case "RSA":
			{
				// Add a RSA Middleware
				var publicKey string
				err = json.Unmarshal(*handlerConfigString.Settings["public_key"], &publicKey)
				if err != nil {
					log.Println("Do not have public_key")
				} else {
					handlerConfig.MiddleWare = model.NewRSAEncryptionMiddleware(model.NewRSACipher([]byte(publicKey)))
					return handlerConfig, "RSA", nil
				}
				break
			}
		case "E2E":
			{
				var publicKey string
				var targetId string
				err = json.Unmarshal(*handlerConfigString.Settings["public_key"], &publicKey)
				if err != nil {
					log.Println("Do not have public_key")
				} else {
					err = json.Unmarshal(*handlerConfigString.Settings["target"], &targetId)
					if err != nil {
						log.Println("Cannot get target")
					}
					handlerConfig.MiddleWare = model.NewE2eEncryptionMiddleware([]byte(targetId), []byte(publicKey))
					return handlerConfig, "E2E", nil
				}
				break
			}
		}
	} else {
		return nil, "", err
	}
	return nil, "", nil
}

func handler(res http.ResponseWriter, req *http.Request) {
	cookie := make(map[string]string)
	rawCookie, err := req.Cookie("cookie")
	if err != nil {
		log.Println("Cannot get Cookie")
		return
	}
	if err := cookieHandler.Decode("cookie", rawCookie.Value, &cookie); err != nil {
		log.Println("Cannot parse cookie")
	}

	id, err := strconv.Atoi(cookie["id"])
	if err != nil {
		log.Println("Cannot get id from cookie")
		return
	}
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		http.NotFound(res, req)
		return
	}
	fmt.Println("Open an WebSocket channel")
	wsHandler := model.NewWsHandler(*conn, *model.NewUser(conn, id), configuration)
	server.CreateHandler <- wsHandler
}

func registerHandler(res http.ResponseWriter, req *http.Request) {
	var account model.Account
	err := json.NewDecoder(req.Body).Decode(&account)
	account.Id = int(atomic.AddUint32(&userIndex, 1)) // get next id
	if err != nil {
		http.Error(res, "Cannot create user", http.StatusBadRequest)
		return
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
		//if login
		setCookie(account.Id, res)
		fmt.Println("User " + account.Username + " has logged in")
		http.Redirect(res, req, "/", 302)
	} else {
		res.WriteHeader(401)
		res.Write([]byte("Cannot authorized"))
	}
}

func generateToken(username string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(username))
	return []byte(hex.EncodeToString(hasher.Sum(nil)))
}

func setCookie(id int, response http.ResponseWriter) {
	value := map[string]string{
		"id": strconv.Itoa(id),
	}
	if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}
