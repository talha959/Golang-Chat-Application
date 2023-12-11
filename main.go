package main

import (
	handleconnections "chat/handleConnections"
	handlemessages "chat/handleMessages"
	homepage "chat/homePage"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homepage.Home)
	http.HandleFunc("/ws", handleconnections.Handle)

	go handlemessages.Message()

	fmt.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
