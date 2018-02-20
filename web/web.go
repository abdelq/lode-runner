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

func proxyServer(ws *websocket.Conn) {
	tcp, err := net.Dial("tcp", flag.Lookup("tcp").Value.String())
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

	log.Printf("Listening on http %s", *httpAddr)
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}
