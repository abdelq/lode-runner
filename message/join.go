package message

import (
	"encoding/json"
	"errors"
	"strings"
)

type JoinMessage struct {
	Name, Room string
	Role       byte
}

func (m *JoinMessage) Parse(data json.RawMessage) error {
	if err := json.Unmarshal(data, m); err != nil {
		return err
	}

	// TODO Replace 0, '&', '0'
	if m.Role == 0 || m.Role == '&' || m.Role == '0' { // Players only
		if m.Name = strings.TrimSpace(m.Name); m.Name == "" {
			return errors.New("invalid name")
		}
	}
	if m.Room = strings.TrimSpace(m.Room); m.Room == "" {
		return errors.New("invalid room")
	}

	return nil
}
