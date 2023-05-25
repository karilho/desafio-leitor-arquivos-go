<<<<<<< HEAD
# desafio-leitor-arquivos-go
Criação de desafio de ler e escrever arquivos .txt com GO
=======
## Desafio Leitor de Arquivos Go ##

A ideia deste desafio é replicar um desafio que eu realizei durante o programa MOVIMENTO CODAR da [Bluesoft](https://bluesoft.com.br/).

Originalmente, o desafio foi feito em Java, mas eu decidi replicá-lo em Go, para fins de aprendizado.


### O Desafio ###

O desafio consiste em criar um programa que leia um arquivo .txt com uma lista de perguntas feitas para cadastrar um candidato.

O programa deve ler o arquivo de perguntas e, para cada pergunta, deve ler a resposta do usuário e armazená-la em um arquivo .txt.

Regras -> 

### FORMULÁRIO PERGUNTAS: ###

1 - Qual o seu nome completo? 

2 - Qual seu e-mail?

3 - Qual sua idade?

4 - Qual seu whatsapp ou celular?


### RESPOSTAS (exemplo):###

NOME ARQUIVO = 1 - LUCASCARRILHO.txt

Lucas Carrilho

lucas3234@gmail.com

32

32999093874

# Critérios #

 - A idade permitida pelo formulário deve ser de no mínimo 16 anos. Caso a idade esteja abaixo do permitido deve ser 
exibida uma mensagem agradecendo o usuário e pedindo para ele tentar a inscrição novamente quando chegar na idade permitida.


 - Quando o usuário terminar de responder as perguntas o formulário dele deve ser salvo no formato TXT com o nome do arquivo 
no padrão {numeroDoFormulario}-{nomeDoCandidato}.txt dentro de uma pasta da sua escolha dentro em seu sistema operacional.


- O nome do arquivo deve ser salvo em letras maíusculas sem acentos e sem espaços para evitar problemas de encoding quando 
rodar em sistemas operacionais diferentes.


 - Como informado deve possibilitar o usuário adicionar uma nova pergunta ao formulário e salvar o txt formulario.txt 
com essa alteração. Ao criar a mesma pergunta novamente o sistema não deve duplicar a pergunta dentro do txt e deve manter apenas uma.


 - Como informado deve possibilitar o usuário remover uma pergunta do formulário e salvar o txt formulario.txt com essa alteração.
não deve ser possível apagar as perguntas de 1 a 4 pois elas são padrão e não podem ser removidas.


 - O programa deve ter uma opção dentro desse menu que permita o administrador listar os nomes dos candidatos cadastrados 
agrupados por idade para fins de análise. Também deve existir uma outra opção que exiba a quantidade de candidatos por idade.


- O programa deve ter uma funcionalidade que permita o administrador encontrar formulários duplicados e apresentá-los em tela.
Um nome duplicado pode ser um homônimo, porém um e-mail + nome iguais com certeza indica uma duplicidade.


- O programa deve permitir realizar a busca se um candidato está cadastrado nos formulários existentes. 
Lembre-se que o que define a unicidade de um cadastro no seu sistema é o nome do candidato + email.


- Ao finalizar a ação de uma das opções dos menus o usuário deve ser enviado para o menu principal novamente.

# EVOLUÇÕES (TODO)

1 - Aplicar boas práticas de programação (renomear arquivos e funcoes).

2 - Criar testes unitários para todos campos.

3 - Criar GOROUTINES na parte de leituras de arquivos do diretório
 
4 - Orientar a OO com uso de mais structs para candidatos e perguntas.

5 - Verificar situação dos pacotes para não acontecer problemas com import cycle.

6 - Integrar com LOCALSTACK (apresentar instalação e uso etc - utilizar bucket S3 simulado para salvar txt.)

7 - Criar Dockerfile para rodar o programa.


## SPRINT 2 (sem previsão)

 - USAR O TEMPLATE PARA TESTAR COM E SEM GOROUTINS PRA VER O DESEMPENHO.

 - INTEGRAR COM DATABASE, CRIAR ENDPOINTS UTILIZANDO ALGUMA API, DOCUMENTAÇÃO

 - MODIFICAR TODO O PROJETO PARA INGLES

 - COLOCAR LOGS EM TODO O PROJETO PARA PODER VER EM MONITORAMENTO.

 - MAPEAR DE TXT PARA JSON
>>>>>>> new/inicial
