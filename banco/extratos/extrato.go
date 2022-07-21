package extratos

import "banco/transacoes"

type Extrato struct {
	Transacoes []transacoes.Transacao
}

func (e *Extrato) AdicionarTransacao(t transacoes.Transacao) {
	e.Transacoes = append(e.Transacoes, t)
}

func (e *Extrato) ObterExtratoComoTexto() string {
	texto := "\tData\t\t\tTipo\t\tValor\t\tSaldo\n"
	for _, t := range e.Transacoes {
		texto += "\t" + t.ObterData() + "\t" + t.ObterTipo() + "\t\t" + t.ObterValor() + "\t" + t.ObterSaldo() + "\n"
		texto += "=====================================================================================\n"
	}
	return texto
}
