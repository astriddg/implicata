package main

import "testing"

func TestSaveDataStore(t *testing.T) {
	var session sessionID = "123"
	url := "http://www.ravelin.com/"

	store := make(dataStore)
	req := request{
		WebsiteURL:   url,
		SessionID:    session,
		CopyAndPaste: make(copyAndPaste),
	}
	save(req, store)

	if store[session].WebsiteURL != url {
		t.Errorf("expected URL: %s, got: %s", url, store[session].WebsiteURL)
	}

	req.ResizeFrom = dimension{
		Height: "200",
		Width:  "200",
	}
	req.ResizeTo = dimension{
		Height: "150",
		Width:  "150",
	}
	save(req, store)

	if store[session].ResizeFrom.Height != "200" {
		t.Errorf("expected ResizeFrom Height: 200, got: %s", store[session].ResizeFrom.Height)
	}
	if store[session].ResizeFrom.Width != "200" {
		t.Errorf("expected ResizeFrom Width 200, got: %s", store[session].ResizeFrom.Width)
	}
	if store[session].ResizeTo.Height != "150" {
		t.Errorf("expected ResizeTo Height: 150, got: %s", store[session].ResizeTo.Height)
	}
	if store[session].ResizeTo.Height != "150" {
		t.Errorf("expected ResizeTo Width: 150, got: %s", store[session].ResizeTo.Width)
	}

	req.CopyAndPaste = copyAndPaste{"div1": true}
	save(req, store)

	if _, ok := store[session].CopyAndPaste["div1"]; !ok {
		t.Error("did not find given element in map")
	}

	req.FormCompletionTime = 10
	save(req, store)

	if store[session].FormCompletionTime != 10 {
		t.Errorf("expected FormCompletionTime: 10, got: %d", store[session].FormCompletionTime)
	}
}
