package main

import (
	"hamologist.com/emojify/pkg/app"
	"net/http"
)

func main() {
	router := app.NewRouter()
	http.ListenAndServe(":3000", router)
}
