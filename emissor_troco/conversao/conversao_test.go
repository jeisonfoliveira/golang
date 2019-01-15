package conversao

import (
	"testing"
)

func TestConverterValorTroco(t *testing.T) {

	t.Log("Verifica troco passando diferentes números")

	t.Log("1 - Teste passando número ímpar")
	sliceNotas := make([][]int, 0, 0)
	a := []int{1, 20}
	sliceNotas = append(sliceNotas, a)
	a = []int{1, 10}
	sliceNotas = append(sliceNotas, a)
	a = []int{1, 5}
	sliceNotas = append(sliceNotas, a)
	valorTroco := 35
	executaConversaoTroco(valorTroco, sliceNotas, t)

	t.Log("2 - Teste passando número par")
	sliceNotas = make([][]int, 0, 0)
	a = []int{1, 50}
	sliceNotas = append(sliceNotas, a)
	a = []int{1, 20}
	sliceNotas = append(sliceNotas, a)
	a = []int{1, 10}
	sliceNotas = append(sliceNotas, a)
	valorTroco = 80
	executaConversaoTroco(valorTroco, sliceNotas, t)

	t.Log("3 - Teste passando número ímpar acima de 100")
	sliceNotas = make([][]int, 0, 0)
	a = []int{1, 100}
	sliceNotas = append(sliceNotas, a)
	a = []int{1, 20}
	sliceNotas = append(sliceNotas, a)
	a = []int{1, 5}
	sliceNotas = append(sliceNotas, a)
	valorTroco = 125
	executaConversaoTroco(valorTroco, sliceNotas, t)

	t.Log("4 - Teste passando número par acima de 100")
	sliceNotas = make([][]int, 0, 0)
	a = []int{1, 100}
	sliceNotas = append(sliceNotas, a)
	a = []int{1, 50}
	sliceNotas = append(sliceNotas, a)
	a = []int{2, 20}
	sliceNotas = append(sliceNotas, a)
	a = []int{1, 5}
	sliceNotas = append(sliceNotas, a)
	a = []int{1, 2}
	sliceNotas = append(sliceNotas, a)
	a = []int{1, 1}
	sliceNotas = append(sliceNotas, a)
	valorTroco = 198
	executaConversaoTroco(valorTroco, sliceNotas, t)


}

func executaConversaoTroco(valorTroco int, sliceNotas [][]int, t *testing.T ) {

	sliceRetorno := ConverterValorTroco(valorTroco)

	if len(sliceRetorno) != len(sliceNotas) {
		t.Errorf("\tTeste com falha: numero de notas diferente do esperado.")
	} else {
		for i := range sliceRetorno {
			//se o tipo de notas é diferente ou a quantidade é diferente
			if (sliceRetorno[i][1] == sliceNotas[i][1]) && (sliceRetorno[i][0] == sliceNotas[i][0]){
				continue
			} else {
				t.Log("\tTeste falhou. Deveria encontrar: ")
				t.Log( sliceNotas )
				t.Log("\tMas encontrou: ")
				t.Fatal( sliceRetorno )
			}
		}
	}
}