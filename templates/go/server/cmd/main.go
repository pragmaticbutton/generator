package main

import (
	"net/http"

	"{{ .Name }}/internal/api"
)

func main() {
	http.ListenAndServe("localhost:8090", api.RegisterRoutes())
}
