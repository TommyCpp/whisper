package model

import (
	"github.com/satori/go.uuid"
)
//todo: 处理控制请求(新连接，关闭连接 .etc)
type Server struct {
	userHandler   map[uuid.UUID]*WsHandler
	QueryHandlers chan HandlerQuery
}

func (server *Server) handle() {
	for {
		select {
		case query := <-server.QueryHandlers:
			{
				userIds := query.Receivers
				var result []chan *Message
				for _, userId := range userIds {
					result = append(result, server.userHandler[userId].MsgToSend)
				}
				query.Source.queryResult <- QueryResult{result, query.Msg}
			}
		}
	}
}
