package backend

import (
	"context"
	"log"
	"os"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/AleksanderWWW/go-webapp/resources"
)

var allProducts []resources.Product = []resources.Product{
	{Id: "C1", Name: "Chair", QuantityInStock: 200, Price: 7850},
	{Id: "C2", Name: "Chair", QuantityInStock: 100, Price: 9850},
	{Id: "B1", Name: "Bed", QuantityInStock: 900, Price: 23378},
}

func RetrieveAll() []resources.Product {
	return allProducts
}

func RetrieveByID(id string) resources.Product {
	for _, p := range allProducts {
		if p.Id == id {
			return p
		}
	}
	return resources.Product{}
}

func ConnectToMongo(db string, collection_name string) mongo.Collection {
	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return *client.Database(db).Collection(collection_name)
}

func CloseMongoConnection(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func InsertProduct(coll mongo.Collection, p resources.Product) {
	result, err := coll.InsertOne(context.TODO(), p)
	if err != nil {
		panic(err)
	}
	log.Printf("Inserted document with _id: %v\n", result.InsertedID)
}
