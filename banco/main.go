package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
	"time"
)

func pagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) (string, float64)
}

func main() {
	conta1 := contas.ContaPoupanca{
		Titular: clientes.Cliente{
			Nome:      "João",
			CPF:       "123.456.789-00",
			Profissao: "Engenheiro",
		},
		Operacao: 13,
		Agencia:  123,
		Conta:    456,
	}

	conta1.Depositar(1000)
	time.Sleep(1 * time.Second)
	pagarBoleto(&conta1, 100.0)
	time.Sleep(1 * time.Second)
	conta1.Depositar(300)
	time.Sleep(2 * time.Second)
	pagarBoleto(&conta1, 200.0)
	conta1.Depositar(50)
	time.Sleep(1 * time.Second)
	conta1.Sacar(1000)
	fmt.Println(conta1.ObterExtratoFormatado())

	contaCorrente := contas.ContaCorrente{
		Titular: clientes.Cliente{
			Nome:      "João",
			CPF:       "123.456.789-00",
			Profissao: "Engenheiro",
		},
		Agencia: 123,
		Conta:   456,
	}

	contaCorrente.Depositar(1000)
	time.Sleep(1 * time.Second)
	pagarBoleto(&contaCorrente, 100.0)
	time.Sleep(1 * time.Second)
	contaCorrente.Depositar(300)
	time.Sleep(2 * time.Second)
	pagarBoleto(&contaCorrente, 200.0)
	contaCorrente.Depositar(50)
	time.Sleep(1 * time.Second)
	contaCorrente.Sacar(1000)
	fmt.Println(contaCorrente.ObterExtratoFormatado())
}
