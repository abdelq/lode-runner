package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	addr     = flag.String("addr", ":3000", "server address")
	upgrader = websocket.Upgrader{}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	player := &Player{conn: conn}

	go player.read()
	go player.write()
}

func main() {
	flag.Parse()

	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
