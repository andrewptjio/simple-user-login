package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	UserName string `json:"username" bson:"username,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`

	Role string `json:"role" bson:"role,omitempty"`
	Name string `json:"name" bson:"name,omitempty"`
	Age  int    `json:"age" bson:"age,omitempty"`
}
