package main

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Sender   string
	Receiver string
	Content  string
}

var clients = make(map[*websocket.Conn]string)
var parior = make(map[string]*websocket.Conn)

func HandleUpgrader(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		panic(err)
		return
	}

	defer conn.Close()

	userId := r.URL.Query().Get("user_id")

	clients[conn] = userId
	parior[userId] = conn

	defer func() {
		delete(clients, conn)
		delete(parior, userId)
	}()

}

func handleMessage(conn *websocket.Conn) {
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("There is some error: ", err)
			continue
		}

		var message Message

		if jsonerr := json.Unmarshal(msg, &message); jsonerr != nil {
			fmt.Println("There is unmarshal error: ", jsonerr)
			continue
		}

		recipientconn, ok := parior[message.Receiver]

		if !ok {
			fmt.Println("There is conn error")
			continue
		}

		senderr := recipientconn.WriteMessage(msgType, []byte(message.Content))
		if senderr != nil {
			fmt.Println("sending error:", senderr)
			continue
		}

	}
}

func main() {
	http.HandleFunc("/ws", HandleUpgrader)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")

	})
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8090", nil)
}
