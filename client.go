package main

import (
	"fmt"
	"net/url"
)

import (
	"log"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	id := "http://localhost:8080/unique-identifer-goes-here"
	pond := "ws://localhost:8080/fish"
	var msg = make([]byte, 512)

	var err error
	var ws *websocket.Conn
	failcount := 1

Loop:
	for {
		ws, err = websocket.Dial(pond, "", id)
		if err != nil {
			log.Println("Connection failed, re-trying ", failcount)
			failcount++
			time.Sleep(5 * time.Second)
			continue
		}
		log.Printf("Connected to %s", pond)

		n, err := ws.Read(msg)

		if err != nil {
			log.Println("Error reading", err)
		}

		log.Printf("Received: %s\n", msg)
		rurl := string(msg[:n])

		u, err := url.ParseRequestURI(rurl)
		if err == nil {
			switch u.Scheme {
			case "http", "https":
				ws.Close()
				fmt.Println(u)
				break Loop
			default:
				log.Println("Non-URL returned:", rurl)
			}
		}

		time.Sleep(1 * time.Second)
	}
}
