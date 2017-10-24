package main

import (
	"encoding/json"
	"log"
)

type message struct {
	Event string
	Data  json.RawMessage
}

// Errors
var errorMsg = map[string]*message{ // TODO Variable name + Structure
	"invalidEvent":     &message{"error", json.RawMessage(`"invalid event"`)},
	"invalidName":      &message{"error", json.RawMessage(`"invalid name"`)},
	"invalidDirection": &message{"error", json.RawMessage(`"invalid direction"`)},
	"invalidRoom":      &message{"error", json.RawMessage(`"invalid room"`)},
	"notAPlayer":       &message{"error", json.RawMessage(`"not a player"`)},
	"notARunner":       &message{"error", json.RawMessage(`"not a runner"`)},
}

func parseJoin(data json.RawMessage, sender *client) {
	var joinData struct{ Name, Room, Role string }
	if err := json.Unmarshal(data, &joinData); err != nil {
		log.Println(err)
		return
	}
	// TODO Trim all + Downcase role

	if joinData.Name == "" {
		sender.out <- errorMsg["invalidName"]
		return
	}
	if joinData.Room == "" {
		sender.out <- errorMsg["invalidRoom"]
		return
	}

	room, ok := rooms[joinData.Room]
	if !ok {
		room = newRoom(joinData.Room)
	}

	switch joinData.Role {
	case "", "runner": // Runner
		room.join <- &join{sender, &runner{name: joinData.Name}}
	case "guard": // Guard
		room.join <- &join{sender, &guard{name: joinData.Name}}
	default: // Spectator
		room.join <- &join{sender, nil}
	}
}

// TODO Move to game package
// TODO Remove repetition with parseDig
func parseMove(data json.RawMessage, sender *client) {
	var moveData struct{ Direction, Room string }
	if err := json.Unmarshal(data, &moveData); err != nil {
		log.Println(err)
		return
	}
	// TODO Trim all + Downcase direction

	if moveData.Direction == "" {
		sender.out <- errorMsg["invalidDirection"]
		return
	}
	if moveData.Room == "" { // TODO Find a room with client
		sender.out <- errorMsg["invalidRoom"]
		return
	}

	if room, ok := rooms[moveData.Room]; ok {
		if player := room.clients[sender]; player != nil {
			go player.move(moveData.Direction, room.game)
		} else {
			sender.out <- errorMsg["notAPlayer"]
		}
	}
}

// TODO Move to game package
// TODO Remove repetition with parseMove
func parseDig(data json.RawMessage, sender *client) {
	var digData struct{ Direction, Room string }
	if err := json.Unmarshal(data, &digData); err != nil {
		log.Println(err)
		return
	}
	// TODO Trim all + Downcase direction

	if digData.Direction == "" {
		sender.out <- errorMsg["invalidDirection"]
		return
	}
	if digData.Room == "" { // TODO Find a room with client
		sender.out <- errorMsg["invalidRoom"]
		return
	}

	if room, ok := rooms[digData.Room]; ok {
		if runner, ok := room.clients[sender].(*runner); ok {
			go runner.dig(digData.Direction, room.game)
		} else {
			sender.out <- errorMsg["notARunner"]
		}
	}
}

func (m *message) parse(sender *client) {
	// TODO Default should send to parser for game directly
	switch m.Event { // TODO Trim + Downcase
	case "join":
		go parseJoin(m.Data, sender)
	case "move":
		go parseMove(m.Data, sender)
	case "dig":
		go parseDig(m.Data, sender)
	default:
		sender.out <- errorMsg["invalidEvent"]
	}
}
