package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"simple-user-login/model"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/crypto/bcrypt"
)

var (
	AppPort         string
	MongoURI        string
	JWTSecretKey    string
	DefaultPassword string
)

// InitConfig: function to load and read env
func InitConfig() {
	// init env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// read env
	AppPort = os.Getenv("APP_PORT")
	MongoURI = os.Getenv("MONGO_URI")
	JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	DefaultPassword = os.Getenv("DEFAULT_PASSWORD")
}

// MongoConnectinoTest : function to test connection with mongo
func MongoConnectionTest() {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoURI))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	// create an admin and a user
	mongoSeed(client)

	fmt.Println("*** Mongo successfully connected and pinged. ***")
}

func mongoSeed(client *mongo.Client) {
	collection := client.Database("deals").Collection("user")

	var (
		admin *model.User
		user  *model.User
	)

	// get admin
	data := collection.FindOne(context.Background(), bson.M{"username": "admin_sample"})

	// create admin if doesn't exist
	data.Decode(&admin)
	if admin == nil {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(DefaultPassword), 14)

		m := model.User{
			UserName: "admin_sample",
			Password: string(hashPassword),
			Role:     "admin",
			Name:     "admin",
			Age:      99,
		}

		// insert data
		collection.InsertOne(context.TODO(), m)
	}

	// get user
	data = collection.FindOne(context.Background(), bson.M{"username": "user_sample"})

	// create a user if doesn't exist
	data.Decode(&user)
	if user == nil {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(DefaultPassword), 14)

		m := model.User{
			UserName: "user_sample",
			Password: string(hashPassword),
			Role:     "user",
			Name:     "user",
			Age:      1,
		}

		// insert data
		collection.InsertOne(context.TODO(), m)
	}

}

// ConnectMongo : function to create a connection with mongo
func ConnectMongo() (*mongo.Client, context.Context, context.CancelFunc, error) {
	// ctx will be used to set deadline for process
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoURI))
	return client, ctx, cancel, err
}

// CloseMongo : function to close the connection with mongo
func CloseMongo(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
