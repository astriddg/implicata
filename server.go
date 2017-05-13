package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type (
	// fieldID is the unique ID of given HTML element.
	fieldID string
	// copyAndPaste contains a list of copy and pasted HTML element IDs.
	copyAndPaste map[fieldID]bool
	//sessionID is the UUID of a unique session.
	sessionID string
)

// dimension contains the width/height of a screen size.
type dimension struct {
	Width  string `json:"width"`
	Height string `json:"height"`
}

// request contains the data of each HTTP request.
type request struct {
	WebsiteURL         string       `json:"website_url"`
	SessionID          sessionID    `json:"session_id"`
	ResizeFrom         dimension    `json:"resize_from"`
	ResizeTo           dimension    `json:"resize_to"`
	CopyAndPaste       copyAndPaste `json:"copy_and_paste"`
	FormCompletionTime int          `json:"form_completion_time"`
}

// String formats data in a readable output.
func (req request) String() string {
	var b bytes.Buffer
	if req.WebsiteURL != "" {
		b.WriteString(fmt.Sprintf("WebsiteURL: %s\n", req.WebsiteURL))
	}
	if req.SessionID != "" {
		b.WriteString(fmt.Sprintf("SessionID: %s\n", req.SessionID))
	}
	if req.ResizeFrom.Width != "" && req.ResizeFrom.Height != "" {
		b.WriteString(fmt.Sprintf("ResizeFrom: %sx%s\n", req.ResizeFrom.Width, req.ResizeFrom.Height))
	}
	if req.ResizeTo.Width != "" && req.ResizeTo.Height != "" {
		b.WriteString(fmt.Sprintf("ResizeTo: %sx%s\n", req.ResizeTo.Width, req.ResizeTo.Height))
	}
	if len(req.CopyAndPaste) > 0 {
		var cp bytes.Buffer
		cp.WriteString("CopyAndPaste: ")
		for key, val := range req.CopyAndPaste {
			cp.WriteString(fmt.Sprintf("#%s=%t;", key, val))
		}
		b.WriteString(fmt.Sprintf("%s\n", cp.String()))
	}
	if req.FormCompletionTime != 0 {
		b.WriteString(fmt.Sprintf("FormCompletionTime: %ds\n", req.FormCompletionTime))
	}
	return b.String()
}

// submitHandler handles HTTP requests for saving data.
func submitHandler(stream chan request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("error reading request body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		req := request{
			CopyAndPaste: make(copyAndPaste),
		}
		if err := json.Unmarshal(body, &req); err != nil {
			log.Printf("error unmarshalling JSON request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		select {
		case stream <- req:
			w.WriteHeader(http.StatusNoContent)
		case <-ctx.Done():
			err := ctx.Err()
			if err == context.DeadlineExceeded {
				w.WriteHeader(http.StatusRequestTimeout)
			}
			log.Printf("context done: %v", err)
		}
	}
}
