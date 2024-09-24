package services

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"todo-app/model"
	"todo-app/structs"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodos() ([]structs.Todo, error) {
	var todoList []structs.Todo

	res, err := model.GetTodos()

	if err != nil {
		log.Fatal("Could not find todo items", err)
	}

	for res.Next(context.TODO()) {
		var todo structs.Todo
		err := res.Decode(&todo)

		if err != nil {
			log.Fatal("Could not decode todo item", err)
			continue
		}

		todoList = append(todoList, todo)
	}

	return todoList, err;
}

func CreateTodo(payload io.ReadCloser) (primitive.ObjectID){
	var todo structs.Todo

	decoder := json.NewDecoder(payload)

			err := decoder.Decode(&todo)
			if err != nil {
				log.Fatal(err)
			}

			res, err := model.CreateTodo(todo)
	
			if err != nil {
				log.Fatal("Could not insert todo item", err)
			}

			return res.InsertedID.(primitive.ObjectID)
}

func DeleteTodo(todoId string) (int64, error){

	id, err := primitive.ObjectIDFromHex(todoId)

	if err != nil {
		log.Fatal("Could not convert todoId to ObjectID", err)
	}

		res, err := model.DeleteTodo(id)

		if err != nil || res.DeletedCount == 0 {
			log.Fatal("Could not delete todo item", err)
		}
	
		return res.DeletedCount, err
}

func UpdateTodo(todoId string, payload io.ReadCloser) int64 {
	id, err := primitive.ObjectIDFromHex(todoId)

	
	if err != nil {
		log.Fatal("Could not convert todoId to ObjectID", err)
	}

	var updatedTodo structs.UpdateTodo

	decoder := json.NewDecoder(payload)
	err = decoder.Decode(&updatedTodo)
	if err != nil {
		log.Fatal("Couldn't decode request body", err)
	}

	res, err := model.UpdateTodo(id, updatedTodo)
	if err != nil {
		log.Fatal("Could not update todo item", err)
	}

	return res.ModifiedCount
}

func UpdateAll() int64 {
	res, err := model.UpdateAll()

	if err != nil {
		log.Fatal("Could not update all todo items", err)
	}

	return res.ModifiedCount
}

