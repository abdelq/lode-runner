package main

import (
	"bufio"
	"log"
	"net"
)

type Client struct {
	conn net.Conn
	out  chan string
	room *Room
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
		c.room.leave <- c
		close(c.out)
		c.conn.Close()
	}()

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}

		// TODO Join a room
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
