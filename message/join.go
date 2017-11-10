package message

import (
	"encoding/json"
	"errors"
	"strings"
)

type JoinMessage struct {
	Name, Room string
	Role       uint8
}

func (m *JoinMessage) Parse(data json.RawMessage) error {
	if err := json.Unmarshal(data, m); err != nil {
		return err
	}

	// TODO Remove hardcoded value
	if m.Role < 2 { // Runner/Guard
		if m.Name = strings.TrimSpace(m.Name); m.Name == "" {
			return errors.New("invalid name")
		}
	}
	if m.Room = strings.TrimSpace(m.Room); m.Room == "" {
		return errors.New("invalid room")
	}

	return nil
}
