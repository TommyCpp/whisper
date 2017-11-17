package model

import "time"

type Message struct {
	Content   string    `json:"content"`
	TimeStamp time.Time `json:"time_stamp"`
	Sender    User      `json:"sender"`
}

func (message *Message) timeFromNow() time.Duration {
	return time.Now().Sub(message.TimeStamp)
}
