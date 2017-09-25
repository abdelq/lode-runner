package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"log"
)

const (
	name = "316k"
	room = "twado"
)

type message struct {
	Event string
	Data  json.RawMessage
}

func join(conn *tls.Conn) {
	data := fmt.Sprintf(`{"name": "%s", "room": "%s"}`, name, room)

	msg := message{
		Event: "join",
		Data:  json.RawMessage(data),
	}

	b, err := json.Marshal(&msg)
	if err != nil {
		log.Println(err)
		return
	}
	conn.Write(b)
	conn.Write([]byte("\n"))

	for {
	}
}

func main() {
	// Flags
	addr := flag.String("addr", ":443", "listener's network address")
	flag.Parse()
	log.SetFlags(log.Ltime)

	// Connect to the server
	conf := tls.Config{InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", *addr, &conf)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	join(conn)
}
