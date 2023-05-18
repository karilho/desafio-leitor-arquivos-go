package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cadastraCandidato() {

	perguntas := mapearFormulario()

	fmt.Println("Digite o nome do candidato:")
	nomeCandidato, _ := lerRespostaDoUsuario()

	// Limpar o nome do candidato removendo espaços em branco extras
	nomeCandidato = strings.TrimSpace(nomeCandidato)

	// Criar o nome do arquivo com base no nome do candidato
	nomeArquivo := nomeCandidato + ".txt"

	// Abrir o arquivo para escrita
	arquivo, err := os.Create(nomeArquivo)
	if err != nil {
		fmt.Printf("Erro ao criar o arquivo: %v", err)
		return
	}
	defer arquivo.Close()

	// Criar um escritor para o arquivo
	escritor := bufio.NewWriter(arquivo)

	// Pedir e gravar as respostas no arquivo
	for _, pergunta := range perguntas {
		fmt.Println(pergunta)
		resposta, _ := lerRespostaDoUsuario()
		escritor.WriteString(resposta + "\n")
	}

	// Certificar que todas as informações foram gravadas no arquivo
	escritor.Flush()

	fmt.Printf("Respostas do candidato '%s' gravadas com sucesso no arquivo '%s'!\n", nomeCandidato, nomeArquivo)
}

// Função para ler a resposta do usuário
func lerRespostaDoUsuario() (string, error) {
	leitorEntradaDoUsuario := bufio.NewReader(os.Stdin)
	resposta, err := leitorEntradaDoUsuario.ReadString('\n')
	if err != nil {
		return "", err
	}
	return resposta, nil
}
