package routers

import (
	"devmatheusguerra/controllers"
	"net/http"
)

func CarregarRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.NovoProduto)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
