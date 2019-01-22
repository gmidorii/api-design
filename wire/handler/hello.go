package handler

import (
	"github.com/gmidorii/api-design/wire/provider"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	hello, err := provider.InitHelloApp()
	if err != nil {
		return
	}
	fmt.Println(hello)
}