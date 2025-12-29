//go:build exercicio00 || ignore
// +build exercicio00 ignore

package main

/*
EXERCÍCIO 09: Escopo e Ponteiros Básicos
Link do exercício: roadmap-com-desafios-golang.md - Etapa 1: Básico - Escopo e Ponteiros Básicos

OBJETIVO:
Entender escopo de variáveis e como usar ponteiros em Go.

CONCEITOS:
ESCOPO:
- Variáveis declaradas em um bloco só existem dentro desse bloco
- Variáveis externas são acessíveis em blocos internos
- Shadowing: Variável interna pode "esconder" variável externa com mesmo nome

PONTEIROS:
- &variavel: Operador de endereço (retorna ponteiro)
- *ponteiro: Operador de dereferência (acessa valor apontado)
- *tipo: Tipo ponteiro (ex: *int)
- nil: Valor zero de ponteiros (ponteiro não inicializado)
- new(tipo): Cria ponteiro para novo valor do tipo

PASSAGEM POR VALOR vs REFERÊNCIA:
- Go passa por valor por padrão (cópia)
- Ponteiros permitem passar por referência (modifica original)
- Slices, maps e channels são passados por referência implicitamente

EXEMPLOS DE APLICAÇÃO:

1. Escopo de Variáveis:
   x := 10  // Escopo do main
   if true {
       y := 20  // Escopo do if
       fmt.Println(x, y)  // 10 20 (x acessível)
   }
   // fmt.Println(y)  // ERRO: y não existe aqui
   
   // Shadowing:
   x := 10
   if true {
       x := 20  // Nova variável x (shadowing)
       fmt.Println(x)  // 20
   }
   fmt.Println(x)  // 10 (x original não foi alterado)

2. Ponteiros Básicos:
   x := 10
   ptr := &x        // ptr é um *int (ponteiro para int)
   fmt.Println(ptr) // Endereço de memória (ex: 0xc0000140a0)
   fmt.Println(*ptr) // 10 (valor apontado)
   
   *ptr = 20        // Modifica x através do ponteiro
   fmt.Println(x)   // 20

3. Criar Ponteiros:
   // Método 1: Usando &
   x := 10
   ptr := &x
   
   // Método 2: Usando new()
   ptr2 := new(int)  // Cria ponteiro para int com valor zero (0)
   *ptr2 = 30
   fmt.Println(*ptr2)  // 30

4. Função com Ponteiro:
   func dobrar(x *int) {
       *x = *x * 2
   }
   
   numero := 5
   dobrar(&numero)
   fmt.Println(numero)  // 10

5. Verificar se Ponteiro é nil:
   var ptr *int
   if ptr == nil {
       fmt.Println("Ponteiro não inicializado")
   }

6. Ponteiros em Structs:
   type Pessoa struct {
       Nome string
   }
   
   p := Pessoa{Nome: "João"}
   ptr := &p
   ptr.Nome = "Maria"  // Go permite acesso direto
   // Equivale a: (*ptr).Nome = "Maria"

DESAFIO (User Story):
Como desenvolvedor de um app financeiro,
eu quero organizar variáveis em escopos e usar ponteiros para dados mutáveis,
para que dados sensíveis não vazem e eu gerencie memória eficientemente.

Critérios de Aceitação:
- Variáveis em escopos diferentes
- Usar ponteiro para alterar valor
- Evitar acesso fora de escopo

INSTRUÇÕES:
1. Crie uma variável saldo no escopo principal
2. Crie uma função que recebe um ponteiro para saldo e adiciona valor
3. Dentro de um bloco if, crie uma variável temporária (escopo local)
4. Use a função para modificar o saldo através do ponteiro
5. Tente acessar a variável temporária fora do bloco (deve dar erro)
6. Exiba o saldo final
*/

import "fmt"

// TODO: Crie uma função que modifica saldo através de ponteiro
// func adicionarSaldo(saldo *float64, valor float64) { ... }

func main() {
	// TODO: Implemente a solução do desafio aqui
	
	// Dica: Declare saldo no escopo do main
	// Dica: Use &saldo para passar ponteiro
	// Dica: Crie variável temporária dentro de um bloco if
	// Dica: Tente acessar variável temporária fora do bloco (comente para não dar erro)
	
	// Exemplo de estrutura:
	// saldo := 1000.0
	// 
	// if true {
	//     valorTemporario := 500.0  // Escopo local
	//     fmt.Printf("Valor temporário: %.2f\n", valorTemporario)
	//     adicionarSaldo(&saldo, valorTemporario)
	// }
	// // fmt.Println(valorTemporario)  // ERRO: não existe aqui
	// 
	// fmt.Printf("Saldo final: R$ %.2f\n", saldo)
}



