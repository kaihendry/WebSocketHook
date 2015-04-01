package main

import (
	"fmt"
	"html"
	"log"
	"sync"
	"net/http"

	"golang.org/x/net/websocket"
)

var mu sync.RWMutex
var sockets = make(map[string]*websocket.Conn)

func fishHandler(ws *websocket.Conn) {
	id := ws.RemoteAddr().String()
	sockets[ws.RemoteAddr().String()] = ws
	// wait here
	log.Println(id, "is waiting")
	msg := make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Received: %s\n", msg[:n])
	// Client should exit
}

func main() {
	http.Handle("/fish", websocket.Handler(fishHandler))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for key, _ := range sockets {
			fmt.Fprintf(w, "%q\n", html.EscapeString(key))
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
