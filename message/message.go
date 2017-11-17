package message

import (
	"encoding/json"
	"strconv"
)

type Message struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

func NewMessage(event, data string) *Message {
	return &Message{event, json.RawMessage(strconv.Quote(data))}
}
