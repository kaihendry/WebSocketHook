package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func fishHandler(ws *websocket.Conn) {

	// I guess it's waiting here

	// Need to send it a message from hook
	// But how?

	log.Println(ws.RemoteAddr(), "is waiting")
	_, err := ws.Write([]byte("htp://example.com"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Sent")
}

func main() {
	http.Handle("/fish", websocket.Handler(fishHandler))
	http.HandleFunc("/hook", func(w http.ResponseWriter, r *http.Request) {
		// Be good to show what's connected
		// And send referer URL (or some other URL) to fishHandler
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
