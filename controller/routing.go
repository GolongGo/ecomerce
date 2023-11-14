package controller

import(
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// userManagemnt
	router.HandleFunc("/register", CreateUser).Methods("POST")
	return router
}