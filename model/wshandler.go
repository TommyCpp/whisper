package model

import (
	"github.com/gorilla/websocket"
	"fmt"
)

/*
WsHandler 负责一个客户端的收发工作，Server将保存一个map，将每一个User映射到一个WsHandler，如果需要向别的User发送单播消息，则在这个map中找到对应的WsHandler,使用send方法发送
*/
type WsHandler struct {
	Conn        websocket.Conn
	Client      User
	MsgToSend   chan *Message
	MsgReceived chan *Message
	Redirect    chan QueryResult
	Close       chan struct{}
	Server      *Server
}

//用于在broadcast中查询接受者的handler,并取出其中的MsgToSend
type HandlerQuery struct {
	Receivers []string
	Source    *WsHandler
	Msg       *Message
}

type QueryResult struct {
	handlerChans []chan *Message //对应接受方的MsgToSend channel
	Msg          *Message
}

func NewWsHandler(conn websocket.Conn, client User) *WsHandler {
	return &WsHandler{
		conn,
		client,
		make(chan *Message),
		make(chan *Message),
		make(chan QueryResult),
		make(chan struct{}),
		nil,
	}
}

func (wsHandler *WsHandler) sendMsg(msg *Message) {
	wsHandler.MsgToSend <- msg
}

func (wsHandler *WsHandler) redirectMsg(handlerChan chan *Message, message *Message) {
	handlerChan <- message
}

func (wsHandler *WsHandler) handle() {
	go wsHandler.read() // 启动Read线程
	for {
		select {
		case msgToSend := <-wsHandler.MsgToSend:
			{
				wsHandler.Conn.WriteJSON(struct {
					Content string
					Sender  string
				}{(msgToSend).Content, (msgToSend).SenderId})
			}
			//	togo: add more handler func
		case msgReceived := <-wsHandler.MsgReceived:
			{
				receiverIds := msgReceived.ReceiverIds
				wsHandler.Server.QueryRedirectTarget <-
					HandlerQuery{
						receiverIds, wsHandler, msgReceived,
					}
				//传入Server的QueryRedirectTarget channel
			}
		case queryResult := <-wsHandler.Redirect: //转发消息
			{
				msg := queryResult.Msg
				for _, handlerChan := range queryResult.handlerChans {
					go wsHandler.redirectMsg(handlerChan, msg)
				}
			}
		case _ = <-wsHandler.Close:
			{
				return
			}
		}
	}
}

func (wsHandler *WsHandler) read() {
	for {
		var message Message
		err := wsHandler.Conn.ReadJSON(&message)
		if err != nil {
			fmt.Println("read:", err)
			wsHandler.close()
			return
		}
		wsHandler.MsgReceived <- &message
	}
}

func (wsHandler *WsHandler) close() {
	wsHandler.Close <- struct{}{}
	wsHandler.Server.CloseHandler <- wsHandler
}
