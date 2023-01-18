package repository

import (
	"context"
	"fmt"
	"simple-user-login/config"
	"simple-user-login/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAllUser : function to read all user
func GetAllUser() (users []*model.User, err error) {
	// mongo set connection
	client, ctx, cancel, err := config.ConnectMongo()
	if err != nil {
		panic(err)
	}
	defer config.CloseMongo(client, ctx, cancel)

	// collection access
	collection := client.Database("deals").Collection("user")

	cursor, err := collection.Find(ctx, options.Find())

	if err = cursor.All(context.TODO(), &users); err != nil {
		panic(err)
	}

	return
}

// GetUser : function to read a user
func GetUser(username string) (user *model.User, err error) {
	// mongo set connection
	client, ctx, cancel, err := config.ConnectMongo()
	if err != nil {
		panic(err)
	}
	defer config.CloseMongo(client, ctx, cancel)

	// collection access
	collection := client.Database("deals").Collection("user")

	// check user
	filter := bson.D{{"username", username}}
	collection.FindOne(context.TODO(), filter).Decode(&user)

	// return error if user doesn't exist
	if user == nil {
		return nil, fmt.Errorf("Unable to get data")
	}

	return
}
