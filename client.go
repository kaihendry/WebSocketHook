package main

import (
	"log"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	origin := "http://localhost:8080/"
	url := "ws://localhost:8080/echo"
	var msg = make([]byte, 512)

	var err error
	var ws *websocket.Conn
	failcount := 1
	for {
		ws, err = websocket.Dial(url, "", origin)
		if err != nil {
			log.Println("Connection failed, re-trying ", failcount)
			failcount++
			time.Sleep(5 * time.Second)
			continue
		}
		log.Printf("Connected to %s", url)

		_, err = ws.Read(msg)
		if err != nil {
			log.Println("Error reading", err)
		}
		log.Printf("Receive: %s\n", msg)

	}

}
