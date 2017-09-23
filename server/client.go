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
	out        chan []byte
	once       sync.Once
}

func newClient(conn net.Conn) *client {
	client := &client{
		conn: conn,
		out:  make(chan []byte),
	}

	// Listeners
	go client.read()
	go client.write()

	return client
}

func (c *client) close() {
	c.once.Do(func() {
		if r := rooms[c.room]; r != nil {
			r.leave <- c
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

	// TODO Use JSON encoder
	for msg := range c.out {
		_, err := c.conn.Write(msg)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func (c *client) join(name, room string) {
	if name == "" || room == "" {
		// TODO Error message
		c.close()
		return
	}
	c.name, c.room = name, room

	r, ok := rooms[room]
	if !ok {
		newRoom(room).join <- c
		return
	}

	// Verify uniqueness of name in room
	for client := range r.clients {
		if client.name == name {
			// TODO Error message
			c.close()
			return
		}
	}
	r.join <- c
}
