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

	//CrevN           []string `json:"cRevN" bson:"cRevN"`
	//Crev            []string `json:"cRev" bson:"cRev"`
	Ca1N            string `json:"ca1N" bson:"ca1N"`
	Ca2N            string `json:"ca2N" bson:"ca2N"`
	Ca3N            string `json:"ca3N" bson:"ca3N"`
	Ca4N            string `json:"ca4N" bson:"ca4N"`
	Ca5N            string `json:"ca5N" bson:"ca5N"`
	Ca6N            string `json:"ca6N" bson:"ca6N"`
	Ca7N            string `json:"ca7N" bson:"ca7N"`
	Cna1            string `json:"cna1" bson:"cna1"`
	Cna2            string `json:"cna2" bson:"cna2"`
	Cna3            string `json:"cna3" bson:"cna3"`
	Cna4            string `json:"cna4" bson:"cna4"`
	Cna5            string `json:"cna5" bson:"cna5"`
	Cna6            string `json:"cna6" bson:"cna6"`
	Cna7            string `json:"cna7" bson:"cna7"`
	Cna9            string `json:"cna9" bson:"cna9"`
	Cna10           string `json:"cna10" bson:"cna10"`
	CnewestRevMonth string `json:"cnewest_Rev_month" bson:"cnewest_Rev_month"`
	CluX            string `json:"cluX" bson:"cluX"`
	CnluXMoM        string `json:"ccnluX_MoM" bson:"cnluX_MoM"`
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
