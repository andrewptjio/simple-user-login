package user

import (
	"context"
	"fmt"
	"simple-user-login/config"
	"simple-user-login/model"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// save : function to create a user
func save(r *requestCreate) (user *model.User, err error) {
	// mongo set connection
	client, ctx, cancel, err := config.ConnectMongo()
	if err != nil {
		panic(err)
	}
	defer config.CloseMongo(client, ctx, cancel)

	// collection access
	collection := client.Database("deals").Collection("user")

	// save required data
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(r.Password), 14)
	m := model.User{
		UserName: r.UserName,
		Password: string(hashPassword),
		Role:     r.Role,
		Name:     r.Name,
		Age:      r.Age,
	}

	// insert data
	result, err := collection.InsertOne(ctx, m)
	if err != nil {
		return nil, fmt.Errorf("Unable to save data")
	}

	// return inserted data
	data := collection.FindOne(context.Background(), bson.M{"_id": result.InsertedID})
	data.Decode(&user)

	return user, err
}

// update : function to update a user
func update(r *requestUpdate) (user *model.User, err error) {
	// mongo set connection
	client, ctx, cancel, err := config.ConnectMongo()
	if err != nil {
		panic(err)
	}
	defer config.CloseMongo(client, ctx, cancel)

	// collection access
	collection := client.Database("deals").Collection("user")

	// filter used
	filter := bson.D{{"username", r.User.UserName}}
	update := bson.D{{"$set", bson.D{{"role", r.Role}, {"name", r.Name}, {"age", r.Age}}}}

	// update data
	if _, err := collection.UpdateOne(context.TODO(), filter, update); err != nil {
		return nil, fmt.Errorf("Unable to update data")
	}

	// return updated data
	data := collection.FindOne(context.Background(), filter)
	data.Decode(&user)

	return user, err
}

// delete : function to delete a user
func delete(r *requestDelete) (user *model.User, err error) {
	// mongo set connection
	client, ctx, cancel, err := config.ConnectMongo()
	if err != nil {
		panic(err)
	}
	defer config.CloseMongo(client, ctx, cancel)

	// collection access
	collection := client.Database("deals").Collection("user")

	// filter used
	filter := bson.D{{"username", r.User.UserName}}

	// update data
	if err := collection.FindOneAndDelete(context.TODO(), filter).Decode(&user); err != nil {
		return nil, fmt.Errorf("Unable to delete data")
	}

	return user, err
}
