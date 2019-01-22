package handler

import (
	"io"
	"net/http"
)

// Ping is ping request handler.
func Ping(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "OK")
	if err != nil {
		return
	}
}

