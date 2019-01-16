package conversao

import (
	"testing"
)

func TestConverterValorTroco(t *testing.T) {

	t.Log("Verifica troco passando diferentes números")

	t.Log("1 - Teste passando número ímpar")
	sliceNotas := [][]int{{1, 20}, {1, 10}, {1, 5}}
	valorTroco := 35
	executaConversaoTroco(valorTroco, sliceNotas, t)

	t.Log("2 - Teste passando número par")
	sliceNotas = [][]int{{1, 50}, {1, 20}, {1, 10}}
	valorTroco = 80
	executaConversaoTroco(valorTroco, sliceNotas, t)

	t.Log("3 - Teste passando número ímpar acima de 100")
	sliceNotas = [][]int{{1, 100}, {1, 20}, {1, 5}}
	valorTroco = 125
	executaConversaoTroco(valorTroco, sliceNotas, t)

	t.Log("4 - Teste passando número par acima de 100")
	sliceNotas = [][]int{{1, 100}, {1, 50}, {2, 20}, {1, 5}, {1, 2}, {1, 1}}
	valorTroco = 198
	executaConversaoTroco(valorTroco, sliceNotas, t)

}

func executaConversaoTroco(valorTroco int, sliceNotas [][]int, t *testing.T) {

	sliceRetorno := ConverterValorTroco(valorTroco)

	if len(sliceRetorno) != len(sliceNotas) {
		t.Errorf("\tTeste com falha: numero de notas diferente do esperado.")
	} else {
		for i := range sliceRetorno {
			//se o tipo de notas é diferente ou a quantidade é diferente
			if (sliceRetorno[i][1] == sliceNotas[i][1]) && (sliceRetorno[i][0] == sliceNotas[i][0]) {
				continue
			} else {
				t.Log("\tTeste falhou. Deveria encontrar: ")
				t.Log(sliceNotas)
				t.Log("\tMas encontrou: ")
				t.Fatal(sliceRetorno)
			}
		}
	}
}
