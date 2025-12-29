// go:build exercicio01 || ignore
// +build exercicio01 ignore

//go:build ignore
// +build ignore

package main

/*
EXERCÍCIO 01: Instalação e Ambiente
Link do exercício: roadmap-com-desafios-golang.md - Etapa 1: Básico - Instalação e Ambiente

OBJETIVO:
Configurar o ambiente Go e criar seu primeiro programa.

CONCEITOS:
- Package main: Define que este é um programa executável
- Import: Importa pacotes da biblioteca padrão do Go
- func main(): Ponto de entrada do programa
- fmt.Println(): Função para imprimir no console

EXEMPLOS DE APLICAÇÃO:

1. Programa básico "Hello World":
   package main
   import "fmt"
   func main() {
       fmt.Println("Olá, Golang!")
   }

2. Verificar versão do Go:
   No terminal: go version

3. Executar programa:
   go run 01-instalacao-e-ambiente.go

4. Compilar programa:
   go build 01-instalacao-e-ambiente.go
   ./01-instalacao-e-ambiente (Linux/Mac) ou 01-instalacao-e-ambiente.exe (Windows)

DESAFIO (User Story):
Como desenvolvedor iniciante,
eu quero configurar o ambiente Go e rodar um programa simples,
para que eu possa começar a codar sem problemas de setup.

Critérios de Aceitação:
- Instalar Go e verificar com `go version`
- Criar e rodar um arquivo .go que imprime uma mensagem
- Configurar VS Code com extensão Go para autocomplete

INSTRUÇÕES:
1. Certifique-se de ter o Go instalado (go.dev)
2. Configure o VS Code com a extensão Go
3. Crie um programa que imprima uma mensagem de boas-vindas
4. Execute o programa com: go run 01-instalacao-e-ambiente.go
*/

import "fmt"

func main() {
	// TODO: Implemente aqui seu primeiro programa Go
	// Imprima uma mensagem de boas-vindas usando fmt.Println()
	
	// Exemplo:
	fmt.Println("Olá, Golang!")
}

