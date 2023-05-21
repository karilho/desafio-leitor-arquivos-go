package service

import "fmt"

func AbrirMenu() {
	var entradaUsuario int

	fmt.Println("Escolha a opção de 1 a 3")

	fmt.Scan(&entradaUsuario)

	switch entradaUsuario {
	case 1:
		CadastrarCandidato()
		return
	case 2:
		RemoverPergunta()
		return
	case 3:
		MapearFormulario()
		return
	}

}
