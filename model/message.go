package model

type Message struct {
	Content     string   `json:"content"`
	SenderId    string   `json:"sender"`
	ReceiverIds []string `json:"receiver"`
}

//func NewMessage() *Message {
//	return &Message{"", time.Now(), nil, nil}
//}
