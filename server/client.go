package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"sync"
)

type client struct {
	name, room string
	conn       net.Conn
	once       sync.Once
	out        chan *message
}

func newClient(conn net.Conn) *client {
	client := &client{
		conn: conn,
		out:  make(chan *message),
	}

	// Listeners
	go client.read()
	go client.write()

	return client
}

func (c *client) close() {
	c.once.Do(func() {
		if room := rooms[c.room]; room != nil {
			room.leave <- c
		}

		close(c.out)
		c.conn.Close()
		log.Printf("Closed connection from %s", c.conn.RemoteAddr())
	})
}

func (c *client) read() {
	defer c.close()

	dec := json.NewDecoder(c.conn)
	for {
		var msg message
		if err := dec.Decode(&msg); err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}
		go msg.parse(c)
	}
}

func (c *client) write() {
	defer c.close()

	enc := json.NewEncoder(c.conn)
	for msg := range c.out {
		if err := enc.Encode(msg); err != nil {
			log.Println(err)
			break
		}
	}
}

func (c *client) join(clientName, roomName string) {
	if clientName == "" || roomName == "" {
		c.close()
		return
	}
	c.name, c.room = clientName, roomName

	room, ok := rooms[roomName]
	if !ok {
		room = newRoom(roomName)
	} else if room.hasClient(clientName) {
		c.close()
		return
	}
	room.join <- c
}
