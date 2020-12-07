package main

import (
	"context"
	"dockergo/helper"
	"dockergo/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var collection = helper.ConnectDB()

func main() {

	r := mux.NewRouter()
	fmt.Println(":9036!")
	r.HandleFunc("/api/customers", getCustomers).Methods("GET")
	r.HandleFunc("/api/customers", createCustomer).Methods("POST")

	log.Fatal(http.ListenAndServe(":9036", r))

}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var Customer models.Customer

	_ = json.NewDecoder(r.Body).Decode(&Customer)

	result, err := collection.InsertOne(context.TODO(), Customer)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var Customers []models.Customer

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var Customer models.Customer

		err := cur.Decode(&Customer)
		if err != nil {
			log.Fatal(err)
		}

		Customers = append(Customers, Customer)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(Customers)
}
