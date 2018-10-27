package model

import (
	"github.com/gorilla/websocket"
	"net"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Addr     net.Addr
}

func NewUser(conn *websocket.Conn, Id int) *User {
	return &User{
		Id,
		conn.LocalAddr().String(),
		conn.LocalAddr()}
}
