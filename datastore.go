package main

import "fmt"

// dataStore stores the state of each unique request.
type dataStore map[sessionID]*data

// process receives data from a stream and saves it to a data store.
func process(stream chan data) {
	store := make(dataStore)
	for {
		req := <-stream
		save(req, store)
	}
}

// save updates a unique data set in a given data store.
func save(req data, store dataStore) {
	session := req.SessionID
	if _, ok := store[session]; !ok {
		store[session] = &req
	}
	store[session].WebsiteURL = req.WebsiteURL
	store[session].SessionID = req.SessionID
	if req.ResizeFrom.Height != "" && req.ResizeFrom.Width != "" {
		store[session].ResizeFrom = req.ResizeFrom
	}
	if req.ResizeTo.Height != "" && req.ResizeTo.Width != "" {
		store[session].ResizeTo = req.ResizeTo
	}
	if len(req.CopyAndPaste) > 0 {
		for k, v := range req.CopyAndPaste {
			store[session].CopyAndPaste[k] = v
		}
	}
	if req.FormCompletionTime > 0 {
		store[session].FormCompletionTime = req.FormCompletionTime
	}
	fmt.Println(store[session])
}
