package controller

import(
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// userManagemnt
	router.HandleFunc("/register", createUser).Methods("POST")
	router.HandleFunc("/editUser", editUser).Methods("PUT")
	router.HandleFunc("/userdetail", GetUserById).Methods("GET")
	return router
}