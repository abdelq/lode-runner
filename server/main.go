package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		id:     uuid.New(),
		socket: conn,
		send:   make(chan []byte),
	}

	go client.read()
	go client.write()
}

func main() {
	addr := flag.String("addr", ":3000", "server address")
	flag.Parse()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
