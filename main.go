package main

import(
	"net/http"

	"ecomerce/controller"
	"ecomerce/config"
)

func main() {

	config.ConnectDatabase()

	router := controller.SetupRouter()

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
