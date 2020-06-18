package global

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DB holds datbase Connection
var DB mongo.Database

func connectToDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal("Error connect to DB: ", err.Error())
	}
	DB = *client.Database(dbname)

}

//NewDBContext returns a new contact according to app performance
func NewDBContext(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d*performance/100)
}

//ConnectToTestDB overwrites database with test database
func ConnectToTestDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal("Error connect to DB: ", err.Error())
	}
	DB = *client.Database(dbname + "_test")
}
