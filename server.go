package main

import (
	"log"
	"net/http"
	"text/template"

	"golang.org/x/net/websocket"
)

var sockets = make(map[string]*websocket.Conn)

func fishHandler(ws *websocket.Conn) {
	id := ws.RemoteAddr().String()
	sockets[ws.RemoteAddr().String()] = ws

	// Wait here
	log.Println(id, "is waiting")
	msg := make([]byte, 512)
	_, err := ws.Read(msg)

	if err != nil {
		// Client should exit
		log.Println(id, "client disconnect", err)
		delete(sockets, id)
	}
}

func hook(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	rurl := r.Referer()
	log.Println("Hook requested on", id, "with referer", rurl)
	if ws, ok := sockets[id]; ok {
		w.Write([]byte("Hooking " + id))
		_, err := ws.Write([]byte(rurl))
		if err != nil {
			log.Fatal("Error whilst writing", err)
		}
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	http.Handle("/fish", websocket.Handler(fishHandler))
	http.HandleFunc("/hook/", hook)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("server/index.html")
		keys := make([]string, 0, len(sockets))
		for k := range sockets {
			keys = append(keys, k)
		}
		t.Execute(w, keys)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
