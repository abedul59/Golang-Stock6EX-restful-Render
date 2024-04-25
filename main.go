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
	CstockID        string `json:"cStockID" bson:"cStockID"`
	CstockName      string `json:"cStockName" bson:"cStockName"`
	CnewestSeason   string `json:"cNewestSeason" bson:"cNewestSeason"`
	CnewestRev      string `json:"cNewestRev" bson:"cNewestRev"`
	Csign1          string `json:"cSign1" bson:"cSign1"`
	Csign2          string `json:"cSign2" bson:"cSign2"`
	Csign3          string `json:"cSign3" bson:"cSign3"`
	Csign4          string `json:"cSign4" bson:"cSign4"`
	Csign5          string `json:"cSign5" bson:"cSign5"`
	Csign6          string `json:"cSign6" bson:"cSign6"`
	CaverageScore   string `json:"cAverageScore" bson:"cAverageScore"`
	ClossGain       string `json:"cLossGain" bson:"cLossGain"`
	CreateDate      string `json:"CreateDate" bson:"CreateDate"`
	CrevN           string `json:"cRevN" bson:"cRevN"`
	Crev            string `json:"cRev" bson:"cRev"`
	Ca1N            string `json:"ca1N" bson:"ca1N"`
	Ca2N            string `json:"ca2N" bson:"ca2N"`
	Ca3N            string `json:"ca3N" bson:"ca3N"`
	Ca4N            string `json:"ca4N" bson:"ca4N"`
	Ca5N            string `json:"ca5N" bson:"ca5N"`
	Ca6N            string `json:"ca6N" bson:"ca6N"`
	Ca7N            string `json:"ca7N" bson:"ca7N"`
	Cna1            string
	Cna2            string
	Cna3            string
	Cna4            string
	Cna5            string
	Cna6            string
	Cna7            string
	Cna9            string
	Cna10           string
	CnewestRevMonth string
	CluX            string
	CnluXMoM        string
	CProfitN        string
	CProfit         string
	Cb1N            string
	Cb2N            string
	Cb3N            string
	Cb4N            string
	Cb5N            string
	Cb6N            string
	Cb7N            string
	Cb8N            string
	Cb1             string
	Cb2             string
	Cb3             string
	Cb4             string
	Cb5             string
	Cb6             string
	Cb7             string
	Cb8             string
	Cb9             string
	Cb10            string
	Cb10P           string
	CInvTON         string
	CInvTO          string
	Ce1N            string
	Ce2N            string
	Ce3N            string
	Ce4N            string
	Ce5N            string
	Ce6N            string
	Ce7N            string
	Ce8N            string
	Ce1             string
	Ce2             string
	Ce3             string
	Ce4             string
	Ce5             string
	Ce6             string
	Ce7             string
	Ce8             string
	CnewestFinQ     string
	CNetIncomeN     string
	CNetIncome      string
	Cc1N            string
	Cc2N            string
	Cc3N            string
	Cc4N            string
	Cc5N            string
	Cc6N            string
	Cc7N            string
	Cc8N            string
	Cc1             string
	Cc2             string
	Cc3             string
	Cc4             string
	Cc5             string
	Cc6             string
	Cc7             string
	Cc8             string
	Cc9             string
	Cc10            string
	Cc11            string
	Cpc9            string
	Cpc10           string
	Cpc11           string
	CEPSN           string
	CEPS            string
	CD1N            string
	CD2N            string
	CD3N            string
	CD4N            string
	CD5N            string
	CD6N            string
	CD7N            string
	CD8N            string
	CD1             string
	CD2             string
	CD3             string
	CD4             string
	CD5             string
	CD6             string
	CD7             string
	CD8             string
	CCashFlowN      string
	CCashFlow       string
	Cf1N            string
	Cf2N            string
	Cf3N            string
	Cf4N            string
	Cf5N            string
	Cf6N            string
	Cf7N            string
	Cf8N            string
	Cf1             string
	Cf2             string
	Cf3             string
	Cf4             string
	Cf5             string
	Cf6             string
	Cf7             string
	Cf8             string
	Cf9             string
	Cf10            string
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
	err := collection.FindOne(context.Background(), bson.M{"CstockID": name}).Decode(&item)
	if err != nil {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(200, item)
}
