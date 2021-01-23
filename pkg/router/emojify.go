package router

import (
	"fmt"
	"github.com/go-chi/chi"
	"hamologist.com/emojify/pkg/model"
	"hamologist.com/emojify/pkg/transformer"
	"net/http"
)

func EmojifyRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", func(writer http.ResponseWriter, request *http.Request) {
		resp, _ := transformer.EmojifyTransformer(model.EmojifyPayload{Message: ""})
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte(fmt.Sprintf("{\"message\": \"%s\"}", resp.Message)))
	})

	return router
}