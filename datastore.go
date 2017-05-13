package main

import "fmt"

// dataStore stores the state of each unique request.
type dataStore map[sessionID]*request

// process receives data from a stream and saves it to a data store.
func process(stream chan request) {
	store := make(dataStore)
	for {
		req := <-stream
		save(req, store)
	}
}

// save updates a unique data set in a given data store.
func save(req request, store dataStore) {
	session := req.SessionID
	data, ok := store[session]
	if !ok {
		store[session] = &req
		fmt.Println(&req)
		return
	}
	if req.ResizeFrom.Height != "" && req.ResizeFrom.Width != "" {
		data.ResizeFrom = req.ResizeFrom
	}
	if req.ResizeTo.Height != "" && req.ResizeTo.Width != "" {
		data.ResizeTo = req.ResizeTo
	}
	if len(req.CopyAndPaste) > 0 {
		for k, v := range req.CopyAndPaste {
			data.CopyAndPaste[k] = v
		}
	}
	if req.FormCompletionTime > 0 {
		data.FormCompletionTime = req.FormCompletionTime
	}
	fmt.Println(data)
}
