package controller

import (
	"encoding/json"
	"net/http"
	"todo-app/services"
	htmlGenerator "todo-app/utils"

	"github.com/go-chi/chi"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {

	todoList, err := services.GetTodos()

	if err != nil {
		http.Error(w, "Could not find todo items", http.StatusInternalServerError)
		return
	}
	
	htmlGenerator.HtmlGenerator(w, todoList)
}

func CreateTodo(w http.ResponseWriter, r *http.Request){

	res := services.CreateTodo(r.Body)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Todo item inserted successfully",
		"id": res,
	})
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "id")

		res, err := services.DeleteTodo(todoId)

		if err != nil {
			http.Error(w, "Could not delete todo item", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Todo item deleted successfully",
			"count": res,
			"id": todoId,
		})
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "id")


		res := services.UpdateTodo(todoId, r.Body)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Todo item updated successfully",
			"count": res,
			"id": todoId,
		})
}

func UpdateAll(w http.ResponseWriter, r *http.Request) {
	res := services.UpdateAll()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "All todo items updated successfully",
		"count": res,
	})
}
