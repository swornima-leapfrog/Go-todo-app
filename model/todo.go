package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"todo-app/db"
	"todo-app/structs"
)

func GetTodos() (*mongo.Cursor, error) {
	res, err := db.TodoDetails.Find(context.TODO(), bson.D{})

	return res, err

}

func CreateTodo(todo structs.Todo) (*mongo.InsertOneResult, error) {
	res, err := db.TodoDetails.InsertOne(context.TODO(), todo)

	return res, err
}

func DeleteTodo(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	res, err := db.TodoDetails.DeleteOne(context.TODO(), bson.M{"_id": id})

	return res, err
}

func UpdateTodo(id primitive.ObjectID, todo structs.UpdateTodo) (*mongo.UpdateResult, error) {
	// Initialize the update query
	update := bson.M{
		"$set": bson.M{},
	}

	// Conditionally add the "title" field only if it's not empty
	if todo.Title != "" {
		update["$set"].(bson.M)["title"] = todo.Title
	}

	// Always update the "done" field
	update["$set"].(bson.M)["done"] = todo.Done

	// Perform the update in MongoDB
	res, err := db.TodoDetails.UpdateOne(context.TODO(), bson.M{"_id": id}, update)

	return res, err
}

func UpdateAll() (*mongo.UpdateResult, error) {
	res, err := db.TodoDetails.UpdateMany(context.TODO(), bson.D{}, bson.M{"$set": bson.M{"done": true}})

	return res, err
}
