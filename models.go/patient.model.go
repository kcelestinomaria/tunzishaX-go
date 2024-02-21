package models

import "gopkg.in/mgo.v2/bson"

type Patient struct {
	patientId     bson.ObjectId `json:"patientId" bson: "_patientId"`
	firstName     string        `json:"firstName" bson: "firstName"`
	secondName    string        `json: "secondName" bson: "secondName"`
	age           int           `json: "age" bson: "age"`
	gender        string        `json: "gender" bson: "age"`
	contactNumber string        `json: "contactNumber" bson: "contactNumber"`
	nextOfKin     string        `json: "nextOfKin" bson: "nextOfKin"`
	dateOfBirth   string        `json: dateOfBirth" bson: "dateOfBirth"`
}
