package helpers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	model "github.com/Mukunth-arya/Ticketbooking/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Data_env() string {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Errorf("Here having key value pair error map[string]string")
	}

	return os.Getenv("password")

}

var (
	password = os.Getenv(Data_env())
	datas    model.Models
)

const (
	db_url      = "mongodb+srv://mukunth:<password>@mycluster.jptcn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	db_database = "Ticketbookdb"
	db_col      = "Ticketbookcl"
)

// Mongodb connction stage
func Mongoconnect() *mongo.Collection {
	Client := options.Client().ApplyURI(db_url)

	Connect, err := mongo.Connect(context.TODO(), Client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection is successfull")
	defer fmt.Println("Collection instance is ready")
	return Connect.Database(db_url).Collection(db_col)

}
