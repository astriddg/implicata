package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// fieldID ..
type fieldID string

// data ...
type data struct {
	WebsiteURL         url.URL          `json:"website_url"`
	SessionID          string           `json:"session_id"`
	ResizeFrom         dimension        `json:"resize_from"`
	ResizeTo           dimension        `json:"resize_to"`
	CopyAndPaste       map[fieldID]bool `json:"copy_and_paste"`
	FormCompletionTime time.Duration    `json:"form_completion_time"`
}

// dimension ...
type dimension struct {
	Width  string
	Height string
}

// submitHandler ...
func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
