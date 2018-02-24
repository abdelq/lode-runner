package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"sync"

	msg "github.com/abdelq/lode-runner/message"
)

type client struct {
	conn net.Conn
	once sync.Once
	out  chan *msg.Message
}

func newClient(conn net.Conn) *client {
	client := &client{conn: conn, out: make(chan *msg.Message, 1)}

	// Listeners
	go client.read()
	go client.write()

	log.Printf("New connection from %s", conn.RemoteAddr())

	return client
}

func (c *client) close() {
	c.once.Do(func() {
		// Leave all joined rooms
		rooms.Range(func(n, r interface{}) bool {
			if _, ok := r.(*room).clients[c]; ok {
				r.(*room).leave <- c
			}
			return true
		})

		//close(c.out)
		c.conn.Close()

		log.Printf("Closed connection from %s", c.conn.RemoteAddr())
	})
}

func (c *client) read() {
	defer c.close()

	dec := json.NewDecoder(c.conn)
	for {
		message := new(message)
		if err := dec.Decode(message); err != nil {
			if err == io.EOF || err == io.ErrClosedPipe {
				break
			}
			if _, ok := err.(net.Error); ok {
				break
			}
			c.out <- msg.NewMessage("error", err.Error())
			continue
		}
		go message.parse(c)
	}
}

func (c *client) write() {
	defer c.close()

	enc := json.NewEncoder(c.conn)
	for msg := range c.out {
		if err := enc.Encode(msg); err != nil {
			if _, ok := err.(net.Error); ok {
				break
			}
			log.Println(err)
		}
	}
}
