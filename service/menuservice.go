package service

import (
	"fmt"
)

func AbrirMenu() {
	imprimirMenuTela()
	var entradaUsuario int
	fmt.Println("Escolha uma opção")

	_, err := fmt.Scan(&entradaUsuario)
	if err != nil {
		fmt.Println("Erro ao ler a entrada do usuário")
		return
	}

	switch entradaUsuario {
	case 1:
		CadastrarCandidato()
		return
	case 2:
		AdicionarPergunta()
	case 3:
		RemoverPergunta()
	case 4:
		ListarCandidatosPorIdade()
		AbrirMenu()
	case 5:
		//PesquisarFormularios()
	case 6:
		fmt.Println("Saindo do sistema...")
		return
	}
}

func imprimirMenuTela() {

	fmt.Println("######### Bem vindo ao sistema de cadastro de candidatos! ###########" + "\n" + "\n" +
		"######### O que deseja fazer? #########" + "\n")

	arrayPerguntas := [6]string{
		"1 - Cadastrar candidato.",
		"2 - Adicionar perguntado formulário.",
		"3 - Remover pergunta do formulário.",
		"4 - Listar formulários cadastrados.",
		"5 - Pesquisar formulários cadastrados.",
		"6 - Sair do site."}

	for _, opcaoMenu := range arrayPerguntas {
		fmt.Println(opcaoMenu)
	}
}
