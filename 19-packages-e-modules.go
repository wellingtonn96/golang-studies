//go:build exercicio19 || ignore
// +build exercicio19 ignore

package main

/*
EXERCÍCIO 19: Packages e Modules
Link do exercício: roadmap-com-desafios-golang.md - Etapa 2: Intermediário - Organização de Código

OBJETIVO:
Aprender a organizar código em packages e criar modules reutilizáveis.

CONCEITOS:
PACKAGES:
- package nome: Define o package
- Nomes em minúsculas: privados (não exportados)
- Nomes em maiúsculas: públicos (exportados)
- Cada diretório = um package
- package main: Programa executável

MODULES:
- go mod init nome: Cria novo module
- go.mod: Arquivo de configuração do module
- go.sum: Checksums de dependências
- go get: Adiciona dependência
- go mod tidy: Limpa dependências não usadas

IMPORTS:
- import "pacote": Import padrão
- import alias "pacote": Import com alias
- import _ "pacote": Import apenas para side effects
- import . "pacote": Import sem qualificador

EXEMPLOS DE APLICAÇÃO:

1. Criar Module:
   go mod init meuprojeto
   // Cria go.mod

2. Package Exportado:
   // Em utils/math.go
   package utils
   
   // Exportado (maiúscula):
   func Soma(a, b int) int {
       return a + b
   }
   
   // Privado (minúscula):
   func subtracao(a, b int) int {
       return a - b
   }

3. Importar Package Local:
   // Em main.go
   package main
   
   import (
       "fmt"
       "meuprojeto/utils"
   )
   
   func main() {
       resultado := utils.Soma(5, 3)
       fmt.Println(resultado)
   }

4. Import com Alias:
   import (
       m "meuprojeto/utils"
   )
   
   resultado := m.Soma(5, 3)

5. Import para Side Effects:
   import _ "database/sql/driver"
   // Executa init() do package

6. Variáveis Exportadas:
   package config
   
   var APIKey string = "secret"  // Exportado
   var internalToken string       // Privado

7. go get Dependência:
   go get github.com/gin-gonic/gin
   // Adiciona ao go.mod

8. Estrutura de Projeto:
   meuprojeto/
   ├── go.mod
   ├── go.sum
   ├── main.go
   ├── utils/
   │   └── math.go
   └── models/
       └── user.go

DESAFIO (User Story):
Como desenvolvedor de ferramenta financeira,
eu quero separar lógica em packages e modules reutilizáveis,
para que equipes colaborem sem duplicação.

Critérios de Aceitação:
- Criar module com `go mod init`
- Package separado com função exportada
- Importar e usar

INSTRUÇÕES:
1. Crie um novo module: go mod init calculadora
2. Crie diretório "calculos" com arquivo calculos.go
3. No package calculos, crie funções exportadas:
   - CalcularJuros(capital, taxa, tempo float64) float64
   - CalcularDesconto(valor, percentual float64) float64
4. No main.go, importe e use essas funções
5. Execute: go run main.go
*/

import "fmt"

// NOTA: Para este exercício, você precisa criar a estrutura:
//
// calculadora/
// ├── go.mod (criado com: go mod init calculadora)
// ├── main.go (este arquivo)
// └── calculos/
//     └── calculos.go
//
// Exemplo de calculos/calculos.go:
//
// package calculos
//
// func CalcularJuros(capital, taxa, tempo float64) float64 {
//     return capital * (taxa / 100) * tempo
// }
//
// func CalcularDesconto(valor, percentual float64) float64 {
//     return valor * (percentual / 100)
// }

func main() {
	// TODO: Importe o package calculos e use as funções
	// import "calculadora/calculos"
	//
	// juros := calculos.CalcularJuros(1000.0, 5.0, 12.0)
	// desconto := calculos.CalcularDesconto(500.0, 10.0)
	//
	// fmt.Printf("Juros: R$ %.2f\n", juros)
	// fmt.Printf("Desconto: R$ %.2f\n", desconto)
	
	fmt.Println("Execute: go mod init calculadora")
	fmt.Println("Crie o diretório calculos/ com as funções exportadas")
	fmt.Println("Importe e use no main.go")
}



