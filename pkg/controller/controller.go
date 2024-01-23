package controller

import (
	"bharatrail_server/pkg/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AddCity(w http.ResponseWriter, req *http.Request) {
	var city models.City

	err := json.NewDecoder(req.Body).Decode(&city)

	if err != nil {
		log.Printf("Unable to decode %v \n", err)
	}

	city, err = models.AddCityToDatabase(city)

	if err != nil {
		log.Printf("Error: %v \n", err)
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(city)

	fmt.Printf("%v\n", city)
}

func AddTrain(w http.ResponseWriter, req *http.Request) {
	var train models.Train

	err := json.NewDecoder(req.Body).Decode(&train)

	if err != nil {
		log.Printf("Unable to decode %v \n", err)
	}

	train, err = models.AddTrainToDatabase(train)

	if err != nil {
		log.Printf("Error: %v \n", err)
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(train)
}

func AddUser(w http.ResponseWriter, req *http.Request) {
}
