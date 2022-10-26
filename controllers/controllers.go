package controllers

import (
	"API_GO/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

func ReturnID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	for _, infosSetores := range models.Personalidades {
		if strconv.Itoa(infosSetores.ID) == id {
			json.NewEncoder(w).Encode(infosSetores)
		}
	}
}
