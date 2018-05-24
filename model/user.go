package model

import (
	"net"
	"github.com/gorilla/websocket"
)

var maxId = 0

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Addr     net.Addr
}

func NewUser(conn *websocket.Conn) *User {
	maxId += 1
	return &User{
		maxId,
		conn.LocalAddr().String(),
		conn.LocalAddr()}
}
