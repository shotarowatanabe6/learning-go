package handler

import (
	"fmt"
	"net/http"
)

func NewServer(port int) *http.Server {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: nil,
	}
	http.HandleFunc("/", helloWorld)

	return server
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
