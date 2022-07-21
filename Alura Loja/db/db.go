package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	// NOTA: O usuário e senha do banco devem ser definidos no arquivo .env
	conexao := "user=postgres dbname=alura_loja password=123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err)
	}
	return db
}
