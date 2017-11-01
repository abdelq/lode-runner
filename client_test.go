package main

import (
	"io"
	"net"
	"testing"
)

// TODO
func TestNewClient(t *testing.T) {}

// TODO
func TestRead(t *testing.T) {}

// TODO
func TestWrite(t *testing.T) {}

func TestClose(t *testing.T) {
	conn, _ := net.Pipe()
	client := newClient(conn)

	// Join rooms
	rooms := []*room{newRoom("Buzz"), newRoom("Rex"), newRoom("Bo")}
	for _, room := range rooms {
		room.join <- &join{client, nil}
	}

	client.close() // First
	client.close() // Second

	// Verify rooms are left
	for _, room := range rooms {
		if _, ok := room.clients[client]; ok {
			t.Errorf("client still in a room")
			break
		}
	}

	// Verify output channel is closed
	if _, ok := <-client.out; ok {
		t.Error("output channel not closed")
	}

	// Verify connection is closed
	if _, err := conn.Read(nil); err != io.ErrClosedPipe {
		t.Error(err)
	}
}
