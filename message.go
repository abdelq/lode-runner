package main

import (
	"encoding/json"
	"strings"

	"github.com/abdelq/lode-runner/game"
	. "github.com/abdelq/lode-runner/message"
)

type message Message

func newMessage(event, data string) *message { // FIXME
	msg := message(*NewMessage(event, data))
	return &msg
}

// TODO Move sections to game package
func (m *message) parse(sender *client) {
	switch evt := strings.ToLower(strings.TrimSpace(m.Event)); evt {
	case "join":
		go parseJoin(m.Data, sender)
	case "move":
		go parseMove(m.Data, sender)
	case "dig":
		go parseDig(m.Data, sender)
	default:
		sender.out <- newMessage("error", "invalid event")
	}
}

func parseJoin(data json.RawMessage, sender *client) {
	msg := new(JoinMessage)
	if err := msg.Parse(data); err != nil {
		sender.out <- newMessage("error", err.Error())
		return
	}

	// Find/Create room
	room, ok := rooms[msg.Room]
	if !ok {
		room = newRoom(msg.Room)
	}

	room.join <- &join{sender, game.NewPlayer(msg.Name, msg.Role)}
}

// TODO Move to game package
func parseMove(data json.RawMessage, sender *client) {
	msg := new(game.Message)
	if err := msg.Parse(data); err != nil {
		sender.out <- newMessage("error", err.Error())
		return
	}

	// Find room name if none sent
	if msg.Room == "" {
		if msg.Room = findRoom(sender); msg.Room == "" {
			sender.out <- newMessage("error", "not in a room")
			return
		}
	}

	if room, ok := rooms[msg.Room]; ok {
		if !room.game.Started() {
			sender.out <- newMessage("error", "game not yet started")
			return
		}

		if player := room.clients[sender]; player != nil {
			go player.Move(msg.Direction, room.game.Level)
		} else {
			sender.out <- newMessage("error", "not a player")
		}
	}
}

// TODO Move to game package
func parseDig(data json.RawMessage, sender *client) {
	msg := new(game.Message)
	if err := msg.Parse(data); err != nil {
		sender.out <- newMessage("error", err.Error())
		return
	}

	// Find room name if none sent
	if msg.Room == "" {
		if msg.Room = findRoom(sender); msg.Room == "" {
			sender.out <- newMessage("error", "not in a room")
			return
		}
	}

	if room, ok := rooms[msg.Room]; ok {
		if !room.game.Started() {
			sender.out <- newMessage("error", "game not yet started")
			return
		}

		if runner, ok := room.clients[sender].(*game.Runner); ok {
			go runner.Dig(msg.Direction, room.game.Level)
		} else {
			sender.out <- newMessage("error", "not a runner")
		}
	}
}
