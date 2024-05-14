package api

import (
	"fmt"
	"io"
	"net/http"
)

type MyHandler struct {
}

func (h *MyHandler) Get(rw http.ResponseWriter, r *http.Request) {

	rw.Write([]byte("Hello from GET!"))
}

func (h *MyHandler) GetWithID(rw http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	rw.Write([]byte("Hello from GET, requested id is: " + id))
}

func (h *MyHandler) Post(rw http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "reading body failed", http.StatusInternalServerError)
		return
	}
	fmt.Println("this is printed from post")

	rw.Write(body)
}
