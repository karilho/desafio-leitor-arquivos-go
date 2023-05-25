package service

import (
	"desafio-leitor-arquivos-go/utils"
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
	case 2:
		AdicionarPergunta()
	case 3:
		RemoverPergunta()
	case 4:
		ListarPorIdade()
		AbrirMenu()
	case 5:
		//Feito
		EncontrarCndDuplicados()
	case 6:
		//Feito
		BuscarCandidato()
	case 7:
		fmt.Println("Saindo do sistema...")
		return
	case 8:
		utils.NameFormater("Teste aRNALDíNO")
	}
}

func imprimirMenuTela() {

	fmt.Println("######### Bem vindo ao sistema de cadastro de candidatos! ###########" + "\n" + "\n" +
		"######### O que deseja fazer? #########" + "\n")

	arrayPerguntas := [7]string{
		"1 - Cadastrar candidato.",
		"2 - Adicionar perguntado formulário.",
		"3 - Remover pergunta do formulário.",
		"4 - Listar formulários cadastrados.",
		"5 - Validar formulários cadastrados.",
		"6 - Pesquisar formulários cadastrados.",
		"7 - Sair do site."}

	for _, opcaoMenu := range arrayPerguntas {
		fmt.Println(opcaoMenu)
	}
}
