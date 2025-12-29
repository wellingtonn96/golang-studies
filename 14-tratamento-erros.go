//go:build exercicio14 || ignore
// +build exercicio14 ignore

package main

/*
EXERCÍCIO 14: Tratamento de Erros (error, panic, recover)
Link do exercício: roadmap-com-desafios-golang.md - Etapa 2: Intermediário - Tratamento de Erros

OBJETIVO:
Aprender a tratar erros de forma idiomática em Go e usar panic/recover quando necessário.

CONCEITOS:
ERROR INTERFACE:
- type error interface { Error() string }
- Funções retornam (resultado, error)
- Sempre verifique erros: if err != nil { }
- errors.New("mensagem"): Cria novo erro
- fmt.Errorf("formato", args): Erro formatado

PANIC E RECOVER:
- panic(valor): Interrompe execução normal
- recover(): Recupera de panic (só funciona em defer)
- defer: Executa função ao final do escopo
- Use panic apenas para erros realmente excepcionais

IDIOMAS GO:
- "Errors are values": Trate erros como valores normais
- Não ignore erros: sempre verifique
- Retorne erros, não panics (na maioria dos casos)

EXEMPLOS DE APLICAÇÃO:

1. Retornar e Verificar Erro:
   func divide(a, b int) (int, error) {
       if b == 0 {
           return 0, errors.New("divisão por zero")
       }
       return a / b, nil
   }
   
   resultado, err := divide(10, 0)
   if err != nil {
       fmt.Println("Erro:", err)
       return
   }
   fmt.Println("Resultado:", resultado)

2. Erro Formatado:
   func validarIdade(idade int) error {
       if idade < 0 {
           return fmt.Errorf("idade inválida: %d (deve ser positiva)", idade)
       }
       if idade > 150 {
           return fmt.Errorf("idade inválida: %d (muito alta)", idade)
       }
       return nil
   }

3. Múltiplos Erros:
   func processar(valor int) (int, error) {
       if valor < 0 {
           return 0, errors.New("valor negativo")
       }
       if valor > 100 {
           return 0, errors.New("valor muito alto")
       }
       return valor * 2, nil
   }

4. Defer Básico:
   func exemplo() {
       defer fmt.Println("Sempre executa no final")
       fmt.Println("Primeiro")
       fmt.Println("Segundo")
   }
   // Saída: Primeiro, Segundo, Sempre executa no final

5. Defer com Recover:
   func exemploSeguro() {
       defer func() {
           if r := recover(); r != nil {
               fmt.Println("Recuperado de panic:", r)
           }
       }()
       
       panic("algo deu errado")
       fmt.Println("Isso não será executado")
   }

6. Panic (use com cuidado):
   func buscarUsuario(id int) *Usuario {
       usuario := buscarNoBanco(id)
       if usuario == nil {
           panic("usuário não encontrado")
       }
       return usuario
   }

7. Wrapping Errors (Go 1.13+):
   import "fmt"
   
   func processarArquivo(nome string) error {
       dados, err := lerArquivo(nome)
       if err != nil {
           return fmt.Errorf("erro ao processar arquivo %s: %w", nome, err)
       }
       // processar dados...
       return nil
   }

8. Verificar Tipo de Erro:
   var ErrDivisaoPorZero = errors.New("divisão por zero")
   
   func divide(a, b int) (int, error) {
       if b == 0 {
           return 0, ErrDivisaoPorZero
       }
       return a / b, nil
   }
   
   _, err := divide(10, 0)
   if err == ErrDivisaoPorZero {
       fmt.Println("Erro específico capturado")
   }

DESAFIO (User Story):
Como responsável por sistema de pagamentos,
eu quero tratar erros de transação e recuperar de panics,
para que o sistema seja robusto em operações financeiras.

Critérios de Aceitação:
- Retornar error em função
- Verificar err != nil
- Usar defer recover para panic

INSTRUÇÕES:
1. Crie uma função "ProcessarPagamento" que recebe valor e saldo
2. A função deve retornar (sucesso bool, error)
3. Valide: valor deve ser positivo, saldo deve ser suficiente
4. Crie uma função "ProcessarTransacao" que chama ProcessarPagamento
5. Use defer recover em ProcessarTransacao para capturar panics
6. Teste com casos válidos e inválidos
*/

import (
	"errors"
	"fmt"
)

// TODO: Crie a função ProcessarPagamento aqui
// func ProcessarPagamento(valor, saldo float64) (bool, error) { ... }

func main() {
	// TODO: Implemente ProcessarTransacao com defer recover
	// func ProcessarTransacao(valor, saldo float64) { ... }
	
	// Dica: Retorne errors.New() para erros específicos
	// Dica: Use defer func() { recover() } para capturar panics
	// Dica: Sempre verifique err != nil
	
	// Exemplo de estrutura:
	// func ProcessarPagamento(valor, saldo float64) (bool, error) {
	//     if valor <= 0 {
	//         return false, errors.New("valor deve ser positivo")
	//     }
	//     if saldo < valor {
	//         return false, fmt.Errorf("saldo insuficiente: R$ %.2f < R$ %.2f", saldo, valor)
	//     }
	//     return true, nil
	// }
	// 
	// func ProcessarTransacao(valor, saldo float64) {
	//     defer func() {
	//         if r := recover(); r != nil {
	//             fmt.Printf("Panic recuperado: %v\n", r)
	//         }
	//     }()
	//     
	//     sucesso, err := ProcessarPagamento(valor, saldo)
	//     if err != nil {
	//         fmt.Printf("Erro no pagamento: %v\n", err)
	//         return
	//     }
	//     
	//     if sucesso {
	//         fmt.Println("Pagamento processado com sucesso!")
	//     }
	// }
	// 
	// ProcessarTransacao(100.0, 500.0)  // Sucesso
	// ProcessarTransacao(100.0, 50.0)  // Erro: saldo insuficiente
	// ProcessarTransacao(-10.0, 500.0) // Erro: valor inválido
}

