package router

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"hamologist.com/emojify/pkg/model"
	"hamologist.com/emojify/pkg/transformer"
	"net/http"
)

func EmojifyRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", func(writer http.ResponseWriter, request *http.Request) {
		payload := model.EmojifyPayload{}
		err := json.NewDecoder(request.Body).Decode(&payload)
		if err != nil {
			http.Error(writer, http.StatusText(500), 500)
			return
		}

		resp, _ := transformer.EmojifyTransformer(payload)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte(fmt.Sprintf("{\"message\": \"%s\"}", resp.Message)))
	})

	return router
}