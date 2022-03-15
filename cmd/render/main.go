package main

import (
	"log"
	"net/http"

	"github.com/alvisyahri97/rest-sample/handler/render"

	"github.com/go-chi/chi"
)

func main() {
	handler := render.NewRenderHandler()
	router := chi.NewRouter()

	// Initialize router.
	router.Get("/{userid}", handler.GetByID)

	log.Print("Starting http server at port 8000")

	// Initialize http server.
	err := http.ListenAndServe(":8001", router)
	if err != nil {
		panic(err)
	}
}
