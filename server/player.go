package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type Player struct {
	conn *websocket.Conn
	room *Room
	//send chan []byte
}

type Message struct {
	Type string
	Data string
}

func (player *Player) read() {
	defer func() {
		player.room.leave <- player
		player.conn.Close()
	}()

	//c.conn.SetReadLimit(maxMessageSize)
	//c.conn.SetReadDeadline(time.Now().Add(pongWait))
	//c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		msg := Message{}

		err := player.conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			break
		}

		switch msg.Type {
		case "join":
			room, ok := rooms[msg.Data]
			if !ok {
				room = newRoom()
				rooms[msg.Data] = room

				go room.run()
			}

			room.join <- player
		}

		//player.room.broadcast <- msg
	}
}

func (player *Player) write() {
	//ticker := time.NewTicker(pingPeriod)

	defer func() {
		//ticker.Stop()
		//player.room.leave <- player
		player.conn.Close()
	}()

	for {
		select {
		/*case message, ok := <-player.send:
			//c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				player.conn.WriteMessage(websocket.CloseMessage, []byte{}) // TODO
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println(err)
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}*/
		}
	}
}
