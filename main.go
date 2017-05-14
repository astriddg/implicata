package main

import (
	"log"
	"net/http"
	"os"
)

// logger is used to safely print simultaneously from multiple goroutines.
var logger = log.New(os.Stdout, "", 0)

func main() {
	// create request stream.
	stream := make(chan request)
	// proccess stream data.
	go process(stream)

	// configure HTTP endpoints.
	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.Handle("/submit", submitHandler(stream))
	//run HTTP server.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
