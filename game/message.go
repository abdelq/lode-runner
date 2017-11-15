package game

import (
	"encoding/json"
	"errors"
	"strings"
)

// TODO Move to msg package?
type Message struct {
	Direction uint8
	Room      string
}

func (m *Message) Parse(data json.RawMessage) error {
	if err := json.Unmarshal(data, m); err != nil {
		return err
	}

	// TODO Valid only LEFT/RIGHT for digging
	if m.Direction > 4 { // NONE is a valid direction
		return errors.New("invalid direction")
	}
	m.Room = strings.TrimSpace(m.Room)

	return nil
}
