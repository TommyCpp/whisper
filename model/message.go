package model

import "time"

type Message struct {
	Content    string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
	Sender     User      `json:"sender"`
	Receiver   []User    `json:"receiver"`
}

func NewMessage() *Message {
	return &Message{"", time.Now(), nil,nil}
}

func (message *Message) TimeFromNow() time.Duration {
	return time.Now().Sub(message.CreateTime)
}

