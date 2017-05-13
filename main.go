package main

import (
	"log"
	"net/http"
)

func main() {
	// create request stream.
	stream := make(chan data)
	// proccess stream data.
	go process(stream)

	// configure HTTP endpoints.
	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.Handle("/submit", submitHandler(stream))
	//run HTTP server.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
