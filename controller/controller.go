package controller

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/gin-gonic/gin"
	"github.com/golang-pagination/database"
	"github.com/golang-pagination/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = database.OpenColletion()

func Populate(c *gin.Context) {

	ctx, cancel_ := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel_()
	for i := 0; i < 50; i++ {
		var product = model.Product{
			Title:       faker.Word(),
			Description: faker.Paragraph(),
			Image:       fmt.Sprintf("http://lorempixel.com/200/200?%s", faker.UUIDDigit()),
			Price:       rand.Intn(90) + 10,
		}
		collection.InsertOne(ctx, product)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func GetProducts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var products []model.Product
	cursor, _ := collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var product model.Product
		cursor.Decode(&product)
		products = append(products, product)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    products,
	})

}

func GetSingleProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var products []model.Product
	filter := bson.M{}
	findOptions := options.Find()
	if s := c.Query("s"); s != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"title": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
				{
					"description": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
			},
		}
	}

	if sort := c.Query("sort"); sort != "" {
		if sort == "asc" {
			findOptions.SetSort(bson.M{"price": 1})
		} else if sort == "desc" {
			findOptions.SetSort(bson.M{"price": -1})
		}
	}

	// total, _ := collection.CountDocuments(ctx, filter)

	// page, _ := strconv.Atoi(c.Query("page"))
	// var perPage int64 = 5
	// findOptions.SetSkip(int64(page-1) * perPage)
	// findOptions.SetLimit(perPage)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Println(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var product model.Product
		cursor.Decode(&product)
		products = append(products, product)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    products,
		// "total":     total,
		// "page":      page,
		// "last_page": math.Ceil(float64(total / perPage)),
	})
}
