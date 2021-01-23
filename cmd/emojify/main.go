package main

import (
	"github.com/hamologist/emojify/pkg/app"
	"net/http"
)

func main() {
	router := app.NewRouter()
	http.ListenAndServe(":3000", router)
}
