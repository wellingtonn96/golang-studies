//go:build exercicio10 || ignore
// +build exercicio10 ignore

package main

/*
EXERCÍCIO 10: Structs e Embedding
Link do exercício: roadmap-com-desafios-golang.md - Etapa 2: Intermediário - Structs e Embedding

OBJETIVO:
Aprender a criar estruturas de dados personalizadas e usar embedding para composição.

CONCEITOS:
STRUCTS:
- type Nome struct { campos }: Define uma estrutura
- Campos nomeados com tipos
- Inicialização: Nome{campo1: valor1, campo2: valor2}
- Acesso: instancia.campo
- Campos anônimos: Embedding

EMBEDDING:
- Struct dentro de outra struct sem nome
- Campos são "promovidos" (acessíveis diretamente)
- Permite composição ao invés de herança
- Útil para compartilhar comportamento comum

EXEMPLOS DE APLICAÇÃO:

1. Struct Básica:
   type Pessoa struct {
       Nome  string
       Idade int
   }
   
   p := Pessoa{Nome: "João", Idade: 30}
   fmt.Println(p.Nome)  // João
   
   // Inicialização sem nomes (ordem importa):
   p2 := Pessoa{"Maria", 25}

2. Struct com Campos Anônimos (Embedding):
   type Animal struct {
       Nome string
   }
   
   type Cachorro struct {
       Animal  // Embedding (sem nome de campo)
       Raca    string
   }
   
   c := Cachorro{
       Animal: Animal{Nome: "Rex"},
       Raca:    "Labrador",
   }
   
   // Campos promovidos - acesso direto:
   fmt.Println(c.Nome)  // Rex (vem de Animal)
   fmt.Println(c.Raca)  // Labrador
   
   // Também pode acessar explicitamente:
   fmt.Println(c.Animal.Nome)  // Rex

3. Struct com Métodos (veremos depois):
   type Retangulo struct {
       Largura float64
       Altura  float64
   }
   
   func (r Retangulo) Area() float64 {
       return r.Largura * r.Altura
   }

4. Ponteiros para Structs:
   p := &Pessoa{Nome: "João", Idade: 30}
   p.Idade = 31  // Go permite acesso direto
   // Equivale a: (*p).Idade = 31

5. Struct Anônima:
   pessoa := struct {
       Nome string
       Idade int
   }{
       Nome: "Temporário",
       Idade: 20,
   }

6. Composição Múltipla:
   type Endereco struct {
       Rua    string
       Cidade string
   }
   
   type Contato struct {
       Email string
       Telefone string
   }
   
   type Cliente struct {
       Endereco  // Embedding
       Contato   // Embedding
       Nome      string
   }
   
   c := Cliente{
       Endereco: Endereco{Rua: "Rua A", Cidade: "São Paulo"},
       Contato:  Contato{Email: "cliente@email.com", Telefone: "123456"},
       Nome:     "João",
   }
   
   fmt.Println(c.Rua)     // Promovido de Endereco
   fmt.Println(c.Email)   // Promovido de Contato

DESAFIO (User Story):
Como responsável pelo catálogo de produtos,
eu quero modelar itens com structs e embedding para atributos compartilhados,
para que eu organize dados de forma eficiente em memória.

Critérios de Aceitação:
- Struct base para atributos comuns
- Embedding em struct derivada
- Acessar campos promovidos

INSTRUÇÕES:
1. Crie uma struct base "Produto" com campos: ID, Nome, Preco
2. Crie uma struct "ProdutoEletronico" que embede Produto e adiciona campo "Garantia"
3. Crie uma struct "ProdutoAlimenticio" que embede Produto e adiciona campo "Validade"
4. Crie instâncias de cada tipo
5. Acesse campos promovidos (ID, Nome, Preco) diretamente nas structs derivadas
6. Exiba todas as informações
*/

import "fmt"

// TODO: Defina as structs aqui
// type Produto struct { ... }
// type ProdutoEletronico struct { ... }
// type ProdutoAlimenticio struct { ... }

func main() {
	// TODO: Implemente a solução do desafio aqui
	
	// Dica: Use embedding para Produto nas structs derivadas
	// Dica: Campos de Produto serão promovidos (acessíveis diretamente)
	
	// Exemplo de estrutura:
	// eletronico := ProdutoEletronico{
	//     Produto: Produto{
	//         ID:    1,
	//         Nome:  "Notebook",
	//         Preco: 2500.0,
	//     },
	//     Garantia: "12 meses",
	// }
	// 
	// alimenticio := ProdutoAlimenticio{
	//     Produto: Produto{
	//         ID:    2,
	//         Nome:  "Leite",
	//         Preco: 5.50,
	//     },
	//     Validade: "2024-12-31",
	// }
	// 
	// fmt.Printf("Eletrônico - ID: %d, Nome: %s, Preço: R$ %.2f, Garantia: %s\n",
	//     eletronico.ID, eletronico.Nome, eletronico.Preco, eletronico.Garantia)
	// fmt.Printf("Alimentício - ID: %d, Nome: %s, Preço: R$ %.2f, Validade: %s\n",
	//     alimenticio.ID, alimenticio.Nome, alimenticio.Preco, alimenticio.Validade)
}

