package main

import (
	"github.com/tommycpp/Whisper/model"
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
	"encoding/json"
)

var server = model.Server{
	UserHandlerMap:      make(map[string]*model.WsHandler),
	QueryRedirectTarget: make(chan model.HandlerQuery),
	CreateHandler:       make(chan *model.WsHandler),
	CloseHandler:        make(chan *model.WsHandler),
}

func main() {
	start(&server)
}

func start(server *model.Server) {
	fmt.Println("Start processing....")
	go server.Handle()
	http.HandleFunc("/login", loginHandler)
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
	wsHandler := model.NewWsHandler(*conn, *model.NewUser(conn))
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
	//todo: Token生成算法
	return []byte(username)
}
