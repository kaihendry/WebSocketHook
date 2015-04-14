package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"golang.org/x/net/websocket"
)

var sockets = make(map[string]*websocket.Conn)

func fishHandler(ws *websocket.Conn) {
	id := ws.RemoteAddr().String() + "-" + ws.Request().RemoteAddr + "-" + ws.Request().UserAgent()
	sockets[id] = ws

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
	log.Println("Request from", r.RemoteAddr)
	err := r.ParseForm()
	if err != nil {
		//handle error http.Error() for example
		return
	}
	log.Println(r.Form)
	w.Write([]byte("Hooking " + strconv.Itoa(len(r.Form["m"])) + " machine(s) with " + r.FormValue("webhook")))

	for _, id := range r.Form["m"] {
		if ws, ok := sockets[id]; ok {
			_, err := ws.Write([]byte(r.FormValue("webhook")))
			if err != nil {
				log.Fatal("Error whilst writing", err)
			} else {
				log.Println("Successfully hooked", id)
			}
		} else {
			log.Println(id, "not waiting")
		}
	}

}

func main() {
	http.Handle("/fish", websocket.Handler(fishHandler))
	http.HandleFunc("/hook/", hook)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/client.html")
	})

	http.HandleFunc("/admin/", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("templates/admin.html")
		keys := make([]string, 0, len(sockets))
		for k := range sockets {
			keys = append(keys, k)
		}
		t.Execute(w, keys)
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
