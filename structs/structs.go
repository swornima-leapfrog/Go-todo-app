package structs

type Todo struct {
	Id    string `bson:"_id,omitempty"`
	Title string `bson:"title"`
	Done  bool   `bson:"done"`
}

type UpdateTodo struct {
	Title string `bson:"title,omitempty"`
	Done  bool   `bson:"done,omitempty"`
}
