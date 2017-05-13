package main

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type mockReader struct{}

func (m mockReader) Read(p []byte) (int, error) {
	return 0, errors.New("read error")
}

type fixture struct {
	body       io.Reader
	httpMethod string
	statusCode int
	ctx        bool
}

func TestStringMethod(t *testing.T) {
	expected := "WebsiteURL: https://www.ravelin.com/\nSessionID: 123456\nResizeFrom: 200x200\nResizeTo: 100x100\nCopyAndPaste: #div1=true;\nFormCompletionTime: 10s\n"
	d := data{
		WebsiteURL: "https://www.ravelin.com/",
		SessionID:  "123456",
		ResizeFrom: dimension{
			Width:  "200",
			Height: "200",
		},
		ResizeTo: dimension{
			Width:  "100",
			Height: "100",
		},
		CopyAndPaste:       copyAndPaste{"div1": true},
		FormCompletionTime: 10,
	}
	output := d.String()

	if output != expected {
		t.Errorf("expected output: \n%sgot:\n%s", expected, output)
	}
}

func TestSubmitHandler(t *testing.T) {
	testTable := make(map[string]fixture)
	testTable["TestIncorrectMethodError"] = fixture{
		body:       nil,
		httpMethod: http.MethodGet,
		statusCode: http.StatusMethodNotAllowed,
	}
	testTable["TestBodyReadError"] = fixture{
		body:       ioutil.NopCloser(mockReader{}),
		httpMethod: http.MethodPost,
		statusCode: http.StatusInternalServerError,
	}
	testTable["TestJSONUnmarshalError"] = fixture{
		body:       ioutil.NopCloser(bytes.NewBufferString("invalid JSON")),
		httpMethod: http.MethodPost,
		statusCode: http.StatusBadRequest,
	}
	testTable["TestSuccessfulSubmit"] = fixture{
		body:       ioutil.NopCloser(bytes.NewBufferString(`{"website_url":"lol","session_id":"1234","form_completion_time":76,"resize_from":{"width":"200","height":"250"},"resize_to":{"width":"100","height":"150"}}`)),
		httpMethod: http.MethodPost,
		statusCode: http.StatusNoContent,
	}
	testTable["TestContextDone"] = fixture{
		body:       ioutil.NopCloser(bytes.NewBufferString(`{}`)),
		httpMethod: http.MethodPost,
		statusCode: http.StatusRequestTimeout,
		ctx:        true,
	}

	for name, test := range testTable {
		t.Run(name, func(t *testing.T) {
			stream := make(chan data, 1)
			defer close(stream)

			handler := submitHandler(stream)
			req, err := http.NewRequest(test.httpMethod, "/submit", test.body)
			if err != nil {
				t.Fatal(err)
			}
			if test.ctx {
				stream <- data{}
				ctx, cancel := context.WithDeadline(req.Context(), time.Now().Add(-7*time.Hour))
				cancel()
				req = req.WithContext(ctx)
			}

			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			if rec.Code != test.statusCode {
				t.Errorf("expected status code: %d, got: %d", test.statusCode, rec.Code)
			}
		})
	}
}
