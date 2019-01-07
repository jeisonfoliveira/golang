package converter

func ConverterValorTroco(valorTotal int) [][]int{
	var valorSoma int
	var proximaNota bool
	var i int

	arrayNotas := [7]int{100, 50, 20, 10, 5, 2, 1}

	mapNotas := map[int]int{
		100: 0,
		50:  0,
		20:  0,
		10:  0,
		5:   0,
		2:   0,
		1:   0,
	}
	//inicializa o valor da nota com 100 reais
	valorNota := arrayNotas[i]
	//enquanto não atingir o valor estipulado total permanece no loop
	for valorSoma < valorTotal {
		if proximaNota {
			i++
			valorNota = arrayNotas[i]
			proximaNota = false
		}
		//enquanto o valor da soma permanecer menor ou igual o valor da nota atual continua somando
		if valorSoma+valorNota <= valorTotal {
			if _, ok := mapNotas[valorNota]; ok {
				mapNotas[valorNota] += 1
				valorSoma += valorNota
			}
		} else {
			proximaNota = true
		}
	}
	i = 0
	valorNota = arrayNotas[i]
	sliceNotas := make([][]int, 0, 0)
	for range mapNotas {
		valorNota = arrayNotas[i]
		if q, _ := mapNotas[valorNota]; q > 0 {
			a := []int{q, valorNota}
			sliceNotas = append(sliceNotas, a)
		}
		i++
	}

	return sliceNotas
}
