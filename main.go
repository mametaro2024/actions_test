package main

import (
	"net/http"
)

func main() {

	handler := HelloHandler{}

	server := http.Server{
		Addr:    ":8000",
		Handler: &handler,
	}
	server.ListenAndServe()
}
