package main

import (
	"net/http"

	"ecomerce/config"
	"ecomerce/controller"
)

func main() {

	config.ConnectDatabase()

	router := controller.SetupRouter()

	http.Handle("/", router)
	http.ListenAndServe(":8090", nil)
}
