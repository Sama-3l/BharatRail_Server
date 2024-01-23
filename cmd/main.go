package main

import (
	"bharatrail_server/pkg/routes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func createPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	addr := ":" + port

	return addr
}

func main() {
	r := routes.Router()
	port := createPort()
	fmt.Printf("Serving api at port%v\n", port)

	log.Fatal(http.ListenAndServe(port, r))
}
