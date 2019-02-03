package handler

import (
	"io"
	"net/http"

	"github.com/gmidorii/api-design/wire/provider"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	hello, err := provider.InitHelloApp()
	if err != nil {
		return
	}
	message, err := hello.GetMessage("id")
	if err != nil {
		return
	}

	io.WriteString(w, message)
}
