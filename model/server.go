package model

import (
	"github.com/satori/go.uuid"
)

//todo: 测试
type Server struct {
	UserHandlerMap      map[uuid.UUID]*WsHandler
	QueryRedirectTarget chan HandlerQuery
	CreateHandler       chan *WsHandler
	CloseHandler        chan *WsHandler
}

func (server *Server) Handle() {
	for {
		select {
		case query := <-server.QueryRedirectTarget:
			{
				userIds := query.Receivers //接受方ID
				var result []chan *Message
				for _, userId := range userIds {
					result = append(result, server.UserHandlerMap[userId].MsgToSend)
				}
				query.Source.Redirect <- QueryResult{result, query.Msg}
			}
		case handler := <-server.CreateHandler: //接收到一个新的Conn
			{
				//handler := NewWsHandler(*conn, *NewUser(conn), make(chan HandlerQuery))
				server.UserHandlerMap[handler.Client.Id] = handler // 添加到Id->Handler表中
				go handler.handle()                                //启动处理线程
			}
		case handler := <-server.CloseHandler:
			{
				handler.Close <- struct{}{} // send signal to close handler
				delete(server.UserHandlerMap, handler.Client.Id)

			}

		}
	}
}
