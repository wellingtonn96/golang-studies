//go:build exercicio24 || ignore
// +build exercicio24 ignore

package main

/*
EXERCÍCIO 24: Worker Pools
Link do exercício: Conceito avançado essencial para desenvolvimento real

OBJETIVO:
Aprender a implementar worker pools para processamento paralelo eficiente.

CONCEITOS:
WORKER POOLS:
- Padrão para limitar número de goroutines
- Workers processam tarefas de uma fila
- Controla concorrência e uso de recursos
- Evita criar goroutines ilimitadas

COMPONENTES:
- Job Channel: Fila de tarefas
- Worker Pool: Número fixo de workers
- Result Channel: Canal para resultados
- WaitGroup: Sincronização

BENEFÍCIOS:
- Controle de recursos
- Melhor performance
- Evita sobrecarga do sistema

EXEMPLOS DE APLICAÇÃO:

1. Worker Pool Básico:
   func worker(id int, jobs <-chan int, results chan<- int) {
       for job := range jobs {
           // Processar job
           resultado := job * 2
           results <- resultado
       }
   }
   
   jobs := make(chan int, 100)
   results := make(chan int, 100)
   
   // Iniciar workers
   for w := 0; w < 3; w++ {
       go worker(w, jobs, results)
   }
   
   // Enviar jobs
   for j := 1; j <= 5; j++ {
       jobs <- j
   }
   close(jobs)
   
   // Coletar resultados
   for r := 1; r <= 5; r++ {
       fmt.Println(<-results)
   }

2. Worker Pool com WaitGroup:
   func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
       defer wg.Done()
       for job := range jobs {
           resultado := processar(job)
           results <- resultado
       }
   }
   
   var wg sync.WaitGroup
   numWorkers := 5
   jobs := make(chan int, 100)
   results := make(chan int, 100)
   
   for w := 0; w < numWorkers; w++ {
       wg.Add(1)
       go worker(w, jobs, results, &wg)
   }
   
   // Enviar jobs e fechar canal
   for j := 1; j <= 10; j++ {
       jobs <- j
   }
   close(jobs)
   
   // Aguardar workers terminarem
   go func() {
       wg.Wait()
       close(results)
   }()
   
   // Processar resultados
   for result := range results {
       fmt.Println(result)
   }

3. Worker Pool com Context:
   func worker(ctx context.Context, jobs <-chan string, results chan<- string) {
       for {
           select {
           case job, ok := <-jobs:
               if !ok {
                   return
               }
               resultado := processar(job)
               results <- resultado
           case <-ctx.Done():
               return
           }
       }
   }
   
   ctx, cancel := context.WithCancel(context.Background())
   defer cancel()
   
   jobs := make(chan string, 100)
   results := make(chan string, 100)
   
   for w := 0; w < 5; w++ {
       go worker(ctx, jobs, results)
   }

4. Worker Pool com Rate Limiting:
   rateLimiter := make(chan struct{}, 10)  // Máximo 10 por vez
   
   func worker(jobs <-chan int) {
       for job := range jobs {
           rateLimiter <- struct{}{}  // Adquire
           processar(job)
           <-rateLimiter  // Libera
       }
   }

DESAFIO:
Como desenvolvedor de sistema de processamento de imagens,
eu quero processar múltiplas imagens em paralelo com controle de recursos,
para que eu maximize throughput sem sobrecarregar o sistema.

Critérios de Aceitação:
- Criar worker pool com número limitado de workers
- Processar lista de tarefas
- Coletar resultados de forma ordenada
- Usar WaitGroup para sincronização

INSTRUÇÕES:
1. Crie uma função "processarTarefa" que simula processamento (time.Sleep)
2. Crie worker pool com 3 workers
3. Crie canal de jobs e canal de results
4. Envie 10 tarefas para processamento
5. Workers devem processar tarefas e enviar resultados
6. Colete e exiba todos os resultados
*/

import (
	"fmt"
	"sync"
	"time"
)

// TODO: Implemente worker pool aqui

func main() {
	// TODO: Implemente a solução do desafio aqui
	
	// Dica: Use make(chan tipo, buffer) para canais
	// Dica: Use sync.WaitGroup para sincronização
	// Dica: Feche canal de jobs após enviar todas as tarefas
	// Dica: Use range para iterar sobre canal
	
	// Exemplo de estrutura:
	// func processarTarefa(id int) int {
	//     time.Sleep(100 * time.Millisecond)  // Simula processamento
	//     return id * 2
	// }
	// 
	// func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	//     defer wg.Done()
	//     for job := range jobs {
	//         fmt.Printf("Worker %d processando tarefa %d\n", id, job)
	//         resultado := processarTarefa(job)
	//         results <- resultado
	//     }
	// }
	// 
	// numWorkers := 3
	// numJobs := 10
	// 
	// jobs := make(chan int, numJobs)
	// results := make(chan int, numJobs)
	// 
	// var wg sync.WaitGroup
	// 
	// // Iniciar workers
	// for w := 1; w <= numWorkers; w++ {
	//     wg.Add(1)
	//     go worker(w, jobs, results, &wg)
	// }
	// 
	// // Enviar jobs
	// for j := 1; j <= numJobs; j++ {
	//     jobs <- j
	// }
	// close(jobs)
	// 
	// // Aguardar workers e fechar results
	// go func() {
	//     wg.Wait()
	//     close(results)
	// }()
	// 
	// // Coletar resultados
	// for result := range results {
	//     fmt.Printf("Resultado: %d\n", result)
	// }
}



