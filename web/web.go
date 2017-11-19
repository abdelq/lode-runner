package web

import (
	"io"
	"log"
	"net"
	"net/http"

	//"golang.org/x/net/websocket"
	"github.com/gorilla/websocket"
)

var httpAddr, tcpAddr *string
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

//func proxyServer(ws *websocket.Conn) {
func proxyServer(w http.ResponseWriter, r *http.Request) {
	tcp, err := net.Dial("tcp", *tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer tcp.Close()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	messageType, reader, err := conn.NextReader()
	if err != nil {
		log.Println(err)
		return
	}
	writer, err := conn.NextWriter(messageType)
	if err != nil {
		log.Println(err)
		return
	}

	go io.Copy(writer, tcp)
	io.Copy(tcp, reader)
}

func Listen(addr, serverAddr *string) {
	httpAddr, tcpAddr = addr, serverAddr

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		proxyServer(w, r)
	})
	http.Handle("/", http.FileServer(http.Dir("public")))

	log.Printf("Listening on %s %s", "http", *httpAddr)
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}
