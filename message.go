package main

import (
	"encoding/json"
	"log"
	//"strings"

	"github.com/abdelq/lode-runner/game"
)

type message struct {
	Event string
	Data  json.RawMessage
}

func parseJoin(data json.RawMessage, sender *client) {
	var joinData struct{ Name, Room, Role string }
	if err := json.Unmarshal(data, &joinData); err != nil {
		log.Println(err)
		return
	}
	//joinData.Name = strings.TrimSpace(joinData.Name)                  // TODO
	//joinData.Room = strings.TrimSpace(joinData.Room)                  // TODO
	//joinData.Role = strings.ToLower(strings.TrimSpace(joinData.Role)) // TODO

	if joinData.Name == "" {
		sender.out <- &message{"error", json.RawMessage(`"invalid name"`)}
		return
	}
	if joinData.Room == "" {
		sender.out <- &message{"error", json.RawMessage(`"invalid room"`)}
		return
	}

	room, ok := rooms[joinData.Room]
	if !ok {
		room = newRoom(joinData.Room)
	}

	switch joinData.Role {
	case "", "runner": // Runner
		room.join <- &join{sender, &game.Runner{Name: joinData.Name}}
	case "guard": // Guard
		room.join <- &join{sender, &game.Guard{Name: joinData.Name}}
	default: // Spectator
		room.join <- &join{sender, nil}
	}
}

// TODO Move to game package
func parseGame(msg *message, sender *client) {
	var gameData struct{ Direction, Room string }
	if err := json.Unmarshal(msg.Data, &gameData); err != nil {
		log.Println(err)
		return
	}
	//gameData.Direction = strings.ToLower(strings.TrimSpace(gameData.Direction)) // TODO
	//gameData.Room = strings.TrimSpace(gameData.Room)                            // TODO

	if gameData.Direction == "" {
		sender.out <- &message{"error", json.RawMessage(`"invalid direction"`)}
		return
	}
	if gameData.Room == "" { // TODO Find a room with client
		sender.out <- &message{"error", json.RawMessage(`"invalid room"`)}
		return
	}

	if room, ok := rooms[gameData.Room]; ok {
		player := room.clients[sender]
		if msg.Event == "move" {
			if player != nil {
				go player.Move(gameData.Direction, room.game)
			} else {
				sender.out <- &message{"error", json.RawMessage(`"not a player"`)}
			}
		} else if msg.Event == "dig" {
			if runner, ok := player.(*game.Runner); ok {
				go runner.Dig(gameData.Direction, room.game)
			} else {
				sender.out <- &message{"error", json.RawMessage(`"not a runner"`)}
			}
		}
	}
}

func (m *message) parse(sender *client) {
	//m.Event = strings.ToLower(strings.TrimSpace(m.Event)) // TODO

	if m.Event == "join" {
		go parseJoin(m.Data, sender)
	} else if m.Event == "move" || m.Event == "dig" { // TODO Move to game package
		go parseGame(m, sender) // TODO Rename
	} else { // TODO Move to game package
		sender.out <- &message{"error", json.RawMessage(`"invalid event"`)}
	}
}
