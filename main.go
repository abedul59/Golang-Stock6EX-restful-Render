package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Item struct {
	CstockID      string `json:"cStockID" bson:"cStockID"`
	CstockName    string `json:"cStockName" bson:"cStockName"`
	CnewestSeason string `json:"cNewestSeason" bson:"cNewestSeason"`
	CnewestRev    string `json:"cNewestRev" bson:"cNewestRev"`
	Csign1        string `json:"cSign1" bson:"cSign1"`
	Csign2        string `json:"cSign2" bson:"cSign2"`
	Csign3        string `json:"cSign3" bson:"cSign3"`
	Csign4        string `json:"cSign4" bson:"cSign4"`
	Csign5        string `json:"cSign5" bson:"cSign5"`
	Csign6        string `json:"cSign6" bson:"cSign6"`
	CaverageScore string `json:"cAverageScore" bson:"cAverageScore"`
	ClossGain     string `json:"cLossGain" bson:"cLossGain"`
	CreateDate    string `json:"CreateDate" bson:"CreateDate"`
}

var collection *mongo.Collection

func main() {
	// Set up MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb+srv://pyfbsdk59:NHd4ZEVmHONPZiYD@mongodb-restful.5xgpkpw.mongodb.net/?retryWrites=true&w=majority&appName=mongodb-restful")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Set up a new gin router
	r := gin.Default()

	// Set up the MongoDB collection
	collection = client.Database("test").Collection("s6r202403")

	// API endpoints for CRUD operations

	r.GET("/s6r202403/:name", readItem)

	// Start the server
	r.Run(":8080")
}

func readItem(c *gin.Context) {
	name := c.Param("name")

	var item Item
	err := collection.FindOne(context.Background(), bson.M{"cStockID": name}).Decode(&item)
	if err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(200, item)
}
