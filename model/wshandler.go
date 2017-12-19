package model

import (
	"github.com/gorilla/websocket"
)

/*
WsHandler 负责一个客户端的收发工作，Server将保存一个map，将每一个User映射到一个WsHandler，如果需要向别的User发送单播消息，则在这个map中找到对应的WsHandler,使用send方法发送
*/
type WsHandler struct {
	Conn      websocket.Conn
	Client    User
	MsgToSend chan *Message
}

func NewWsHandler(conn websocket.Conn, client User) *WsHandler {
	return &WsHandler{
		conn,
		client,
		make(chan *Message),
	}
}

func (wsHandler *WsHandler) sendMsg(msg *Message) {
	wsHandler.MsgToSend <- msg
}

func (wsHandler *WsHandler) handle() {
	for {
		select {
		case msgToSend := <-wsHandler.MsgToSend:
			{
				wsHandler.Conn.WriteJSON(&msgToSend)
			}
		//	togo: add more handler func
		}
	}
}
