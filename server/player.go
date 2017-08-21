package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type Player struct {
	conn *websocket.Conn
	room *Room
}

type Message struct {
	Type string
	Data string
}

func (player *Player) read() {
	defer player.conn.Close()

	for {
		msg := Message{}

		err := player.conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			break
		}

		if msg.Type == "join" {
			room, ok := rooms[msg.Data]

			if !ok {
				room = newRoom()
				go room.run()

				rooms[msg.Data] = room
			}

			room.join <- player
		}
	}
}
