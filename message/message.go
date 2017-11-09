package message

import (
	"encoding/json"
	"strconv"
)

type Message struct {
	Event string
	Data  json.RawMessage
}

func NewMessage(event, data string) *Message {
	return &Message{event, json.RawMessage(strconv.Quote(data))}
}
