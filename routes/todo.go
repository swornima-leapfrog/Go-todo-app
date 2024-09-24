package routes

import (
	"net/http"

	"todo-app/controller"

	"github.com/go-chi/chi"
)

func TodoRoute() http.Handler{

	rg := chi.NewRouter()

	rg.Group(func(r chi.Router) {
		r.Post("/", controller.CreateTodo)
		r.Put("/all", controller.UpdateAll)
		r.Put("/{id}", controller.UpdateTodo)
		r.Delete("/{id}", controller.DeleteTodo)
	})

	return rg

}
