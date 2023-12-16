package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_messagesHandler_ServeHTTP(t *testing.T) {
	cases := []struct {
		name               string
		method             string
		endpoint           string
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:               "OK",
			method:             http.MethodGet,
			endpoint:           "/messages/1",
			expectedStatusCode: http.StatusOK,
			expectedBody:       "test1",
		},
		{
			name:               "BadRequest",
			method:             http.MethodGet,
			endpoint:           "/messages/",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "",
		},
		{
			name:               "MethodNotAllowed",
			method:             http.MethodPost,
			endpoint:           "/messages/",
			expectedStatusCode: http.StatusMethodNotAllowed,
			expectedBody:       "",
		},
	}

	for _, c := range cases {
		r, _ := http.NewRequest(c.method, c.endpoint, nil)
		w := httptest.NewRecorder()
		h := newMessageHandler()
		h.ServeHTTP(w, r)

		if w.Code != c.expectedStatusCode {
			t.Fatalf("extected %d but get %d for status code", c.expectedStatusCode, w.Code)
		}

		body := strings.TrimSpace(w.Body.String())
		if body != c.expectedBody {
			t.Fatalf("extected %s but get %s for body", c.expectedBody, w.Body)
		}
	}
}
