package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var unknownMessage = errors.New("unknown message number")

type message struct {
	id    int
	value string
}

type messagesHandler struct {
	messages []message
}

func (h *messagesHandler) lookup(id int) (*message, error) {
	for _, m := range h.messages {
		if id == m.id {
			return &m, nil
		}
	}
	return nil, unknownMessage
}

func newMessageHandler() *messagesHandler {
	return &messagesHandler{
		messages: []message{
			{id: 1, value: "test1"},
			{id: 2, value: "test2"},
			{id: 3, value: "test3"},
			{id: 4, value: "test4"},
			{id: 5, value: "test5"},
		},
	}
}

func (h *messagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id, err := strconv.Atoi(r.URL.Path[len("/messages/"):])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		m, err := h.lookup(id)
		if err == unknownMessage {
			http.Error(w, unknownMessage.Error(), http.StatusNotFound)
			return
		}

		fmt.Fprintf(w, m.value)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
