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
	r.HandleFunc("/api/bairros", controllers.TodasBairros).Methods("Get")
	r.HandleFunc("/api/bairros/{ID}", controllers.ReturnID).Methods("Get")
	r.HandleFunc("/api/bairros/", controllers.Creatbairro).Methods("Post")
	r.HandleFunc("/api/bairros/{ID}", controllers.Deletbairro).Methods("Delete")
	r.HandleFunc("/api/bairros/{ID}", controllers.Editbairro).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
