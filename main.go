package main

import (
	"chatr/server"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"fmt"
)


func main() {
	http.Handle("/socket", websocket.Handler(server.SocketHandler))
	http.HandleFunc("/", server.RootHandler)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
