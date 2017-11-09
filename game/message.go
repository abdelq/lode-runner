package game

import (
	"encoding/json"
	"errors"
	"strings"
)

type Message struct {
	Direction direction
	Room      string
}

func (m *Message) Parse(data json.RawMessage) error {
	if err := json.Unmarshal(data, m); err != nil {
		return err
	}

	if m.Direction > RIGHT { // TODO Comment
		return errors.New("invalid direction")
	}
	if m.Room = strings.TrimSpace(m.Room); m.Room == "" { // TODO Temporary
		return errors.New("invalid room")
	}

	return nil
}
