package routes

import (
	"../controllers"
	"github.com/gorilla/mux"
)

//GetRouter handling routes
func GetRouter() (router *mux.Router) {
	r := mux.NewRouter().StrictSlash(true)
	router = r.PathPrefix("/api").Subrouter()

	router.HandleFunc("/planets", controllers.List).Methods("GET")
	router.HandleFunc("/planets/{id}", controllers.GetByID).Methods("GET")
	router.HandleFunc("/planets", controllers.Create).Methods("POST")
	router.HandleFunc("/planets", controllers.Update).Methods("PUT")
	router.HandleFunc("/planets/{id}", controllers.Delete).Methods("DELETE")

	return
}
