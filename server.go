package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func fishHandler(ws *websocket.Conn) {

	// I guess it's waiting here

	// Need to send it a message from hook
	// But how?

	log.Println(ws.RemoteAddr(), "is waiting")
	_, err := ws.Write([]byte("http://example.com"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Sent")
}

func main() {
	http.Handle("/fish", websocket.Handler(fishHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
