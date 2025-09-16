package service

import (
	"github.com/Sambit99/Basic-HRMS-GO-MongoDB/pkg/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var db *mongo.Collection

func init() {
	config.ConnectDB()
	db = config.GetDB()
}
