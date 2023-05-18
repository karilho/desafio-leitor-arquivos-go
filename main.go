package main

import (
	"fmt"
)

func main() {

	var entradaUsuario int

	fmt.Println("Escolha a opção de 1 a 3")

	fmt.Scan(&entradaUsuario)

	switch entradaUsuario {
	case 1:
		cadastraCandidato()
	case 2:
		mapearFormulario()
		//fmt.Println("Chamou o 2")
	}

}
