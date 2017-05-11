package main

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockReader struct{}

func (m mockReader) Read(p []byte) (int, error) {
	return 0, errors.New("read error")
}

type fixture struct {
	body       io.Reader
	httpMethod string
	statusCode int
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

	for name, test := range testTable {
		t.Run(name, func(t *testing.T) {
			handler := http.HandlerFunc(submitHandler)
			req, err := http.NewRequest(test.httpMethod, "/submit", test.body)
			if err != nil {
				t.Fatal(err)
			}
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			if rec.Code != test.statusCode {
				t.Errorf("expected status code: %d, got: %d", test.statusCode, rec.Code)
			}
		})
	}
}
