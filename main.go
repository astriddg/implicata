package main

import (
	"log"
	"net/http"
)

func main() {
	stream := make(chan data)
	go process(stream)

	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.Handle("/submit", submitHandler(stream))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
