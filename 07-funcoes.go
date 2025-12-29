//go:build exercicio07 || ignore
// +build exercicio07 ignore

package main

/*
EXERCÍCIO 07: Funções (declaração, múltiplos retornos, variádicos)
Link do exercício: roadmap-com-desafios-golang.md - Etapa 1: Básico - Funções

OBJETIVO:
Aprender a criar e usar funções em Go, incluindo múltiplos retornos e parâmetros variádicos.

CONCEITOS:
DECLARAÇÃO DE FUNÇÃO:
- func nomeFuncao(parametros) tipoRetorno { }
- func nomeFuncao(parametros) (tipo1, tipo2) { }  // Múltiplos retornos
- Parâmetros do mesmo tipo: func soma(a, b int) int

MÚLTIPLOS RETORNOS:
- Go permite retornar múltiplos valores
- Nomear retornos: func calc() (soma int, produto int)
- Ignorar valores: _, resultado := funcao()

PARÂMETROS VARIÁDICOS:
- ...tipo: Aceita número variável de argumentos
- Acessado como slice dentro da função
- Deve ser o último parâmetro

FUNÇÕES COMO VALORES:
- Funções são cidadãos de primeira classe
- Podem ser atribuídas a variáveis
- Podem ser passadas como parâmetros

EXEMPLOS DE APLICAÇÃO:

1. Função simples:
   func soma(a, b int) int {
       return a + b
   }

   resultado := soma(5, 3)  // 8

2. Função com múltiplos retornos:
   func divide(a, b int) (int, int) {
       return a / b, a % b
   }

   quociente, resto := divide(10, 3)  // 3, 1

3. Retornos nomeados:
   func calcula(x, y int) (soma int, produto int) {
       soma = x + y
       produto = x * y
       return  // Retorna automaticamente soma e produto
   }

4. Parâmetros variádicos:
   func soma(nums ...int) int {
       total := 0
       for _, n := range nums {
           total += n
       }
       return total
   }

   resultado := soma(1, 2, 3, 4, 5)  // 15
   resultado2 := soma(10, 20)        // 30

5. Função como valor:
   multiplica := func(a, b int) int {
       return a * b
   }

   resultado := multiplica(4, 5)  // 20

6. Função anônima:
   func() {
       fmt.Println("Função anônima executada")
   }()

7. Ignorar retornos:
   _, erro := funcaoQueRetornaDoisValores()
   // Ignora o primeiro retorno, usa apenas o erro

DESAFIO (User Story):
Como contador,
eu quero uma função reutilizável para calcular impostos sobre múltiplos itens com taxas variáveis,
para que eu agilize cálculos financeiros repetitivos.

Critérios de Aceitação:
- Função com variádicos para itens
- Retornar múltiplos valores (total, imposto)
- Chamar e exibir resultados

INSTRUÇÕES:
1. Crie uma função que aceite múltiplos valores de itens (variádico)
2. A função deve calcular o total e o imposto (ex: 10% do total)
3. Retorne ambos os valores (total, imposto)
4. Chame a função com diferentes valores e exiba os resultados

OPCIONAL (Recomendado para prática avançada):
- Considere criar uma struct "ResultadoImposto" para retornar os valores
- Isso torna o código mais legível: type ResultadoImposto struct { Total float64; Imposto float64 }
- Em vez de múltiplos retornos, retorne: return ResultadoImposto{Total: total, Imposto: imposto}
*/

// TODO: Crie a função calcularImposto aqui
// Ela deve aceitar múltiplos valores (variádico) e retornar (total, imposto)
// Exemplo: func calcularImposto(itens ...float64) (float64, float64)

func main() {
	// TODO: Chame a função calcularImposto com alguns valores
	// Exemplo: total, imposto := calcularImposto(100.0, 200.0, 150.0)
	// Exiba os resultados formatados
	
	// Dica: Use variádicos: func calcularImposto(itens ...float64)
	// Dica: Calcule total somando todos os itens
	// Dica: Calcule imposto como 10% do total
	// Dica: Retorne (total, imposto)
	
	// Exemplo de estrutura:
	// func calcularImposto(itens ...float64) (float64, float64) {
	//     total := 0.0
	//     for _, item := range itens {
	//         total += item
	//     }
	//     imposto := total * 0.10
	//     return total, imposto
	// }
}



