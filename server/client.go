package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"sync"
)

type client struct {
	player player
	room   *room
	conn   net.Conn
	once   sync.Once
	out    chan *message
}

func newClient(conn net.Conn) *client {
	client := &client{conn: conn, out: make(chan *message)}

	// Listeners
	go client.read()
	go client.write()

	log.Printf("New connection from %s", conn.RemoteAddr())
	return client
}

func (c *client) close() {
	c.once.Do(func() {
		if c.room != nil {
			c.room.leave <- c
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
		msg := &message{}
		if err := dec.Decode(msg); err != nil {
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
