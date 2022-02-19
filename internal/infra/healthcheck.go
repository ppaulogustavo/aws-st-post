package infra

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHealthCheck(r *chi.Mux) {
	r.Get("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("pong"))
	})
}
