package main

import "fmt"

// dataStore stores the state of each unique request.
type dataStore map[sessionID]*data

// process receives data from a stream and saves to a data store.
func process(stream chan data) {
	store := make(dataStore)
	for {
		req := <-stream
		save(req, store)
	}
}

// save updates a data set in a given data store.
func save(req data, store dataStore) {
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
