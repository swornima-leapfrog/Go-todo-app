package main

import (
	"log"
	"net/http"
	"todo-app/controller"
	"todo-app/routes"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
)

var todoDetails *mongo.Collection

func main() {

	router := chi.NewRouter()

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	router.Get("/", controller.GetTodos)
	router.Mount("/", routes.TodoRoute())
	
	log.Fatal(http.ListenAndServe(":3000", router))
}
