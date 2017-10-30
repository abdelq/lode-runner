package game

import (
	"encoding/json"
	"errors"
	"strings"
)

type Message struct {
	Direction uint8
	Room      string
}

func (m *Message) Parse(data json.RawMessage) error {
	if err := json.Unmarshal(data, m); err != nil {
		return err
	}

	if m.Direction > 3 {
		return errors.New("invalid direction")
	}

	m.Room = strings.TrimSpace(m.Room)
	if m.Room == "" { // TODO Temporary
		return errors.New("invalid room")
	}

	return nil
}
