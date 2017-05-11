package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// fieldID ...
type fieldID string

// copyAndPaste ...
type copyAndPaste map[fieldID]bool

// data ...
type data struct {
	WebsiteURL         string       `json:"website_url"`
	SessionID          string       `json:"session_id"`
	ResizeFrom         dimension    `json:"resize_from"`
	ResizeTo           dimension    `json:"resize_to"`
	CopyAndPaste       copyAndPaste `json:"copy_and_paste"`
	FormCompletionTime int          `json:"form_completion_time"`
}

// dimension ...
type dimension struct {
	Width  string `json:"width"`
	Height string `json:"height"`
}

// submitHandler ...
func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	d := data{
		CopyAndPaste: make(copyAndPaste),
	}
	if err := json.Unmarshal(body, &d); err != nil {
		log.Printf("error unmarshalling JSON request body: %v", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	fmt.Println(d)
}
