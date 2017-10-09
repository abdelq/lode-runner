package main

import (
	"encoding/json"
	"log"
)

type message struct {
	Event string
	Data  *json.RawMessage
}

func parseJoin(data *json.RawMessage, sender *client) {
	var playerData struct {
		Name, Room string
		IsGuard    bool `json:",omitempty"`
	}
	if err := json.Unmarshal(*data, &playerData); err != nil {
		log.Println(err)
		return
	}

	// TODO verify duplicate names
	if playerData.Name == "" {
		//sender.out <- &message{"error", &json.RawMessage(`"invalid name"`)}
		return
	}
	if playerData.Room == "" {
		//sender.out <- &message{"error", &json.RawMessage(`"invalid room"`)}
		return
	}

	if playerData.IsGuard {
		sender.player = &guard{name: playerData.Name}
	} else {
		sender.player = &runner{name: playerData.Name}
	}

	if room, ok := rooms[playerData.Room]; ok {
		room.join <- sender
	} else {
		newRoom(playerData.Room).join <- sender
	}
}

func parseMove(data *json.RawMessage, player player) {
	var direction string
	if err := json.Unmarshal(*data, direction); err != nil {
		log.Println(err)
		return
	}

	if player != nil {
		go player.move(direction)
	}
}

func parseDig(data *json.RawMessage, runner *runner) {
	var direction string
	if err := json.Unmarshal(*data, direction); err != nil {
		log.Println(err)
		return
	}

	if runner != nil {
		go runner.dig(direction)
	}
}

func (m *message) parse(sender *client) {
	switch m.Event {
	case "join":
		go parseJoin(m.Data, sender)
	case "move":
		go parseMove(m.Data, sender.player)
	case "dig":
		go parseDig(m.Data, sender.player.(*runner))
	default:
		err := json.RawMessage(`"invalid event"`)
		sender.out <- &message{"error", &err}
	}
}
