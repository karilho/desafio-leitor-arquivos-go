package service

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func MapearFormulario() []string {

	//Ler o arquivo
	conteudoDoFormulario, err := os.ReadFile("formulario.txt")
	if err != nil {
		log.Fatalf("Erro ao ler arquivo: %v", err)
	}

	//pegar o conteudo do arquivo e transformar em string, sem espaços
	stringsPerguntas := strings.Split(string(conteudoDoFormulario), "\n")
	var slicePerguntas []string

	//Percorre e adiciona no slice
	for i, pergunta := range stringsPerguntas {
		pergunta = strings.TrimSpace(pergunta)
		pergunta = strings.TrimPrefix(pergunta, fmt.Sprintf("%d - ", i+1))
		if pergunta != "" {
			slicePerguntas = append(slicePerguntas, pergunta)
		}
	}

	//Percorre e imprime na tela
	//for _, pergunta := range slicePerguntas {
	//	fmt.Println(pergunta)
	//}
	//fmt.Println(slicePerguntas)
	return slicePerguntas
}
