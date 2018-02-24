package main

import (
	"encoding/json"
	"strings"

	"github.com/abdelq/lode-runner/game"
	msg "github.com/abdelq/lode-runner/message"
)

type message msg.Message

// TODO Move sections to game package
func (m *message) parse(sender *client) {
	switch evt := strings.ToLower(strings.TrimSpace(m.Event)); evt {
	case "join":
		parseJoin(m.Data, sender)
	case "move":
		parseMove(m.Data, sender)
	case "dig":
		parseDig(m.Data, sender)
	case "list":
		names := make([]string, 0, 16) // XXX
		rooms.Range(func(name, room interface{}) bool {
			names = append(names, name.(string))
			return true
		})
		roomNames, _ := json.Marshal(names)
		sender.out <- &msg.Message{"list", roomNames}
	case "kill":
		var roomName string
		if err := json.Unmarshal(m.Data, &roomName); err != nil {
			return
		}

		if r, ok := rooms.Load(roomName); ok {
			if r, ok := r.(*room); ok {
				r.game.Kill()
			}
		}
	default:
		sender.out <- msg.NewMessage("error", "invalid event")
	}
}

func parseJoin(data json.RawMessage, sender *client) {
	message := new(msg.JoinMessage)
	if err := message.Parse(data); err != nil {
		sender.out <- msg.NewMessage("error", err.Error())
		return
	}

	// Find/Create room
	r, _ := rooms.LoadOrStore(message.Room, newRoom(message.Room))

	r.(*room).join <- &join{sender, game.NewPlayer(message, sender.out)}
}

// TODO Move to game package
func parseMove(data json.RawMessage, sender *client) {
	message := new(msg.GameMessage)
	if err := message.Parse(data); err != nil {
		sender.out <- msg.NewMessage("error", err.Error())
		return
	}

	// Find room name if none sent
	if message.Room == "" {
		if message.Room = findRoom(sender); message.Room == "" {
			sender.out <- msg.NewMessage("error", "not in a room")
			return
		}
	}

	if r, ok := rooms.Load(message.Room); ok {
		if r, ok := r.(*room); ok {
			if !r.game.Started() {
				sender.out <- msg.NewMessage("error", "game not started")
				return
			}

			if player := r.clients[sender]; player != nil {
				player.Move(message.Direction)
			} else {
				sender.out <- msg.NewMessage("error", "not a player")
			}
		}
	}
}

// TODO Move to game package
func parseDig(data json.RawMessage, sender *client) {
	message := new(msg.GameMessage)
	if err := message.Parse(data); err != nil {
		sender.out <- msg.NewMessage("error", err.Error())
		return
	}

	// Find room name if none sent
	if message.Room == "" {
		if message.Room = findRoom(sender); message.Room == "" {
			sender.out <- msg.NewMessage("error", "not in a room")
			return
		}
	}

	if r, ok := rooms.Load(message.Room); ok {
		if r, ok := r.(*room); ok {
			if !r.game.Started() {
				sender.out <- msg.NewMessage("error", "game not started")
				return
			}

			if runner, ok := r.clients[sender].(*game.Runner); ok {
				runner.Dig(message.Direction)
			} else {
				sender.out <- msg.NewMessage("error", "not a runner")
			}
		}
	}
}
