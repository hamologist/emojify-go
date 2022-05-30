package router

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/hamologist/emojify-go/pkg/model"
	"github.com/hamologist/emojify-go/pkg/transformer"
	"net/http"
)

func EmojifyRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", func(writer http.ResponseWriter, request *http.Request) {
		payload := model.EmojifyPayload{}
		err := json.NewDecoder(request.Body).Decode(&payload)
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(500)
			writer.Write([]byte("{\"error\": \"Invalid payload\"}"))
			return
		}

		if payload.Message == "" {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(500)
			writer.Write([]byte("{\"error\": \"Payload contains an empty or missing message field\"}"))
			return
		}

		resp, _ := transformer.EmojifyTransformer(payload)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte(fmt.Sprintf("{\"message\": \"%s\"}", resp.Message)))
	})

	return router
}
