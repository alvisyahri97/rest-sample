package main

import (
	"RESTAPISample/handler/user"
	"RESTAPISample/repo/users"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	handler := user.NewUserHandler(users.NewUser())
	router := chi.NewRouter()

	// Initialize router.
	router.Route("/user/{userid}", func(r chi.Router) {
		r.Get("/", handler.GetByID)
		r.Put("/", handler.UpdateByID)
		r.Delete("/", handler.DeleteByID)
	})
	router.Post("/user", handler.InsertUser)

	log.Print("Starting http server at port 8000")

	// Initialize http server.
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		panic(err)
	}
}
