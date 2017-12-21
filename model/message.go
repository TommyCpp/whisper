package model

import (
	"time"
	"github.com/satori/go.uuid"
)

type Message struct {
	Content     string      `json:"content"`
	CreateTime  time.Time   `json:"create_time"`
	SenderId    uuid.UUID   `json:"sender"`
	ReceiverIds []uuid.UUID `json:"receiver"`
}

func NewMessage() *Message {
	return &Message{"", time.Now(), nil, nil}
}

func (message *Message) TimeFromNow() time.Duration {
	return time.Now().Sub(message.CreateTime)
}
