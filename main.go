package main 

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
	"github.com/kcelestinomaria/tunzishaX-go/controllers"
	//"fmt"
	//"context"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	req := httprouter.New()

	userSession := controllers.NewPatientController(getSession())

	req.GET("/patient/:patientId", userSession.GetPatient)
	req.POST("/patient", userSession.CreatePatient)
	//req.UPDATE("")
	req.DELETE("/patient/:patientId", userSession.DeletePatient)

	http.ListenAndServe("localhost:9000", req)
}

func getSession() *mgo.Session{

	/*
	const connectingString = "mongodb+srv://celestino127:C0mpa$$i0n127@cluster0.5qsdpkx.mongodb.net/Simple-Node-API?retryWrites=true&w=majority"
	// Our database is MongoDB Atlas(On Cloud)

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectingString).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}


	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!") */

	 // for local Mongo DB connections
	databaseSession, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}
	return databaseSession
	
}