package main

import (
	"encoding/json"
	"log"
)

type message struct {
	Event string
	Data  json.RawMessage
}

func parseJoin(data json.RawMessage, sender *client) {
	var joinData struct {
		Name, Room string
		Role       uint8
	}
	if err := json.Unmarshal(data, &joinData); err != nil {
		log.Println(err)
		return
	}

	// Validate name and room
	if joinData.Name == "" || joinData.Room == "" {
		sender.out <- &message{"error",
			json.RawMessage(`"invalid name or room"`)}
		return
	}

	room, ok := rooms[joinData.Room]
	if !ok {
		room = newRoom(joinData.Room)
	}

	switch joinData.Role {
	case 0: // Runner
		room.join <- &join{sender, &runner{name: joinData.Name}}
	case 1: // Guard
		room.join <- &join{sender, &guard{name: joinData.Name}}
	default: // Spectator
		room.join <- &join{sender, nil}
	}
}

func parseMove(data json.RawMessage, sender *client) {
	var moveData struct{ Direction, Room string }
	if err := json.Unmarshal(data, &moveData); err != nil {
		log.Println(err)
		return
	}

	// Validate direction and room
	if moveData.Direction == "" || moveData.Room == "" {
		sender.out <- &message{"error",
			json.RawMessage(`"invalid direction or room"`)}
		return
	}

	// TODO Find a room with client if none declared
	if room, ok := rooms[moveData.Room]; ok {
		if player := room.clients[sender]; player != nil {
			go player.move(moveData.Direction, room.game)
		} else {
			sender.out <- &message{"error",
				json.RawMessage(`"not a player"`)}
		}
	}
}

func parseDig(data json.RawMessage, sender *client) {
	var digData struct{ Direction, Room string }
	if err := json.Unmarshal(data, &digData); err != nil {
		log.Println(err)
		return
	}

	// Validate direction and room
	if digData.Direction == "" || digData.Room == "" {
		sender.out <- &message{"error",
			json.RawMessage(`"invalid direction or room"`)}
		return
	}

	// TODO Find a room with client if none declared
	if room, ok := rooms[digData.Room]; ok {
		if runner, ok := room.clients[sender].(*runner); ok {
			go runner.dig(digData.Direction, room.game)
		} else {
			sender.out <- &message{"error",
				json.RawMessage(`"not a runner"`)}
		}
	}
}

func (m *message) parse(sender *client) {
	switch m.Event {
	case "join":
		go parseJoin(m.Data, sender)
	case "move":
		go parseMove(m.Data, sender)
	case "dig":
		go parseDig(m.Data, sender)
	default:
		sender.out <- &message{"error",
			json.RawMessage(`"invalid event"`)}
	}
}
