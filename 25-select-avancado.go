//go:build exercicio25 || ignore
// +build exercicio25 ignore

package main

/*
EXERCÍCIO 25: Select Avançado e Multiplexação
Link do exercício: Conceito avançado essencial para desenvolvimento real

OBJETIVO:
Dominar o uso de select para multiplexação de channels e casos avançados.

CONCEITOS:
SELECT STATEMENT:
- Multiplexa múltiplos channels
- Bloqueia até um case esteja pronto
- Se múltiplos prontos, escolhe aleatoriamente
- default: Não bloqueia (non-blocking)

CASOS DE USO:
- Timeout em operações
- Cancelamento com context
- Priorização de channels
- Non-blocking operations

EXEMPLOS DE APLICAÇÃO:

1. Select Básico:
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

2. Select com Timeout:
   ch := make(chan string)
   
   select {
   case msg := <-ch:
       fmt.Println(msg)
   case <-time.After(2 * time.Second):
       fmt.Println("Timeout!")
   }

3. Select com Default (Non-blocking):
   select {
   case msg := <-ch:
       fmt.Println(msg)
   default:
       fmt.Println("Nenhuma mensagem disponível")
   }

4. Select com Context:
   ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
   defer cancel()
   
   ch := make(chan string)
   
   select {
   case msg := <-ch:
       fmt.Println(msg)
   case <-ctx.Done():
       fmt.Println("Cancelado:", ctx.Err())
   }

5. Select em Loop:
   for {
       select {
       case msg := <-ch1:
           fmt.Println("Ch1:", msg)
       case msg := <-ch2:
           fmt.Println("Ch2:", msg)
       case <-done:
           return
       }
   }

6. Select com Send e Receive:
   ch := make(chan int, 1)
   
   select {
   case ch <- 42:
       fmt.Println("Enviado")
   case val := <-ch:
       fmt.Println("Recebido:", val)
   default:
       fmt.Println("Nada aconteceu")
   }

7. Select com Múltiplos Timeouts:
   ch := make(chan string)
   
   select {
   case msg := <-ch:
       fmt.Println(msg)
   case <-time.After(1 * time.Second):
       fmt.Println("Timeout 1s")
   case <-time.After(5 * time.Second):
       fmt.Println("Timeout 5s")
   }

8. Select com Prioridade:
   ch1 := make(chan string)  // Alta prioridade
   ch2 := make(chan string)  // Baixa prioridade
   
   for {
       select {
       case msg := <-ch1:
           fmt.Println("Alta:", msg)
       default:
           select {
           case msg := <-ch1:
               fmt.Println("Alta:", msg)
           case msg := <-ch2:
               fmt.Println("Baixa:", msg)
           }
       }
   }

DESAFIO:
Como desenvolvedor de sistema de monitoramento,
eu quero processar eventos de múltiplas fontes com timeout,
para que eu responda rapidamente a eventos críticos.

Critérios de Aceitação:
- Processar eventos de múltiplos channels
- Implementar timeout para operações
- Priorizar eventos críticos
- Usar context para cancelamento

INSTRUÇÕES:
1. Crie 3 channels: eventos críticos, eventos normais, e done
2. Crie goroutines que enviam eventos para cada channel
3. Use select para processar eventos (priorizar críticos)
4. Implemente timeout de 3 segundos
5. Use context para cancelamento
6. Processe eventos até timeout ou cancelamento
*/

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// TODO: Implemente a solução do desafio aqui
	
	// Dica: Use select para multiplexar channels
	// Dica: Use time.After para timeout
	// Dica: Use context.WithTimeout para cancelamento
	// Dica: Use default para non-blocking quando necessário
	
	// Exemplo de estrutura:
	// criticos := make(chan string, 10)
	// normais := make(chan string, 10)
	// done := make(chan bool)
	// 
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()
	// 
	// // Goroutines enviando eventos
	// go func() {
	//     for i := 0; i < 5; i++ {
	//         criticos <- fmt.Sprintf("Crítico %d", i)
	//         time.Sleep(200 * time.Millisecond)
	//     }
	// }()
	// 
	// go func() {
	//     for i := 0; i < 10; i++ {
	//         normais <- fmt.Sprintf("Normal %d", i)
	//         time.Sleep(100 * time.Millisecond)
	//     }
	// }()
	// 
	// // Processar eventos
	// for {
	//     select {
	//     case evento := <-criticos:
	//         fmt.Printf("[CRÍTICO] %s\n", evento)
	//     case evento := <-normais:
	//         fmt.Printf("[NORMAL] %s\n", evento)
	//     case <-ctx.Done():
	//         fmt.Println("Timeout ou cancelado")
	//         return
	//     }
	// }
}

