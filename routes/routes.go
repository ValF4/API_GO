package routes

import (
	"API_GO/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{ID}", controllers.ReturnID).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
