package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Employee struct {
	ID     bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string        `json:"name" bson:"name"`
	Salary float64       `json:"salary" bson:"salary"`
	Age    int64         `json:"age,omitempty" bson:"age,omitempty"`
}

type NewEmployeeDto struct {
	Name   string  `json:"name" bson:"name"`
	Salary float64 `json:"salary" bson:"salary"`
	Age    int64   `json:"age,omitempty" bson:"age,omitempty"`
}

type UpdateEmployeeDto struct {
	ID     bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string        `json:"name" bson:"name"`
	Salary float64       `json:"salary" bson:"salary"`
	Age    int64         `json:"age,omitempty" bson:"age,omitempty"`
}
