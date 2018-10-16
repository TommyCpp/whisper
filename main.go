package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/tommycpp/Whisper/config"
	"github.com/tommycpp/Whisper/model"
	"net/http"
)

var server = model.Server{
	UserHandlerMap:      make(map[string]*model.WsHandler),
	QueryRedirectTarget: make(chan model.HandlerQuery),
	CreateHandler:       make(chan *model.WsHandler),
	CloseHandler:        make(chan *model.WsHandler),
}

var configuration = config.Config

func main() {
	start(&server)
}

func start(server *model.Server) {
	err := config.ReadConfig("./config/config.json", configuration)
	sql := GetSqlConnection() //get database connection
	defer sql.Close()         //close database connection
	if err == nil {
		fmt.Println("Start processing....")
		go server.Handle()
		http.HandleFunc("/login", loginHandler)
		http.HandleFunc("/", handler)
		http.ListenAndServe("localhost:8086", nil)
	} else {
		fmt.Println("Cannot read config file.")
		fmt.Println(err)
	}

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

func loginHandler(res http.ResponseWriter, req *http.Request) {
	//todo: 测试
	var account model.Account
	err := json.NewDecoder(req.Body).Decode(&account) // read User
	if err != nil {
		http.Error(res, "Cannot authentication", http.StatusUnauthorized)
		return
	}
	fmt.Println("User " + account.Username + " has logged in")
	json.NewEncoder(res).Encode(struct {
		Token []byte `json:"token"`
	}{generateToken(account.Username)})
	return
}

func generateToken(username string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(username))
	return []byte(hex.EncodeToString(hasher.Sum(nil)))
}
