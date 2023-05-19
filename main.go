package main

import (
	"desafio-leitor-arquivos-go/service"
	"fmt"
)

func main() {

	var entradaUsuario int

	fmt.Println("Escolha a opção de 1 a 3")

	fmt.Scan(&entradaUsuario)

	switch entradaUsuario {
	case 1:
		service.CadastrarCandidato()
	case 2:
		service.MapearFormulario()
		fmt.Println("Chamou exibir formulários")
	}

}
