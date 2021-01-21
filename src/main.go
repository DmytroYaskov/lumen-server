package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func service(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Print("read:", err)
			break
		}
		log.Printf("Received: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Print("write:", err)
			break
		}
	}
}

func main() {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	log.SetFlags(0)
	http.HandleFunc("/", service)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
