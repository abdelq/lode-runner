package main

import (
	"encoding/json"
	"log"
)

type Message struct {
	client *Client
	Event  string
	Data   json.RawMessage
}

func (m *Message) parse() {
	switch m.Event {
	case "join":
		data := &struct {
			Name string
			Room string
		}{}

		if err := json.Unmarshal(m.Data, data); err != nil {
			log.Println(err)
			break
		}

		go m.client.join(data.Name, data.Room)
	}
}
