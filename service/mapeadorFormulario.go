package service

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Candidato struct {
	Nome            string
	Email           string
	Idade           int
	Cidade          string
	RespostasExtras []string
}

func MapearCandidatos() ([]Candidato, error) {
	var candidatos []Candidato

	// Ler o diretório "candidatos"
	files, err := ioutil.ReadDir("candidatos")
	if err != nil {
		fmt.Printf("Erro ao ler o diretório 'candidatos': %v\n", err)
		return nil, nil
	}

	// Percorrer cada arquivo .txt
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			candidato, err := mapearCandidatoArquivo(file.Name())
			if err != nil {
				fmt.Printf("Erro ao mapear o candidato do arquivo '%s': %v\n", file.Name(), err)
				continue
			}
			candidatos = append(candidatos, candidato)
		}
	}

	return candidatos, err
}

func mapearCandidatoArquivo(filename string) (Candidato, error) {
	candidato := Candidato{}

	// Abrir o arquivo
	file, err := os.Open("candidatos/" + filename)
	if err != nil {
		return candidato, fmt.Errorf("Erro ao abrir o arquivo '%s': %v", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0

	// Ler as linhas do arquivo
	for scanner.Scan() {
		line := scanner.Text()
		// Atribuir valores às posições corretas da struct
		switch counter {
		case 0:
			candidato.Nome = line
		case 1:
			candidato.Email = line
		case 2:
			idade, err := strconv.Atoi(line)
			if err != nil {
				return candidato, fmt.Errorf("Erro ao converter idade para inteiro: %v", err)
			}
			candidato.Idade = idade
		case 3:
			candidato.Cidade = line
		default:
			candidato.RespostasExtras = append(candidato.RespostasExtras, line)
		}
		counter++
	}

	if err := scanner.Err(); err != nil {
		return candidato, fmt.Errorf("Erro ao ler as linhas do arquivo '%s': %v", filename, err)
	}

	return candidato, nil
}

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

	fmt.Println(slicePerguntas)
	return slicePerguntas
}
