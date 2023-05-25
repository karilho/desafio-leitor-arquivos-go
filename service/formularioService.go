package service

import (
	"bufio"
	"fmt"
	"os"
)

func AdicionarPergunta() {
	perguntasVelhas := MapearFormulario()

	fmt.Println("Digite a pergunta que deseja adicionar")
	leitorPerguntaLinha := bufio.NewScanner(os.Stdin)
	leitorPerguntaLinha.Scan()
	leitorPerguntaLinha.Scan()
	novaPergunta := leitorPerguntaLinha.Text()
	perguntasVelhas = append(perguntasVelhas, novaPergunta)
	fmt.Println("Captured question:", novaPergunta)

	novoFormulario, err := os.Create("formulario.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo")
		return
	}

	for i, pergunta := range perguntasVelhas {
		novoFormulario.WriteString(fmt.Sprintf("%d - ", i+1) + pergunta + "\n")
	}

	fmt.Println(&perguntasVelhas)

	defer novoFormulario.Close()

	fmt.Println("Formulário atualizado com sucesso!")

	fmt.Println(&novoFormulario)

	AbrirMenu()
}

func RemoverPergunta() {
	perguntas := MapearFormulario()

	fmt.Println("Escolha a pergunta que deseja remover")
	for i, pergunta := range perguntas {
		fmt.Println(i+1, " - "+pergunta)
	}

	var numeroPerguntaRemover int
	fmt.Scan(&numeroPerguntaRemover)
	numeroPerguntaRemover--

	if numeroPerguntaRemover < 4 || numeroPerguntaRemover > len(perguntas) {
		fmt.Println("Não é possível remover as perguntas de 1 a 4, por favor, digite um número válido.")
		RemoverPergunta()
	}
	perguntas = append(perguntas[:numeroPerguntaRemover], perguntas[numeroPerguntaRemover+1:]...)

	fmt.Println("Pergunta removida com sucesso!")

	fmt.Println(perguntas)

	novoFormulario, err := os.Create("formulario.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo")
		return
	}

	for i, pergunta := range perguntas {
		novoFormulario.WriteString(fmt.Sprintf("%d - ", i+1) + pergunta + "\n")
	}

	defer novoFormulario.Close()

	fmt.Println("Formulário atualizado com sucesso!")

	fmt.Println(&novoFormulario)

	AbrirMenu()
}
