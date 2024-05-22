package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients []websocket.Conn

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		clients = append(clients, *conn)
		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			fmt.Printf("%s send: %s", conn.RemoteAddr(), string(msg))
			for _,client := range clients {
				if err := client.WriteMessage(msgType, msg); err != nil {
					log.Fatal(err)
					return
				}

			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")

	})
	println("Your server run in :8080")
	http.ListenAndServe("192.168.137.163:80", nil)

}
