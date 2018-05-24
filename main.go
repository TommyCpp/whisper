package main

import (
	"github.com/tommycpp/Whisper/model"
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
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
