package contas

import (
	"banco/clientes"
	"banco/extratos"
	"banco/transacoes"
)

type ContaPoupanca struct {
	Titular                  clientes.Cliente
	Agencia, Conta, Operacao int
	saldo                    float64
	extrato                  extratos.Extrato
}

func (c *ContaPoupanca) Sacar(valorSaque float64) (string, float64) {
	podeSacar := valorSaque <= c.saldo && valorSaque > 0
	if podeSacar {
		c.saldo -= valorSaque

		transacao := transacoes.NovaTransacao("SAQ", valorSaque, c.saldo)
		c.extrato.AdicionarTransacao(transacao)

		return "Saque realizado com sucesso", c.saldo
	}

	return "saldo insuficiente", c.saldo
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito

		transacao := transacoes.NovaTransacao("DEP", valorDeposito, c.saldo)
		c.extrato.AdicionarTransacao(transacao)

		return "Depósito realizado com sucesso", c.saldo
	}
	return "Valor de depósito inválido", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}

func (c *ContaPoupanca) ObterExtratoFormatado() string {
	return c.extrato.ObterExtratoComoTexto()
}

func (c *ContaPoupanca) ObterExtratoBruto() extratos.Extrato {
	return c.extrato
}
