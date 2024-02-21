package controllers

import (
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type PatientController struct{
	databaseSession *mgo.Session
}

func NewPatientController(databaseSession *mgo.Session) *PatientController{

	return &PatientController{databaseSession}
}

//GetPatient

//CreatePatient

//DeletePatient