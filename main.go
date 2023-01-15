package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsHandler)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	defer conn.Close()

	reader(conn)
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(p))

		msg := "You sent : " + string(p)

		if err := conn.WriteMessage(messageType, []byte(msg)); err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	fmt.Println("hello")
	setupRoutes()
	http.ListenAndServe(":8181", nil)
}
