package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/abdelq/lode-runner/game"
)

type message struct {
	Event string
	Data  json.RawMessage
}

func newErrorMessage(err string) *message {
	return &message{"error", json.RawMessage(strconv.Quote(err))}
}

func newJoinMessage(name string, role uint8) *message {
	return &message{"join", json.RawMessage(fmt.Sprintf(
		`{"name": %q, "role": %d}`, name, role,
	))}
}

func newLeaveMessage(name string, role uint8) *message {
	return &message{"leave", json.RawMessage(fmt.Sprintf(
		`{"name": %q, "role": %d}`, name, role,
	))}
}

func (m *message) parse(sender *client) {
	m.Event = strings.ToLower(strings.TrimSpace(m.Event)) // TODO
	switch m.Event {
	case "join":
		go parseJoin(m.Data, sender)
	case "move":
		go parseMove(m.Data, sender)
	case "dig":
		go parseDig(m.Data, sender)
	default:
		sender.out <- newErrorMessage("invalid event")
	}
}

type joinMessage struct {
	Name, Room string
	Role       uint8
}

func (m *joinMessage) parse(data json.RawMessage) error {
	if err := json.Unmarshal(data, m); err != nil {
		return err
	}

	m.Name = strings.TrimSpace(m.Name)
	if m.Name == "" {
		return errors.New("invalid name")
	}

	m.Room = strings.TrimSpace(m.Room)
	if m.Room == "" {
		return errors.New("invalid room")
	}

	return nil
}

func parseJoin(data json.RawMessage, sender *client) {
	var joinMessage *joinMessage
	if err := joinMessage.parse(data); err != nil {
		sender.out <- newErrorMessage(err.Error())
		return
	}

	room, ok := rooms[joinMessage.Room]
	if !ok {
		room = newRoom(joinMessage.Room)
	}

	switch joinMessage.Role {
	case 0: // Runner
		room.join <- &join{sender, &game.Runner{Name: joinMessage.Name}}
	case 1: // Guard
		room.join <- &join{sender, &game.Guard{Name: joinMessage.Name}}
	default: // Spectator
		room.join <- &join{sender, nil}
	}
}

func parseMove(data json.RawMessage, sender *client) {
	var moveMessage *game.Message
	if err := moveMessage.Parse(data); err != nil {
		sender.out <- newErrorMessage(err.Error())
		return
	}

	// TODO Find room with client if none sent
	if room, ok := rooms[moveMessage.Room]; ok {
		if player := room.clients[sender]; player != nil {
			go player.Move(moveMessage.Direction, room.game)
		} else {
			sender.out <- newErrorMessage("not a player")
		}
	}
}

func parseDig(data json.RawMessage, sender *client) {
	var digMessage *game.Message
	if err := digMessage.Parse(data); err != nil {
		sender.out <- newErrorMessage(err.Error())
		return
	}

	// TODO Find room with client if none sent
	if room, ok := rooms[digMessage.Room]; ok {
		if runner, ok := room.clients[sender].(*game.Runner); ok {
			go runner.Dig(digMessage.Direction, room.game)
		} else {
			sender.out <- newErrorMessage("not a runner")
		}
	}
}
