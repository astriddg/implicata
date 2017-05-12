package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	stream := make(chan data)
	store := make(dataStore)
	go func() {
		for {
			req := <-stream
			if _, ok := store[req.SessionID]; !ok {
				store[req.SessionID] = &req
			}
			store[req.SessionID].WebsiteURL = req.WebsiteURL
			store[req.SessionID].SessionID = req.SessionID
			if req.ResizeFrom.Height != "" && req.ResizeFrom.Width != "" {
				store[req.SessionID].ResizeFrom = req.ResizeFrom
			}
			if req.ResizeTo.Height != "" && req.ResizeTo.Width != "" {
				store[req.SessionID].ResizeTo = req.ResizeTo
			}
			if len(req.CopyAndPaste) > 0 {
				for k, v := range req.CopyAndPaste {
					store[req.SessionID].CopyAndPaste[k] = v
				}
			}
			if req.FormCompletionTime > 0 {
				store[req.SessionID].FormCompletionTime = req.FormCompletionTime
			}
			fmt.Println(store[req.SessionID])
		}
	}()
	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.Handle("/submit", submitHandler(stream))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
