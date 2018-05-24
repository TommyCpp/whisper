package model

import (
	"fmt"
	"strconv"
)

//todo: 测试
type Server struct {
	UserHandlerMap      map[string]*WsHandler
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
				server.UserHandlerMap[strconv.Itoa(handler.Client.Id)] = handler // 添加到Id->Handler表中
				handler.Server = server                           //将Server指针添加到handler中
				fmt.Println(server.UserHandlerMap)
				go handler.handle()                                                                                 //启动处理线程
				go handler.read()                                                                                   //启动Read线程
				handler.MsgToSend <- &Message{"Your Id is " + strconv.Itoa(handler.Client.Id), "0", nil} //Return the id
			}
		case handler := <-server.CloseHandler:
			{
				handler.Close <- struct{}{} // send signal to close handler
				delete(server.UserHandlerMap, strconv.Itoa(handler.Client.Id))

			}

		}
	}
}
