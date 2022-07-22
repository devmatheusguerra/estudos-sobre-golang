package controllers

import (
	"devmatheusguerra/rest/database"
	"devmatheusguerra/rest/models"
	"devmatheusguerra/rest/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)
	json.NewEncoder(w).Encode(personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func CriarNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Create(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}

func DeletarPersonalidade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode("SUCCESS")

}

func EditarPersonalidade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	if personalidade.Id == 0 {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	json.NewDecoder(r.Body).Decode(&personalidade)
	personalidade_id, err := strconv.Atoi(id) // Evita que o ID seja alterado
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	personalidade.Id = personalidade_id
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}

func Autenticar(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	if user.Username == "matheus" && user.Password == "123" {
		token := utils.GerarJWT(user.Username)
		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
	}

}
