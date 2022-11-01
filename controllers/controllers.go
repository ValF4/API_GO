package controllers

import (
	"API_GO/database"
	"API_GO/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func TodasBairros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p []models.AllBairros
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func ReturnID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	var b models.AllBairros
	database.DB.First(&b, id)
	json.NewEncoder(w).Encode(b)
}

func Creatbairro(w http.ResponseWriter, r *http.Request) {
	var newbairro models.AllBairros
	json.NewDecoder(r.Body).Decode(&newbairro)
	database.DB.Create(&newbairro)
	json.NewEncoder(w).Encode(newbairro)

}

func Deletbairro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	var b models.AllBairros
	database.DB.Delete(&b, id)
	json.NewEncoder(w).Encode(b)
}

func Editbairro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	var b models.AllBairros
	database.DB.First(&b, id)
	json.NewDecoder(r.Body).Decode(&b)
	database.DB.Save(&b)
	json.NewEncoder(w).Encode(b)
}
