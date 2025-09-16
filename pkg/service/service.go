package service

import (
	"context"
	"log"

	"github.com/Sambit99/Basic-HRMS-GO-MongoDB/pkg/config"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var db *mongo.Collection
var ctx = context.TODO()

func init() {
	config.ConnectDB()
	db = config.GetDB()
}

func GetEmployees() []bson.M {

	cursor, err := db.Find(ctx, bson.M{}, nil)

	if err != nil {
		panic(err)
	}

	defer cursor.Close(ctx)

	var employees []bson.M

	for cursor.Next(ctx) {
		var emp bson.M

		if err := cursor.Decode(&emp); err != nil {
			log.Fatal("Error while decoding Employee", err.Error())
		}

		employees = append(employees, emp)
	}

	return employees
}
