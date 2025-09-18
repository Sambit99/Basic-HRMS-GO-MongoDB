package service

import (
	"context"
	"fmt"
	"log"

	"github.com/Sambit99/Basic-HRMS-GO-MongoDB/pkg/config"
	"github.com/Sambit99/Basic-HRMS-GO-MongoDB/pkg/model"
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

func GetEmployee(empId string) model.Employee {

	ID, err := bson.ObjectIDFromHex(empId)

	if err != nil {
		log.Fatal("Error while parsing Employee ID", err.Error())
	}

	var employee model.Employee

	filter := bson.M{"_id": ID}

	err = db.FindOne(ctx, filter, nil).Decode(&employee)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No Records found")
		} else {
			log.Fatal("Error while finding Employee by id", err.Error())
		}
	}

	return employee
}

func CreateEmployee(emp model.NewEmployeeDto) string {

	newEmployee, err := db.InsertOne(ctx, emp, nil)

	if err != nil {
		log.Fatal("Error while creating user")
	}

	oid, ok := newEmployee.InsertedID.(bson.ObjectID)
	if !ok {
		log.Fatal("Inserted ID is not of type ObjectID")
	}

	return oid.Hex()
}

func DeleteEmployee(id string) bool {

	parsedId, err := bson.ObjectIDFromHex(id)

	if err != nil {
		log.Fatal("Error while parsing Employee ID", err.Error())
	}

	filter := bson.M{"_id": parsedId}

	result, err := db.DeleteOne(ctx, filter, nil)

	if err != nil {
		log.Fatal(err)
	}

	return result.DeletedCount > 0
}

func UpdateEmployee(updatedEmployee model.UpdateEmployeeDto) bool {
	filter := bson.M{"_id": updatedEmployee.ID}

	strId := updatedEmployee.ID.Hex()

	existingEmp := GetEmployee(strId)

	if updatedEmployee.Name != "" {
		existingEmp.Name = updatedEmployee.Name
	}

	if updatedEmployee.Age != 0 {
		existingEmp.Age = updatedEmployee.Age
	}

	if updatedEmployee.Salary != 0 {
		existingEmp.Salary = updatedEmployee.Salary
	}

	update := bson.M{
		"$set": bson.M{
			"name":   existingEmp.Name,
			"age":    existingEmp.Age,
			"salary": existingEmp.Salary,
		},
	}

	result, err := db.UpdateOne(ctx, filter, update)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No documents found")
		} else {
			log.Fatal("Error while updating user", err.Error())
		}
	}

	return result.ModifiedCount > 0
}
