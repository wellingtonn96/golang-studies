//go:build exercicio15 || ignore
// +build exercicio15 ignore

package main

/*
EXERCÍCIO 15: Concorrência Básica - Goroutines e Channels
Link do exercício: roadmap-com-desafios-golang.md - Etapa 2: Intermediário - Concorrência Básica

OBJETIVO:
Aprender a criar goroutines e usar channels para comunicação entre goroutines.

CONCEITOS:
GOROUTINES:
- go funcao(): Executa função em goroutine separada
- Leves: milhares podem rodar simultaneamente
- Não retornam valores diretamente
- Programa não espera goroutines terminar (sem sincronização)

CHANNELS:
- make(chan tipo): Cria channel
- ch <- valor: Envia valor para channel
- valor := <-ch: Recebe valor do channel
- Bloqueia até haver comunicação
- Channels são seguros para comunicação entre goroutines

TIPOS DE CHANNELS:
- Unbuffered: make(chan tipo) - Sincroniza sender e receiver
- Buffered: make(chan tipo, tamanho) - Buffer permite múltiplos valores

EXEMPLOS DE APLICAÇÃO:

1. Goroutine Básica:
   func tarefa() {
       fmt.Println("Executando tarefa")
   }
   
   go tarefa()
   time.Sleep(time.Second)  // Aguarda goroutine terminar

2. Channel Básico:
   ch := make(chan string)
   
   go func() {
       ch <- "mensagem"
   }()
   
   msg := <-ch
   fmt.Println(msg)  // mensagem

3. Channel com Buffer:
   ch := make(chan int, 2)  // Buffer de 2
   ch <- 1
   ch <- 2
   // Não bloqueia até buffer encher
   
   fmt.Println(<-ch)  // 1
   fmt.Println(<-ch)  // 2

4. Múltiplas Goroutines:
   ch := make(chan int)
   
   for i := 0; i < 3; i++ {
       go func(id int) {
           ch <- id
       }(i)
   }
   
   for i := 0; i < 3; i++ {
       fmt.Println(<-ch)
   }

5. Channel Bidirecional vs Unidirecional:
   // Bidirecional (padrão):
   ch := make(chan int)
   
   // Send-only:
   func enviar(ch chan<- int) {
       ch <- 42
   }
   
   // Receive-only:
   func receber(ch <-chan int) {
       valor := <-ch
       fmt.Println(valor)
   }

6. Select (multiplexação):
   ch1 := make(chan string)
   ch2 := make(chan string)
   
   go func() { ch1 <- "um" }()
   go func() { ch2 <- "dois" }()
   
   select {
   case msg1 := <-ch1:
       fmt.Println(msg1)
   case msg2 := <-ch2:
       fmt.Println(msg2)
   }

7. Range sobre Channel:
   ch := make(chan int)
   
   go func() {
       for i := 0; i < 5; i++ {
           ch <- i
       }
       close(ch)  // Fecha channel
   }()
   
   for valor := range ch {
       fmt.Println(valor)
   }

8. Timeout com Select:
   ch := make(chan string)
   
   select {
   case msg := <-ch:
       fmt.Println(msg)
   case <-time.After(time.Second):
       fmt.Println("Timeout")
   }

DESAFIO (User Story):
Como desenvolvedor de e-commerce,
eu quero usar goroutines para processar pedidos paralelamente,
para que o sistema escale com múltiplos usuários.

Critérios de Aceitação:
- Lançar goroutine com go
- Usar channel para comunicação
- Aguardar e exibir resultado

INSTRUÇÕES:
1. Crie uma função "ProcessarPedido" que recebe ID do pedido e channel
2. A função deve simular processamento (time.Sleep) e enviar resultado pelo channel
3. No main, crie um channel e lance 3 goroutines para processar pedidos diferentes
4. Receba e exiba os resultados de cada pedido
5. Use time.Sleep ou WaitGroup para aguardar todas as goroutines
*/

import (
	"fmt"
	"time"
)

// TODO: Crie a função ProcessarPedido aqui
// func ProcessarPedido(id int, ch chan string) { ... }

func main() {
	// TODO: Implemente a solução do desafio aqui
	
	// Dica: Use make(chan string) para criar channel
	// Dica: Use go para lançar goroutines
	// Dica: Use <-ch para receber valores
	// Dica: Use time.Sleep para simular processamento
	
	// Exemplo de estrutura:
	// func ProcessarPedido(id int, ch chan string) {
	//     time.Sleep(time.Second)  // Simula processamento
	//     ch <- fmt.Sprintf("Pedido %d processado", id)
	// }
	// 
	// ch := make(chan string)
	// 
	// go ProcessarPedido(1, ch)
	// go ProcessarPedido(2, ch)
	// go ProcessarPedido(3, ch)
	// 
	// for i := 0; i < 3; i++ {
	//     resultado := <-ch
	//     fmt.Println(resultado)
	// }
}

