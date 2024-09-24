package htmlGenerator

import (
	"fmt"
	"html/template"
	"net/http"
	"todo-app/structs"
)

func HtmlGenerator(w http.ResponseWriter, todoList []structs.Todo) {

	t, err := template.ParseFiles("static/index.html")

	if err != nil {
		http.Error(w, "Unable to parse template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	err = t.Execute(w, todoList)

	if err != nil {
		http.Error(w, "Unable to generate HTML", http.StatusInternalServerError)
		fmt.Println("Error executing template:",err)
	}
}
