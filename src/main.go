package main

import (
	"log"
	"net/http"

	"lumen-server/server"
)

func main() {

	lumenServer := server.NewServer()

	log.SetFlags(0)
	http.HandleFunc("/", lumenServer.ConnectioHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
