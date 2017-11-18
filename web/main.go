package web

import (
	"io"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/websocket"
)

var httpAddr, tcpAddr *string

func proxyServer(ws *websocket.Conn) {
	tcp, err := net.Dial("tcp", *tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer tcp.Close()

	go io.Copy(ws, tcp)
	io.Copy(tcp, ws)
}

func Listen(addr, serverAddr *string) {
	httpAddr, tcpAddr = addr, serverAddr

	http.Handle("/ws", websocket.Handler(proxyServer))
	http.Handle("/", http.FileServer(http.Dir("public")))

	log.Printf("Listening on %s %s", "http", *httpAddr)
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}
