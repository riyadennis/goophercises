package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func handleRoutes() {
	http.HandleFunc("/home", homePageHandler)
	http.HandleFunc("/ws", websocketHandler)
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	for {
		mt, m, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}
		err = conn.WriteMessage(mt, m)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "type %d, message: %s", mt, string(m))
	}
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}

func main() {
	handleRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
