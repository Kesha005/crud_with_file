package main

import (
	//"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)



var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients []Client



type Client struct{
	Conn websocket.Conn
	Name string
}


func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		
		client_name := r.Header.Get("Client-Name")
		if client_name ==""{
			panic("client name cannot be empty")

			return 
		}

		client:= &Client{Name: client_name,Conn: *conn}

		clients = append(clients, *client)

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")

	})
	println("Your server run in :8080")
	log.Fatal(http.ListenAndServe(":8090", nil))

}
