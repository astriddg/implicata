package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var httpPort uint

func init() {
	// configure optional server port
	flag.UintVar(&httpPort, "http", 8080, "Port number of HTTP server")
	flag.Parse()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.HandleFunc("/submit", submitHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil))
}
