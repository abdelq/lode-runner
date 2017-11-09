package main

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/abdelq/lode-runner/game"
	. "github.com/abdelq/lode-runner/message"
)

type message Message

// TODO Remove
func newMessage(event, data string) *message {
	return &message{event, json.RawMessage(strconv.Quote(data))}
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
	joinMessage := new(JoinMessage)
	if err := joinMessage.Parse(data); err != nil {
		sender.out <- newMessage("error", err.Error())
		return
	}

	// Find/Create room
	room, ok := rooms[joinMessage.Room]
	if !ok {
		room = newRoom(joinMessage.Room)
	}

	room.join <- &join{sender,
		game.NewPlayer(joinMessage.Name, joinMessage.Role),
	}
}

// TODO Move to game package
func parseMove(data json.RawMessage, sender *client) {
	moveMessage := new(game.Message)
	if err := moveMessage.Parse(data); err != nil {
		sender.out <- newMessage("error", err.Error())
		return
	}

	// Find room name if none sent
	if moveMessage.Room == "" {
		moveMessage.Room = findRoom(sender)
		// TODO newMessage("error", "not in a room")
	}

	if room, ok := rooms[moveMessage.Room]; ok {
		if !room.game.Started() {
			sender.out <- newMessage("error", "game not yet started")
			return
		}

		if player := room.clients[sender]; player != nil {
			go player.Move(moveMessage.Direction, room.game.Level)
		} else {
			sender.out <- newMessage("error", "not a player")
		}
	}
}

// TODO Move to game package
func parseDig(data json.RawMessage, sender *client) {
	digMessage := new(game.Message)
	if err := digMessage.Parse(data); err != nil {
		sender.out <- newMessage("error", err.Error())
		return
	}

	// Find room name if none sent
	if digMessage.Room == "" {
		digMessage.Room = findRoom(sender)
		// TODO newMessage("error", "not in a room")
	}

	if room, ok := rooms[digMessage.Room]; ok {
		if !room.game.Started() {
			sender.out <- newMessage("error", "game not yet started")
			return
		}

		if runner, ok := room.clients[sender].(*game.Runner); ok {
			go runner.Dig(digMessage.Direction, room.game.Level)
		} else {
			sender.out <- newMessage("error", "not a runner")
		}
	}
}
