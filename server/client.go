package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	id     uuid.UUID
	socket *websocket.Conn
	room   *Room
	send   chan []byte
}

func (c *Client) join(roomName string) {
	room, ok := rooms[roomName]
	if !ok {
		room = newRoom(roomName)
	}

	c.room = room
	room.join <- c
}

func (c *Client) read() {
	defer func() {
		c.room.leave <- c
		c.socket.Close()
	}()

	for {
		msg := Message{}
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			break
		}

		if msg.Event == "join" {
			c.join(string(msg.Data))
		}

		// TODO Broadcast
	}
}

func (c *Client) write() {
	defer c.socket.Close()

	for {
		select {
		case msg, ok := <-c.send:
			if !ok {
				// Channel closed by room
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
