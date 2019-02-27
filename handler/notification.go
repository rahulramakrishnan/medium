package handler

import (
	"io"
	"net/http"
)

type (
	// Notification is an http handler interface for
	// sending and receiving notifications.
	Notification interface {
		SendSMS(w http.ResponseWriter, req *http.Request)
	}
	notification struct{}
)

func (n *notification) SendSMS(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}
