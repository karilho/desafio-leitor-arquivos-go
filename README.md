# desafio-leitor-arquivos-go
Criação de desafio de ler e escrever arquivos .txt com GO

### Menu Listar Formulários Cadastrados

O programa deve ter uma opção dentro desse menu que permita o administrador listar os nomes dos candidatos cadastrados agrupados por idade para fins de análise.
Também deve existir uma outra opção que exiba a quantidade de candidatos por idade

### Menu Validar Formulários

O programa deve ter uma funcionalidade que permita o administrador encontrar formulários duplicados e apresentá-los em tela. Um nome duplicado pode ser um homônimo, porém um e-mail + nome iguais com certeza indica uma duplicidade.

Dica: Salve o formulário do candidato no txt com uma estrutura que permita você realizar a varredura buscando se o email informado já existe.

### Menu Pesquisar Formulários Cadastrados

O programa deve permitir realizar a busca se um candidato está cadastrado nos formulários existentes. Lembre-se que o que define a unicidade de um cadastro no seu sistema é o nome do candidato + email.


### Outras regras

As classes devem estar divididas em pacotes. As classes que manipulam os arquivos devem ser acessadas apenas por uma outra classe que forneça as funcionalidades para o menu principal.

Cuidado com os padrões camelCase do java. Respeite a padronização.

Pense numa modelagem flexível e nas boas práticas de codificação e clean code.

Ao finalizar a ação de uma das opções dos menus o usuário deve ser enviado para o menu principal

