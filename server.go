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

func hook(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	log.Println("Here with", id)
	w.Write([]byte("Attempting to hook " + id))
	if w, ok := sockets[id]; ok {
		log.Println("Hooking", id)
		_, err := w.Write([]byte("http://example.com"))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Expecting to delete/client up socket")
	}
}

func main() {
	http.Handle("/fish", websocket.Handler(fishHandler))
	http.HandleFunc("/hook/", hook)
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
