package main

import (
	"devmatheusguerra/routers"
	"net/http"
)

func main() {
	routers.CarregarRotas()
	http.ListenAndServe(":8000", nil)
}
