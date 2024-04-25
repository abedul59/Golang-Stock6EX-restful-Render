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
	cStockID      string
	CStockName    string
	CNewestSeason string
	CNewestRev    string

	cSign1        string
	cSign2        string
	cSign3        string
	cSign4        string
	cSign5        string
	cSign6        string
	cAverageScore string
	cLossGain     string
	CreateDate    string

	cRevN string
	cRev  string
	ca1N  string
	ca2N  string
	ca3N  string
	ca4N  string
	ca5N  string
	ca6N  string
	ca7N  string

	cna1              string
	cna2              string
	cna3              string
	cna4              string
	cna5              string
	cna6              string
	cna7              string
	cna9              string
	cna10             string
	cnewest_Rev_month string
	cluX              string
	cnluX_MoM         string

	cProfitN string
	cProfit  string
	cb1N     string
	cb2N     string
	cb3N     string
	cb4N     string
	cb5N     string
	cb6N     string
	cb7N     string
	cb8N     string

	cb1           string
	cb2           string
	cb3           string
	cb4           string
	cb5           string
	cb6           string
	cb7           string
	cb8           string
	cb9           string
	cb10          string
	cb10p         string
	cInvTON       string
	cInvTO        string
	ce1N          string
	ce2N          string
	ce3N          string
	ce4N          string
	ce5N          string
	ce6N          string
	ce7N          string
	ce8N          string
	ce1           string
	ce2           string
	ce3           string
	ce4           string
	ce5           string
	ce6           string
	ce7           string
	ce8           string
	cnewest_Fin_Q string

	cNetIncomeN string
	cNetIncome  string
	cc1N        string
	cc2N        string
	cc3N        string
	cc4N        string
	cc5N        string
	cc6N        string
	cc7N        string
	cc8N        string
	cc1         string
	cc2         string
	cc3         string
	cc4         string
	cc5         string
	cc6         string
	cc7         string
	cc8         string
	cc9         string
	cc10        string
	cc11        string
	cpc9        string
	cpc10       string
	cpc11       string

	cEPSN string
	cEPS  string
	cd1N  string
	cd2N  string
	cd3N  string
	cd4N  string
	cd5N  string
	cd6N  string
	cd7N  string
	cd8N  string
	cd1   string
	cd2   string
	cd3   string
	cd4   string
	cd5   string
	cd6   string
	cd7   string
	cd8   string

	cCashFlowN string
	cCashFlow  string
	cf1N       string
	cf2N       string
	cf3N       string
	cf4N       string
	cf5N       string
	cf6N       string
	cf7N       string
	cf8N       string
	cf1        string
	cf2        string
	cf3        string
	cf4        string
	cf5        string
	cf6        string
	cf7        string
	cf8        string
	cf9        string
	cf10       string
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
