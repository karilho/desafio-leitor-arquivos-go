package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func mapearFormulario() []string {

	//Ler o arquivo
	conteudoDoFormulario, err := os.ReadFile("formulario.txt")
	if err != nil {
		log.Fatalf("Erro ao ler arquivo: %v", err)
	}

	//pegar o conteudo do arquivo e transformar em string, sem espaços
	stringsPerguntas := strings.Split(string(conteudoDoFormulario), "\n")
	var slicePerguntas []string

	//Percorre e adiciona no slice
	for _, pergunta := range stringsPerguntas {
		pergunta = strings.TrimSpace(pergunta)
		if pergunta != "" {
			slicePerguntas = append(slicePerguntas, pergunta)
		}
	}

	//Percorre e imprime na tela
	for _, pergunta := range slicePerguntas {
		fmt.Println(pergunta)

	}

	return slicePerguntas;
}