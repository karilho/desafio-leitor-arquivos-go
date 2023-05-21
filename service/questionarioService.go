package service

import "fmt"

func AdicionarPergunta() {

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
	perguntas = append(perguntas[:numeroPerguntaRemover], perguntas[numeroPerguntaRemover+1:]...)

	fmt.Println("Pergunta removida com sucesso!")
	fmt.Println(perguntas)
}
