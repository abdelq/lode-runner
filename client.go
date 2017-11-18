package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"sync"
)

type client struct {
	conn net.Conn
	once sync.Once
	out  chan message
}

func newClient(conn net.Conn) *client {
	client := &client{conn: conn, out: make(chan message, 5)} // XXX

	// Listeners
	go client.read()
	go client.write()

	log.Printf("New connection from %s", conn.RemoteAddr())

	return client
}

func (c *client) close() {
	c.once.Do(func() {
		// Leave all joined rooms
		for _, room := range rooms {
			if _, ok := room.clients[c]; ok {
				room.leave <- c
			}
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
		msg := new(message)
		if err := dec.Decode(msg); err != nil {
			if err == io.EOF {
				break
			}
			c.out <- newMessage("error", err.Error())
			continue
		}
		go msg.parse(c)
	}
}

func (c *client) write() {
	//defer c.close() // XXX

	enc := json.NewEncoder(c.conn)
	for msg := range c.out {
		if err := enc.Encode(msg); err != nil {
			log.Println(err)
			break
		}
	}
}
