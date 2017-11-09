package main

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/abdelq/lode-runner/game"
	. "github.com/abdelq/lode-runner/message"
)

type message Message

func newMessage(event, data string) *message {
	return &message{event, json.RawMessage(strconv.Quote(data))}
}

// TODO Move part to game package
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
		sender.out <- newMessage("error", "invalid event")
	}
}

func parseJoin(data json.RawMessage, sender *client) {
	joinMessage := new(JoinMessage)
	if err := joinMessage.Parse(data); err != nil {
		sender.out <- newMessage("error", err.Error())
		return
	}

	room, ok := rooms[joinMessage.Room]
	if !ok {
		room = newRoom(joinMessage.Room)
	}

	// TODO Move to game package
	switch joinMessage.Role {
	case 0: // Runner
		room.join <- &join{sender, &game.Runner{Name: joinMessage.Name}}
	case 1: // Guard
		room.join <- &join{sender, &game.Guard{Name: joinMessage.Name}}
	default: // Spectator
		room.join <- &join{sender, nil}
	}
}

// TODO Move to game package
func parseMove(data json.RawMessage, sender *client) {
	moveMessage := new(game.Message)
	if err := moveMessage.Parse(data); err != nil {
		sender.out <- newMessage("error", err.Error())
		return
	}

	// TODO Find room with client if none sent
	if room, ok := rooms[moveMessage.Room]; ok {
		if !room.game.Started() {
			sender.out <- newMessage("error", "game not yet started")
			return
		}

		if player := room.clients[sender]; player != nil {
			go player.Move(room.game.Level, moveMessage.Direction)
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

	// TODO Find room with client if none sent
	if room, ok := rooms[digMessage.Room]; ok {
		if !room.game.Started() {
			sender.out <- newMessage("error", "game not yet started")
			return
		}

		if runner, ok := room.clients[sender].(*game.Runner); ok {
			go runner.Dig(room.game.Level, digMessage.Direction)
		} else {
			sender.out <- newMessage("error", "not a runner")
		}
	}
}
