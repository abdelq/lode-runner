package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type message struct {
	Event string
	Data  json.RawMessage
}

func (m *message) parse(sender *client) {
	switch m.Event {
	case "join":
		var data struct{ Name, Room string }
		if err := json.Unmarshal(m.Data, &data); err != nil {
			log.Println(err)
			break
		}

		go sender.join(
			strings.TrimSpace(data.Name),
			strings.TrimSpace(data.Room),
		)
	default:
		err := fmt.Sprintf("%s is not a valid event", m.Event)
		log.Println(err)
		sender.out <- &message{"error", json.RawMessage(`"` + err + `"`)}
	}
}
