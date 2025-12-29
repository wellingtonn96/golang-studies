//go:build exercicio18 || ignore
// +build exercicio18 ignore

package main

/*
EXERCÍCIO 18: Testes e Benchmarking
Link do exercício: roadmap-com-desafios-golang.md - Etapa 3: Avançado - Testes e Benchmarking

OBJETIVO:
Aprender a escrever testes unitários e benchmarks em Go.

CONCEITOS:
TESTES:
- Arquivo *_test.go: Contém testes
- func TestNome(t *testing.T): Função de teste
- t.Error(), t.Fatal(): Reporta falha
- t.Log(): Log durante teste
- go test: Executa testes
- go test -v: Modo verbose

BENCHMARKS:
- func BenchmarkNome(b *testing.B): Função de benchmark
- b.N: Número de iterações
- go test -bench=.: Executa benchmarks
- go test -bench=. -benchmem: Inclui memória

TABLE-DRIVEN TESTS:
- Padrão comum em Go
- Slice de casos de teste
- Loop sobre casos

EXEMPLOS DE APLICAÇÃO:

1. Teste Básico:
   // Em math_test.go
   package math
   
   import "testing"
   
   func TestSoma(t *testing.T) {
       resultado := Soma(2, 3)
       esperado := 5
       if resultado != esperado {
           t.Errorf("Soma(2, 3) = %d; esperado %d", resultado, esperado)
       }
   }

2. Teste com Múltiplos Casos (Table-Driven):
   func TestSoma(t *testing.T) {
       casos := []struct {
           a, b, esperado int
       }{
           {2, 3, 5},
           {0, 0, 0},
           {-1, 1, 0},
           {10, -5, 5},
       }
       
       for _, c := range casos {
           resultado := Soma(c.a, c.b)
           if resultado != c.esperado {
               t.Errorf("Soma(%d, %d) = %d; esperado %d", c.a, c.b, resultado, c.esperado)
           }
       }
   }

3. Teste com Subtests:
   func TestOperacoes(t *testing.T) {
       t.Run("Soma", func(t *testing.T) {
           if Soma(2, 3) != 5 {
               t.Error("Falha na soma")
           }
       })
       
       t.Run("Multiplicacao", func(t *testing.T) {
           if Multiplicacao(2, 3) != 6 {
               t.Error("Falha na multiplicação")
           }
       })
   }

4. Benchmark Básico:
   func BenchmarkSoma(b *testing.B) {
       for i := 0; i < b.N; i++ {
           Soma(100, 200)
       }
   }

5. Benchmark Comparativo:
   func BenchmarkSoma(b *testing.B) {
       b.Run("Pequenos", func(b *testing.B) {
           for i := 0; i < b.N; i++ {
               Soma(1, 2)
           }
       })
       
       b.Run("Grandes", func(b *testing.B) {
           for i := 0; i < b.N; i++ {
               Soma(1000000, 2000000)
           }
       })
   }

6. Teste de Erro:
   func TestDivisaoPorZero(t *testing.T) {
       _, err := Dividir(10, 0)
       if err == nil {
           t.Fatal("Esperava erro, mas não houve")
       }
   }

7. Teste de Helpers:
   func assertEqual(t *testing.T, resultado, esperado int) {
       t.Helper()  // Marca como helper
       if resultado != esperado {
           t.Errorf("Esperado %d, obteve %d", esperado, resultado)
       }
   }
   
   func TestExemplo(t *testing.T) {
       assertEqual(t, Soma(2, 3), 5)
   }

DESAFIO (User Story):
Como responsável por qualidade,
eu quero escrever testes unitários e benchmarks para funções críticas,
para que garanta precisão e performance.

Critérios de Aceitação:
- Função Test com t.Error
- Benchmark com b.N
- Rodar go test -bench

INSTRUÇÕES:
1. Crie um arquivo 18-testes-e-benchmarking_test.go
2. Implemente função CalcularDesconto(valor, percentual float64) float64
3. Escreva testes para CalcularDesconto usando table-driven tests
4. Escreva benchmark para CalcularDesconto
5. Execute: go test -v
6. Execute: go test -bench=. -benchmem
*/

// Função a ser testada (implemente aqui ou em arquivo separado)
// func CalcularDesconto(valor, percentual float64) float64 {
//     return valor * (percentual / 100)
// }

// NOTA: Os testes devem estar em um arquivo separado:
// 18-testes-e-benchmarking_test.go
//
// Exemplo de conteúdo do arquivo de teste:
//
// package main
//
// import "testing"
//
// func TestCalcularDesconto(t *testing.T) {
//     casos := []struct {
//         nome      string
//         valor     float64
//         percentual float64
//         esperado  float64
//     }{
//         {"Desconto 10%", 100.0, 10.0, 90.0},
//         {"Desconto 50%", 200.0, 50.0, 100.0},
//         {"Sem desconto", 100.0, 0.0, 100.0},
//         {"Desconto total", 100.0, 100.0, 0.0},
//     }
//
//     for _, c := range casos {
//         t.Run(c.nome, func(t *testing.T) {
//             resultado := CalcularDesconto(c.valor, c.percentual)
//             if resultado != c.esperado {
//                 t.Errorf("CalcularDesconto(%.2f, %.2f) = %.2f; esperado %.2f",
//                     c.valor, c.percentual, resultado, c.esperado)
//             }
//         })
//     }
// }
//
// func BenchmarkCalcularDesconto(b *testing.B) {
//     for i := 0; i < b.N; i++ {
//         CalcularDesconto(1000.0, 15.0)
//     }
// }

func main() {
	// Este arquivo contém apenas a função a ser testada
	// Os testes devem estar em arquivo *_test.go separado
}

