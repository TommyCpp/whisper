package model

import (
	"github.com/satori/go.uuid"
	"net"
	"github.com/gorilla/websocket"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Addr     net.Addr
}

func NewUser(conn *websocket.Conn) *User {
	return &User{
		uuid.NewV4(),
		conn.LocalAddr().String(),
		conn.LocalAddr()}
}
