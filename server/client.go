package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
)

type Client struct {
	name, room string
	conn       net.Conn
	out        chan string
}

func newClient(conn net.Conn) *Client {
	client := &Client{conn: conn, out: make(chan string)}

	go client.read()
	go client.write()

	return client
}

func (c *Client) read() {
	reader := bufio.NewReader(c.conn)

	defer func() {
		if rooms[c.room] != nil {
			rooms[c.room].leave <- c
		}
		close(c.out)
		c.conn.Close()
	}()

	for {
		data, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}

		msg := &Message{client: c}
		if err := json.Unmarshal(data, msg); err != nil {
			log.Println(err)
			break
		}
		go msg.parse()
	}
}

func (c *Client) write() {
	writer := bufio.NewWriter(c.conn)

	defer c.conn.Close()

	for msg := range c.out {
		_, err := writer.WriteString(msg)
		if err != nil {
			log.Println(err)
			break
		}
		writer.Flush()
	}
}

func (c *Client) join(name, room string) {
	if name == "" || room == "" {
		return
	}
	c.name, c.room = name, room

	if _, ok := rooms[room]; !ok {
		newRoom(room)
	}
	rooms[room].join <- c
}
