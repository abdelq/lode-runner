package web

import (
	"flag"
	"io"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/websocket"
)

var httpAddr = flag.String("http", ":7331", "HTTP network address")
var TCPAddr *string

func proxyServer(ws *websocket.Conn) {
	tcp, err := net.Dial("tcp", *TCPAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer tcp.Close()

	go io.Copy(ws, tcp)
	io.Copy(tcp, ws)
}

func Listen() {
	http.Handle("/ws", websocket.Handler(proxyServer))
	http.Handle("/", http.FileServer(http.Dir("public")))

	log.Printf("Listening on %s %s", "http", *httpAddr)
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}
