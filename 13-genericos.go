//go:build exercicio13 || ignore
// +build exercicio13 ignore

package main

/*
EXERCÍCIO 13: Genéricos (Go 1.18+)
Link do exercício: roadmap-com-desafios-golang.md - Etapa 2: Intermediário - Genéricos

OBJETIVO:
Aprender a usar genéricos (type parameters) para criar código reutilizável.

CONCEITOS:
GENÉRICOS:
- Introduzidos no Go 1.18
- Permite funções e tipos parametrizados por tipo
- Sintaxe: func Nome[T Tipo]() { }
- Constraints: Limita tipos aceitos

TYPE PARAMETERS:
- [T Tipo]: Parâmetro de tipo
- [T, U Tipo]: Múltiplos parâmetros
- Constraints: [T comparable], [T ~int], [T any]

CONSTRAINTS:
- any: Qualquer tipo (equivalente a interface{})
- comparable: Tipos que podem ser comparados (==, !=)
- ~tipo: Tipos com tipo subjacente específico
- Interface como constraint: Define métodos necessários

EXEMPLOS DE APLICAÇÃO:

1. Função Genérica Simples:
   func Max[T comparable](a, b T) T {
       if a > b {
           return a
       }
       return b
   }
   
   // Erro: comparable não permite >
   // Use constraint customizada

2. Função Genérica com Constraint:
   type Numeric interface {
       ~int | ~int8 | ~int16 | ~int32 | ~int64 |
       ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
       ~float32 | ~float64
   }
   
   func Max[T Numeric](a, b T) T {
       if a > b {
           return a
       }
       return b
   }
   
   fmt.Println(Max(5, 3))        // 5
   fmt.Println(Max(4.5, 6.7))   // 6.7

3. Função Genérica com Slice:
   func Soma[T Numeric](nums []T) T {
       var total T
       for _, n := range nums {
           total += n
       }
       return total
   }
   
   fmt.Println(Soma([]int{1, 2, 3}))        // 6
   fmt.Println(Soma([]float64{1.5, 2.5}))   // 4.0

4. Tipo Genérico:
   type Stack[T any] struct {
       items []T
   }
   
   func (s *Stack[T]) Push(item T) {
       s.items = append(s.items, item)
   }
   
   func (s *Stack[T]) Pop() (T, bool) {
       if len(s.items) == 0 {
           var zero T
           return zero, false
       }
       item := s.items[len(s.items)-1]
       s.items = s.items[:len(s.items)-1]
       return item, true
   }
   
   stack := Stack[int]{}
   stack.Push(1)
   stack.Push(2)
   item, _ := stack.Pop()
   fmt.Println(item)  // 2

5. Múltiplos Type Parameters:
   func Trocar[T, U any](a T, b U) (U, T) {
       return b, a
   }
   
   x, y := Trocar(10, "texto")
   fmt.Println(x, y)  // texto 10

6. Constraint com Métodos:
   type Stringer interface {
       String() string
   }
   
   func Print[T Stringer](item T) {
       fmt.Println(item.String())
   }

DESAFIO (User Story):
Como analista de dados,
eu quero funções genéricas para operar em tipos numéricos diferentes,
para que eu reutilize código em relatórios variados.

Critérios de Aceitação:
- Função genérica com type parameters
- Constraints para tipos
- Testar com int e float

INSTRUÇÕES:
1. Defina uma constraint "Numeric" que aceite int e float64
2. Crie uma função genérica "Media" que calcule a média de um slice numérico
3. Crie uma função genérica "EncontrarMax" que encontre o maior valor
4. Teste ambas as funções com slices de int e float64
5. Exiba os resultados
*/

import "fmt"

// TODO: Defina a constraint e funções genéricas aqui
// type Numeric interface { ... }
// func Media[T Numeric](nums []T) float64 { ... }
// func EncontrarMax[T Numeric](nums []T) T { ... }

func main() {
	// TODO: Implemente a solução do desafio aqui
	
	// Dica: Use constraint ~int | ~float64
	// Dica: Para média, pode precisar converter para float64
	// Dica: Use len(nums) para contar elementos
	
	// Exemplo de estrutura:
	// type Numeric interface {
	//     ~int | ~float64
	// }
	// 
	// func Media[T Numeric](nums []T) float64 {
	//     var soma float64
	//     for _, n := range nums {
	//         soma += float64(n)
	//     }
	//     return soma / float64(len(nums))
	// }
	// 
	// func EncontrarMax[T Numeric](nums []T) T {
	//     if len(nums) == 0 {
	//         var zero T
	//         return zero
	//     }
	//     max := nums[0]
	//     for _, n := range nums[1:] {
	//         if float64(n) > float64(max) {
	//             max = n
	//         }
	//     }
	//     return max
	// }
	// 
	// ints := []int{10, 20, 30, 40}
	// floats := []float64{1.5, 2.5, 3.5, 4.5}
	// 
	// fmt.Printf("Média dos ints: %.2f\n", Media(ints))
	// fmt.Printf("Máximo dos ints: %d\n", EncontrarMax(ints))
	// fmt.Printf("Média dos floats: %.2f\n", Media(floats))
	// fmt.Printf("Máximo dos floats: %.2f\n", EncontrarMax(floats))
}



