package message

import (
	"fmt"
	"time"
)

type Message struct {
	Id   int64  `json:"id"`
	Text string `json:"text"`
}

func NewMessage(text string) *Message {
	return &Message{
		Id:   time.Now().UnixMilli(),
		Text: text,
	}
}

func (m Message) String() string {
	return fmt.Sprintf("Message (Id = %v, Text = %s)", m.Id, m.Text)
}
