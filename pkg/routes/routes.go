package routes

import (
	"bharatrail_server/pkg/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/addCity", controller.AddCity).Methods("POST")
	router.HandleFunc("/addTrain", controller.AddTrain).Methods("POST")
	router.HandleFunc("/addUser", controller.AddUser).Methods("POST")
	router.HandleFunc("/getTrain/{trainId}", controller.GetTrain).Methods("GET")

	return router
}
