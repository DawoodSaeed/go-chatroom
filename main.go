package main

import (
	"net/http"

	"github.com/DawoodSaeed/go-chatroom/socket"
)

func main() {
	socket := socket.NewSocket()
	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	// This is the golang socket route;
	http.HandleFunc("/ws", socket.EstablishSocketConn)
	http.ListenAndServe(":8080", nil)

}
