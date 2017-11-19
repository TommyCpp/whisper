package Whisper

import (
	"github.com/satori/go.uuid"
	"time"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}





type Message struct {
	Content    string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
	Sender     User      `json:"sender"`
}

func NewMessage() *Message {
	return &Message{"", time.Now(), nil}
}

func (message *Message) timeFromNow() time.Duration {
	return time.Now().Sub(message.CreateTime)
}
