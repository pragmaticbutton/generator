package api

import "net/http"

func RegisterRoutes() http.Handler {

	h := &MyHandler{}

	mux := http.NewServeMux()
	mux.Handle("/", http.NotFoundHandler())
	mux.HandleFunc("GET /{$}", h.Get)
	mux.HandleFunc("GET /{id}/{$}", h.GetWithID)
	mux.HandleFunc("POST /{$}", h.Post)

	return mux

}
