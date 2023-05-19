package service

import (
	"bufio"
	"fmt"
	"os"
)

func CadastrarCandidato() {

	sliceRespostas := []string{}
	perguntas := MapearFormulario()
	scannerEscritorRespostas := bufio.NewScanner(os.Stdin)

	//para o fscanner funcionar, é preciso dar um enter a mais
	scannerEscritorRespostas.Scan()

	for _, pergunta := range perguntas {
		fmt.Println(pergunta)
		scannerEscritorRespostas.Scan()
		resposta := scannerEscritorRespostas.Text()
		sliceRespostas = append(sliceRespostas, resposta)
	}
	fmt.Println(perguntas)
	fmt.Println(sliceRespostas)

	nomeArquivo := sliceRespostas[0] + ".txt"

	arquivo, err := os.Create(nomeArquivo)
	if err != nil {
		fmt.Printf("Erro ao criar o arquivo: %v", err)
		return
	}

	for _, resposta := range sliceRespostas {
		arquivo.WriteString(resposta + "\n")
	}

	defer arquivo.Close()
}
