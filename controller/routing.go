package controller

import (
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// userManagemnt
	router.HandleFunc("/register", createUser).Methods("POST")
	router.HandleFunc("/editUser/{id}", editUser).Methods("PUT")
	router.HandleFunc("/userdetail/{id}", getUserById).Methods("GET")
	router.HandleFunc("/deleteUser/{id}", deleteUserbyId).Methods("DELETE")
	router.HandleFunc("/getListUser", getListUser).Methods("GET")
	return router
}
