package model

import (
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

/*
WsHandler 负责一个客户端的收发工作，Server将保存一个map，将每一个User映射到一个WsHandler，如果需要向别的User发送单播消息，则在这个map中找到对应的WsHandler,使用send方法发送
*/
type WsHandler struct {
	Conn          websocket.Conn
	Client        User
	MsgToSend     chan *Message
	MsgReceived   chan *Message
	queryHandlers chan HandlerQuery
	queryResult   chan QueryResult
}

type HandlerQuery struct {
	Receivers []uuid.UUID
	Source    *WsHandler
	Msg       *Message
}

type QueryResult struct {
	handlerChans []chan *Message
	Msg          *Message
}

func NewWsHandler(conn websocket.Conn, client User, queryHandler chan HandlerQuery) *WsHandler {
	return &WsHandler{
		conn,
		client,
		make(chan *Message),
		make(chan *Message),
		queryHandler,
		make(chan QueryResult),
	}
}

func (wsHandler *WsHandler) sendMsg(msg *Message) {
	wsHandler.MsgToSend <- msg
}

func (wsHandler *WsHandler) redirectMsg(handlerChan chan *Message, message *Message) {
	handlerChan <- message
}

func (wsHandler *WsHandler) handle() {
	for {
		select {
		case msgToSend := <-wsHandler.MsgToSend:
			{
				wsHandler.Conn.WriteJSON(&msgToSend)
			}
			//	togo: add more handler func
		case msgReceived := <-wsHandler.MsgReceived:
			{
				receiverIds := msgReceived.ReceiverIds
				wsHandler.queryHandlers <-
					HandlerQuery{
						receiverIds, wsHandler, msgReceived,
					}
			}
		case queryResult := <-wsHandler.queryResult:
			{
				msg := queryResult.Msg
				for _, handlerChan := range queryResult.handlerChans {
					go wsHandler.redirectMsg(handlerChan, msg)
				}
			}
		}
	}
}
