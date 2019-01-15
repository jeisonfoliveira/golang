package main

import "fmt"

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

func emitirNotificacao(n notificacao) {

	n.notificar()

}
