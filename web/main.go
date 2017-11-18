package web

import (
	"io"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/websocket"
)

// XXX
func EchoServer(ws *websocket.Conn) {
	conn, err := net.Dial("tcp", ":1337")
	if err != nil {
		log.Println(err)
	}

	defer conn.Close() // XXX

	go io.Copy(ws, conn)
	io.Copy(conn, ws)
}

func Listen(addr *string) {
	mux := http.NewServeMux()

	mux.Handle("/ws", websocket.Handler(EchoServer)) // XXX /ws and name of function
	mux.Handle("/", http.FileServer(http.Dir("./public")))

	log.Fatal(http.ListenAndServe(":8080", mux)) // XXX
}
