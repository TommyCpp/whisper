package model

import (
	"github.com/satori/go.uuid"
	"github.com/gorilla/websocket"
)

//todo: 处理控制请求(关闭连接 .etc)
type Server struct {
	userHandlerMap      map[uuid.UUID]*WsHandler
	QueryRedirectTarget chan HandlerQuery
	CreateHandler       chan *websocket.Conn
}

func (server *Server) handle() {
	for {
		select {
		case query := <-server.QueryRedirectTarget:
			{
				userIds := query.Receivers //接受方ID
				var result []chan *Message
				for _, userId := range userIds {
					result = append(result, server.userHandlerMap[userId].MsgToSend)
				}
				query.Source.Redirect <- QueryResult{result, query.Msg}
			}
		case conn := <-server.CreateHandler: //接收到一个新的Conn
			{
				handler := NewWsHandler(*conn, *NewUser(conn), make(chan HandlerQuery))
				server.userHandlerMap[handler.Client.Id] = handler // 添加到Id->Handler表中
				go handler.handle() //启动处理线程
			}

		}
	}
}
