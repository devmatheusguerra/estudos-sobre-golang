package main

import (
	"devmatheusguerra/rest/database"
	"devmatheusguerra/rest/routers"
	"fmt"
)

func main() {
	fmt.Println("Iniciando o servidor rest com GO")
	database.ConectaComBancoDeDados()
	routers.HandleRequest()
}
