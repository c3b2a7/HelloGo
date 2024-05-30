package http

import "net/http"

func NewHttpServer() {
	mux := http.NewServeMux()
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
