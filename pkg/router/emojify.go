package router

import (
	"github.com/go-chi/chi"
	"net/http"
)

func EmojifyRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte("{\"message\": \"Under Development\"}"))
	})

	return router
}