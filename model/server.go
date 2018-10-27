package model

import (
	"fmt"
	"strconv"
)

type Server struct {
	UserHandlerMap      map[string]*WsHandler
	QueryRedirectTarget chan HandlerQuery
	ConfigHandler       chan *IdAndHandlerConfig
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
					if handler, ok := server.UserHandlerMap[userId]; ok {
						result = append(result, handler.MsgToSend)
					}
				}
				query.Source.Redirect <- QueryResult{result, query.Msg}
			}
		case handler := <-server.CreateHandler: //接收到一个新的Conn
			{
				//handler := NewWsHandler(*conn, *NewUser(conn), make(chan HandlerQuery))
				server.UserHandlerMap[strconv.Itoa(handler.Client.Id)] = handler // 添加到Id->Handler表中
				handler.Server = server                                          //将Server指针添加到handler中
				fmt.Println(server.UserHandlerMap)
				go handler.handle()                                                                      //启动处理线程
				handler.MsgToSend <- &Message{"Your Id is " + strconv.Itoa(handler.Client.Id), "0", nil} //Return the id
			}
		case handler := <-server.CloseHandler:
			{
				handler.Conn.Close()
				delete(server.UserHandlerMap, strconv.Itoa(handler.Client.Id))
				close(handler.MsgToSend)
				close(handler.MsgReceived)
				close(handler.Redirect)
				close(handler.Close)
			}
		case idAndConfig := <-server.ConfigHandler:
			{
				id := idAndConfig.Id
				config := idAndConfig.Config
				if handler, ok := server.UserHandlerMap[id]; ok {
					handler.ConfigHandler <- config
				}

			}

		}
	}
}
