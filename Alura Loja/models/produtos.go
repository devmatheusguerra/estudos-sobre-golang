package models

import "devmatheusguerra/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func CriarProduto(id int, nome string, descricao string, preco float64, quantidade int) Produto {
	p := Produto{}
	p.Id = id
	p.Nome = nome
	p.Descricao = descricao
	p.Preco = preco
	p.Quantidade = quantidade
	return p
}

func BuscarTodosProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	selectProdutos, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}
	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		p.Id = id

		produtos = append(produtos, p)

	}
	defer db.Close()
	return produtos
}

func InserirProduto(nome string, descricao string, preco float64, quantidade int) bool {
	db := db.ConectaComBancoDeDados()
	_, err := db.Query("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)", nome, descricao, preco, quantidade)
	if err != nil {
		panic(err.Error())
		return false
	}
	defer db.Close()
	return true
}

func DeletarProduto(id int) bool {
	db := db.ConectaComBancoDeDados()
	_, err := db.Query("DELETE FROM produtos WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
		return false
	}
	defer db.Close()
	return true
}

func EditarProduto(id int, nome string, descricao string, preco float64, quantidade int) bool {
	db := db.ConectaComBancoDeDados()
	_, err := db.Query("UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5", nome, descricao, preco, quantidade, id)
	if err != nil {
		panic(err.Error())
		return false
	}
	defer db.Close()
	return true

}
