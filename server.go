package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func echoHandler(ws *websocket.Conn) {
	_, err := ws.Write([]byte("exit"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Sent")
}

func main() {
	http.Handle("/echo", websocket.Handler(echoHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
