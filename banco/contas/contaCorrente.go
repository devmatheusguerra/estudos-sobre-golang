package contas

import (
	"banco/clientes"
	"banco/extratos"
	"banco/transacoes"
)

type ContaCorrente struct {
	Titular        clientes.Cliente
	Agencia, Conta int
	saldo          float64
	extrato        extratos.Extrato
}

func (c *ContaCorrente) Sacar(valorSaque float64) (string, float64) {
	podeSacar := valorSaque <= c.saldo && valorSaque > 0
	if podeSacar {
		c.saldo -= valorSaque

		transacao := transacoes.NovaTransacao("SAQ", valorSaque, c.saldo)
		c.extrato.AdicionarTransacao(transacao)

		return "Saque realizado com sucesso", c.saldo
	}

	return "saldo insuficiente", c.saldo
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito

		transacao := transacoes.NovaTransacao("DEP", valorDeposito, c.saldo)
		c.extrato.AdicionarTransacao(transacao)

		return "Depósito realizado com sucesso", c.saldo
	}
	return "Valor de depósito inválido", c.saldo
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

func (c *ContaCorrente) ObterExtratoFormatado() string {
	return c.extrato.ObterExtratoComoTexto()
}

func (c *ContaCorrente) ObterExtratoBruto() extratos.Extrato {
	return c.extrato
}
