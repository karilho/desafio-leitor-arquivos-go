package service

import (
	"bufio"
	"desafio-leitor-arquivos-go/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func CadastrarCandidato() {

	sliceRespostas := []string{}
	perguntas := MapearFormulario()
	candidatosJaCadastrados, _ := MapearCandidatos()
	var contadorDeCandidatosCadastrados string = strconv.Itoa(len(candidatosJaCadastrados) + 1)

	scannerEscritorRespostas := bufio.NewScanner(os.Stdin)
	//para o fscanner funcionar, é preciso dar um enter a mais
	scannerEscritorRespostas.Scan()

	for i, pergunta := range perguntas {
		fmt.Println(pergunta)
		scannerEscritorRespostas.Scan()
		resposta := scannerEscritorRespostas.Text()
		if i == 2 {
			for {
				idadeStr := resposta
				idade, err := strconv.Atoi(idadeStr)
				if err != nil || idade < 16 {
					fmt.Println("Idade inválida. Por favor, insira uma idade válida maior ou igual a 16 anos.")
					fmt.Println(pergunta)
					scannerEscritorRespostas.Scan()
					resposta = scannerEscritorRespostas.Text()
				} else {
					break
				}
			}
		}
		sliceRespostas = append(sliceRespostas, resposta)
	}

	nomeArquivoCerto := contadorDeCandidatosCadastrados + "-" + utils.NameFormater(sliceRespostas[0])
	nomeArquivo := filepath.Join("candidatos", nomeArquivoCerto)
	arquivo, err := os.Create(nomeArquivo)
	if err != nil {
		fmt.Printf("Erro ao criar o arquivo: %v", err)
		return
	}

	for _, resposta := range sliceRespostas {
		arquivo.WriteString(resposta + "\n")
	}

	defer arquivo.Close()

	// Specify the path of the file you want to upload
	filePath := nomeArquivo

	// Upload the file to S3
	err = utils.UploadFile(utils.AWS_S3_BUCKET, filePath, utils.AWS_S3_REGION)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File uploaded successfully!")

	fmt.Println("Candidato cadastrado com sucesso!")

	AbrirMenu()
}

func ListarPorIdade() {
	candidatos, _ := MapearCandidatos()

	nomeEIdade := make(map[int][]string)
	contadorIdade := make(map[int]int) // Mapa para armazenar a contagem de candidatos por idade

	// Agrupa os nomes dos candidatos por idade
	for _, candidato := range candidatos {
		nomeEIdade[candidato.Idade] = append(nomeEIdade[candidato.Idade], candidato.Nome)
		contadorIdade[candidato.Idade]++
	}

	// Exibe os nomes dos candidatos agrupados por idade
	for idade, nomes := range nomeEIdade {
		fmt.Printf("Idade: %d\n", idade)
		for _, nome := range nomes {
			fmt.Println("Nome(s):" + nome)
		}

		fmt.Printf("Candidatos com %d anos = %d candidato(s)\n", idade, contadorIdade[idade])
		fmt.Println()
	}
}

func EncontrarCndDuplicados() {
	candidatos, _ := MapearCandidatos()

	formulariosDuplicados := make(map[string][]Candidato)

	// Verifica duplicidades com base no nome e e-mail
	for _, candidato := range candidatos {
		chave := candidato.Nome + "|" + candidato.Email
		formulariosDuplicados[chave] = append(formulariosDuplicados[chave], candidato)
	}

	// Exibe os formulários duplicados
	for chave, duplicados := range formulariosDuplicados {
		if len(duplicados) > 1 {
			fmt.Println("Formulário duplicado:")
			fmt.Println("Chave: " + chave)
			for _, candidato := range duplicados {
				fmt.Println("Nome: " + candidato.Nome)
				fmt.Println("Email: " + candidato.Email)
				fmt.Println()
			}
		}
	}
}

func BuscarCandidato() {
	candidatos, _ := MapearCandidatos()

	var nomePesquisado string
	fmt.Println("Digite o nome do candidato:")
	_, err := fmt.Scan(&nomePesquisado)
	if err != nil {
		fmt.Printf("Erro ao ler o nome do candidato: %v", err)
		return
	}

	// Verifica se o candidato está cadastrado
	encontrado := false
	for _, candidato := range candidatos {
		if candidato.Nome == nomePesquisado {
			//Todo: verificar se o email é igual
			fmt.Println("Candidato encontrado:")
			fmt.Println(candidato)
			encontrado = true
			break
		}
	}

	if !encontrado {
		fmt.Println("Candidato não encontrado.")
	}
}
