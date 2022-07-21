package transacoes

import (
	"fmt"
	"time"
)

type Transacao struct {
	data, tipo   string
	valor, saldo float64
}

func NovaTransacao(tipo string, valor float64, novoSaldo float64) Transacao {
	return Transacao{
		data:  time.Now().Format("02/01/2006 15:04:05"),
		tipo:  tipo,
		valor: valor,
		saldo: novoSaldo,
	}
}

func (t *Transacao) ObterData() string {
	return t.data
}

func (t *Transacao) ObterTipo() string {
	return t.tipo
}

func (t *Transacao) ObterValor() string {
	return fmt.Sprintf("R$ %.2f", t.valor)
}

func (t *Transacao) ObterSaldo() string {
	return fmt.Sprintf("R$ %.2f", t.saldo)
}
