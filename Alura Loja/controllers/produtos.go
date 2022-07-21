package controllers

import (
	"devmatheusguerra/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscarTodosProdutos()
	templates.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func NovoProduto(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço")
			return
		}

		quantidade := r.FormValue("quantidade")
		quantidadeInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade")
			return
		}

		result := models.InserirProduto(nome, descricao, precoFloat, quantidadeInt)

		if result {
			http.Redirect(w, r, "/", 301)
		} else {
			http.Redirect(w, r, "/new", 400)
		}
	}

	http.Redirect(w, r, "/", 301)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, _ := strconv.Atoi(id)
	result := models.DeletarProduto(idInt)

	if result {
		http.Redirect(w, r, "/", 301)
	} else {
		http.Redirect(w, r, "/", 400)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Erro na conversão do id")
		return
	}
	nome := r.URL.Query().Get("nome")
	descricao := r.URL.Query().Get("descricao")

	preco := r.URL.Query().Get("preco")
	precoFloat, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		log.Println("Erro na conversão do preço")
		return
	}
	quantidade := r.URL.Query().Get("quantidade")
	quantidadeInt, err := strconv.Atoi(quantidade)
	if err != nil {
		log.Println("Erro na conversão da quantidade")
		return
	}

	produto := models.CriarProduto(idInt, nome, descricao, precoFloat, quantidadeInt)
	templates.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Erro na conversão do id")
		return
	}
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")

	preco := r.FormValue("preco")
	precoFloat, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		log.Println("Erro na conversão do preço")
		return
	}
	quantidade := r.FormValue("quantidade")
	quantidadeInt, err := strconv.Atoi(quantidade)
	if err != nil {
		log.Println("Erro na conversão da quantidade")
		return
	}

	result := models.EditarProduto(idInt, nome, descricao, precoFloat, quantidadeInt)

	if result {
		http.Redirect(w, r, "/", 301)
	} else {
		http.Redirect(w, r, "/", 400)
	}

}
