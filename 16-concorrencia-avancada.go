//go:build exercicio16 || ignore
// +build exercicio16 ignore

package main

/*
EXERCÍCIO 16: Concorrência Avançada (Mutex, WaitGroup, Context)
Link do exercício: roadmap-com-desafios-golang.md - Etapa 3: Avançado - Concorrência Avançada

OBJETIVO:
Aprender a sincronizar goroutines e gerenciar concorrência com Mutex, WaitGroup e Context.

CONCEITOS:
MUTEX (Mutual Exclusion):
- sync.Mutex: Protege dados compartilhados
- Lock(): Bloqueia acesso
- Unlock(): Libera acesso
- Evita race conditions

WAITGROUP:
- sync.WaitGroup: Aguarda goroutines terminarem
- Add(n): Adiciona n goroutines
- Done(): Marca uma goroutine como concluída
- Wait(): Bloqueia até todas terminarem

CONTEXT:
- context.Context: Propaga cancelamento e timeout
- context.Background(): Context raiz
- context.WithTimeout(): Context com timeout
- context.WithCancel(): Context cancelável
- ctx.Done(): Channel que fecha quando cancelado

EXEMPLOS DE APLICAÇÃO:

1. Mutex Básico:
   var mu sync.Mutex
   var contador int
   
   func incrementar() {
       mu.Lock()
       contador++
       mu.Unlock()
   }

2. WaitGroup:
   var wg sync.WaitGroup
   
   for i := 0; i < 5; i++ {
       wg.Add(1)
       go func(id int) {
           defer wg.Done()
           fmt.Println("Goroutine", id)
       }(i)
   }
   
   wg.Wait()
   fmt.Println("Todas terminaram")

3. Context com Timeout:
   ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
   defer cancel()
   
   select {
   case <-time.After(5 * time.Second):
       fmt.Println("Operação concluída")
   case <-ctx.Done():
       fmt.Println("Timeout!")
   }

4. Context com Cancelamento:
   ctx, cancel := context.WithCancel(context.Background())
   
   go func() {
       time.Sleep(2 * time.Second)
       cancel()  // Cancela contexto
   }()
   
   <-ctx.Done()
   fmt.Println("Cancelado")

5. Mutex + WaitGroup:
   var mu sync.Mutex
   var wg sync.WaitGroup
   var soma int
   
   for i := 0; i < 10; i++ {
       wg.Add(1)
       go func(valor int) {
           defer wg.Done()
           mu.Lock()
           soma += valor
           mu.Unlock()
       }(i)
   }
   
   wg.Wait()
   fmt.Println("Soma:", soma)

6. RWMutex (Read-Write Mutex):
   var rwmu sync.RWMutex
   var dados map[string]int
   
   // Múltiplos leitores simultâneos:
   func ler(chave string) int {
       rwmu.RLock()
       defer rwmu.RUnlock()
       return dados[chave]
   }
   
   // Apenas um escritor:
   func escrever(chave string, valor int) {
       rwmu.Lock()
       defer rwmu.Unlock()
       dados[chave] = valor
   }

DESAFIO (User Story):
Como administrador de dashboard em tempo real,
eu quero sincronizar acessos com mutex e gerenciar goroutines com waitgroup e context,
para que evite race conditions e cancele tarefas longas.

Critérios de Aceitação:
- Usar Mutex para dado compartilhado
- WaitGroup para aguardar
- Context para timeout

INSTRUÇÕES:
1. Crie um contador compartilhado protegido por Mutex
2. Crie múltiplas goroutines que incrementam o contador
3. Use WaitGroup para aguardar todas as goroutines
4. Use Context com timeout para cancelar goroutines que demoram muito
5. Exiba o valor final do contador
*/

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// TODO: Implemente a solução do desafio aqui
	
	// Dica: Use sync.Mutex para proteger contador
	// Dica: Use sync.WaitGroup para aguardar goroutines
	// Dica: Use context.WithTimeout para limitar tempo
	// Dica: Use defer para garantir Unlock e Done
	
	// Exemplo de estrutura:
	// var mu sync.Mutex
	// var wg sync.WaitGroup
	// contador := 0
	// 
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()
	// 
	// for i := 0; i < 10; i++ {
	//     wg.Add(1)
	//     go func(id int) {
	//         defer wg.Done()
	//         
	//         select {
	//         case <-ctx.Done():
	//             fmt.Printf("Goroutine %d cancelada\n", id)
	//             return
	//         case <-time.After(1 * time.Second):
	//             mu.Lock()
	//             contador++
	//             mu.Unlock()
	//             fmt.Printf("Goroutine %d incrementou\n", id)
	//         }
	//     }(i)
	// }
	// 
	// wg.Wait()
	// fmt.Printf("Contador final: %d\n", contador)
}



