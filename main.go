package main

import (
	//"fmt"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients []Client

type Msgs struct {
	To      int
	Sender  int
	Message string
}

type Client struct {
	Conn websocket.Conn
	Name string
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)

		defer conn.Close()

		client_name := r.Header.Get("Client-Name")
		if client_name == "" {
			panic("client name cannot be empty")

			return
		}

		client := &Client{Name: client_name, Conn: *conn}

		clients = append(clients, *client)

		for {

			var msg Msgs

			msgType, message, err := conn.ReadMessage()
			if err != nil {
				panic(err)
				break
			}
			if msgType == websocket.TextMessage {
				if msgerr := json.Unmarshal(message, &msg); msgerr != nil {

					continue
				}

			}

			for _, client := range clients{
				client.Conn.WriteMessage(msgType,message)
			}

		}

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")

	})
	println("Your server run in :8080")
	log.Fatal(http.ListenAndServe(":8090", nil))

}
