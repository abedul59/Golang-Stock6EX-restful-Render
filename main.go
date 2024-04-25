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

	CSign1        string
	CSign2        string
	CSign3        string
	CSign4        string
	CSign5        string
	CSign6        string
	CAverageScore string
	CLossGain     string
	CreateDate    string

	CRevN string
	CRev  string
	Ca1N  string
	Ca2N  string
	Ca3N  string
	Ca4N  string
	Ca5N  string
	Ca6N  string
	Ca7N  string

	Cna1              string
	Cna2              string
	Cna3              string
	Cna4              string
	Cna5              string
	Cna6              string
	Cna7              string
	Cna9              string
	Cna10             string
	Cnewest_Rev_month string
	CluX              string
	CnluX_MoM         string

	CProfitN string
	CProfit  string
	Cb1N     string
	Cb2N     string
	Cb3N     string
	Cb4N     string
	Cb5N     string
	Cb6N     string
	Cb7N     string
	Cb8N     string

	Cb1           string
	Cb2           string
	Cb3           string
	Cb4           string
	Cb5           string
	Cb6           string
	Cb7           string
	Cb8           string
	Cb9           string
	Cb10          string
	Cb10p         string
	CInvTON       string
	CInvTO        string
	Ce1N          string
	Ce2N          string
	Ce3N          string
	Ce4N          string
	Ce5N          string
	Ce6N          string
	Ce7N          string
	Ce8N          string
	Ce1           string
	Ce2           string
	Ce3           string
	Ce4           string
	Ce5           string
	Ce6           string
	Ce7           string
	Ce8           string
	Cnewest_Fin_Q string

	CNetIncomeN string
	CNetIncome  string
	Cc1N        string
	Cc2N        string
	Cc3N        string
	Cc4N        string
	Cc5N        string
	Cc6N        string
	Cc7N        string
	Cc8N        string
	Cc1         string
	Cc2         string
	Cc3         string
	Cc4         string
	Cc5         string
	Cc6         string
	Cc7         string
	Cc8         string
	Cc9         string
	Cc10        string
	Cc11        string
	Cpc9        string
	Cpc10       string
	Cpc11       string

	CEPSN string
	CEPS  string
	Cd1N  string
	Cd2N  string
	Cd3N  string
	Cd4N  string
	Cd5N  string
	Cd6N  string
	Cd7N  string
	Cd8N  string
	Cd1   string
	Cd2   string
	Cd3   string
	Cd4   string
	Cd5   string
	Cd6   string
	Cd7   string
	Cd8   string

	CCashFlowN string
	CCashFlow  string
	Cf1N       string
	Cf2N       string
	Cf3N       string
	Cf4N       string
	Cf5N       string
	Cf6N       string
	Cf7N       string
	Cf8N       string
	Cf1        string
	Cf2        string
	Cf3        string
	Cf4        string
	Cf5        string
	Cf6        string
	Cf7        string
	Cf8        string
	Cf9        string
	Cf10       string
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
