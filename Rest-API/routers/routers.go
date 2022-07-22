package routers

import (
	"devmatheusguerra/rest/controllers"
	"devmatheusguerra/rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentType)
	r.Use(middleware.JWT)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("GET")
	r.HandleFunc("/personalidades", controllers.CriarNovaPersonalidade).Methods("POST")
	r.HandleFunc("/personalidades/{id}", controllers.DeletarPersonalidade).Methods("DELETE")
	r.HandleFunc("/personalidades/{id}", controllers.EditarPersonalidade).Methods("PUT")

	r.HandleFunc("/auth", controllers.Autenticar).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))
}
