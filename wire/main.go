package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gmidorii/api-design/wire/handler"
)

func run(port string) error {
	http.HandleFunc("/ping", handler.Ping)
	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func main() {
	port := flag.String("p", "8080", "port")
	flag.Parse()

	if err := run(*port); err != nil {
		log.Println(err)
	}
}
