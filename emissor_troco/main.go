package main

import (
	"emissor_troco/conversao"
	"fmt"
	"strconv"
)

func main() {

	valorTotal := identificaValorConversao()

	sliceNotas := conversao.ConverterValorTroco(valorTotal)

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
