package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TodoStatus string

const (
	StatusCompleted TodoStatus = "completed"
	StatusIncomplete  TodoStatus = "incomplete"	
)

type Todo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Status      TodoStatus         `bson:"status" json:"status"`
}