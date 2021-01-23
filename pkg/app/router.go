package app

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"hamologist.com/emojify/pkg/router"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/emojify", router.EmojifyRouter())

	return r
}