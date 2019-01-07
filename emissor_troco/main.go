package main

import (
	"emissor_troco/converter"
	"fmt"
	"strconv"
)

type notificacao interface {
	notificar()
}

type funcionario struct {
	nome string
}

func (c funcionario) notificar() {
	fmt.Println("Movimentação realizada pelo funcionario: " + c.nome)
}

type gerente struct {
	nome string
}

func (g gerente) notificar() {
	fmt.Println("Movimentação realizada pelo gerente: " + g.nome)
}

func main() {

	valorTotal := identificaValorConversao()

	sliceNotas := converter.ConverterValorTroco(valorTotal)

	fmt.Println("O troco é: ")
	for _, value := range sliceNotas {
		fmt.Println(strconv.Itoa(value[0]) + " nota(s) de " + strconv.Itoa(value[1]) + " reais")
	}

	caixa1 := funcionario{"Joao"}
	emitirNotificacao(caixa1)

	gerente1 := gerente{"Jose"}
	emitirNotificacao(gerente1)

}

func identificaValorConversao() int {

	var valorTotal int
	fmt.Println("Digite o valor para conversão:")
	fmt.Scanf("%d", &valorTotal)

	return valorTotal
}

func emitirNotificacao(n notificacao) {

	n.notificar()

}
