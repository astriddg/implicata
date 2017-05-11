package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// fieldID is the HTML ID of given dom element.
type fieldID string

// copyAndPaste contains a list of copy and pasted HTML elements.
type copyAndPaste map[fieldID]bool

// dimension contains the width/height of a screen size.
type dimension struct {
	Width  string `json:"width"`
	Height string `json:"height"`
}

// data contains the data of each HTTP request.
type data struct {
	WebsiteURL         string       `json:"website_url"`
	SessionID          string       `json:"session_id"`
	ResizeFrom         dimension    `json:"resize_from"`
	ResizeTo           dimension    `json:"resize_to"`
	CopyAndPaste       copyAndPaste `json:"copy_and_paste"`
	FormCompletionTime int          `json:"form_completion_time"`
}

// String formats data in a readable output.
func (d data) String() string {
	var b bytes.Buffer
	if d.WebsiteURL != "" {
		b.WriteString(fmt.Sprintf("WebsiteURL: %s\n", d.WebsiteURL))
	}
	if d.SessionID != "" {
		b.WriteString(fmt.Sprintf("SessionID: %s\n", d.SessionID))
	}
	if d.ResizeFrom.Width != "" && d.ResizeFrom.Height != "" {
		b.WriteString(fmt.Sprintf("ResizeFrom: %sx%s\n", d.ResizeFrom.Width, d.ResizeFrom.Height))
	}
	if d.ResizeTo.Width != "" && d.ResizeTo.Height != "" {
		b.WriteString(fmt.Sprintf("ResizeTo: %sx%s\n", d.ResizeTo.Width, d.ResizeTo.Height))
	}
	if len(d.CopyAndPaste) > 0 {
		var cp bytes.Buffer
		cp.WriteString("CopyAndPaste: ")
		for key, val := range d.CopyAndPaste {
			cp.WriteString(fmt.Sprintf("#%s=%t;", key, val))
		}
		b.WriteString(fmt.Sprintf("%s\n", cp.String()))
	}
	if d.FormCompletionTime != 0 {
		b.WriteString(fmt.Sprintf("FormCompletionTime: %ds\n", d.FormCompletionTime))
	}
	return b.String()
}

// submitHandler handles HTTP requests for saving data.
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(d)
	w.WriteHeader(http.StatusNoContent)
}
