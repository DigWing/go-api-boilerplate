package db

import (
    "context"
    "os"
    "time"
    "github.com/mongodb/mongo-go-driver/mongo"
)

var MongoDB *mongo.Database

func Connect() {
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    client, _ := mongo.Connect(ctx, os.Getenv("MONGO_HOST"))

    MongoDB = client.Database(os.Getenv("MONGO_DB"))
}
