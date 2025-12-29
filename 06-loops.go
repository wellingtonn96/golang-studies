//go:build exercicio06 || ignore
// +build exercicio06 ignore

package main

/*
EXERCÍCIO 06: Loops (for, range)
Link do exercício: roadmap-com-desafios-golang.md - Etapa 1: Básico - Loops

OBJETIVO:
Aprender a criar loops e iterar sobre coleções em Go.

CONCEITOS:
FOR CLÁSSICO:
- for init; condição; incremento { }: Loop tradicional
- for condição { }: Loop while-like
- for { }: Loop infinito (use break para sair)

RANGE:
- for idx, val := range slice: Itera sobre slice com índice e valor
- for idx := range slice: Apenas índices
- for _, val := range slice: Apenas valores (ignora índice)
- for key, val := range map: Itera sobre map
- for key := range map: Apenas chaves
- for idx, char := range string: Itera sobre string (retorna rune)

CONTROLE DE LOOP:
- break: Sai do loop imediatamente
- continue: Pula para próxima iteração
- goto: Vai para label (evite usar)

EXEMPLOS DE APLICAÇÃO:

1. For clássico:
   for i := 0; i < 5; i++ {
       fmt.Println(i)
   }

2. For while-like:
   i := 0
   for i < 5 {
       fmt.Println(i)
       i++
   }

3. For infinito:
   for {
       // código
       if condicao {
           break
       }
   }

4. Range em slice:
   numeros := []int{10, 20, 30}
   for idx, valor := range numeros {
       fmt.Printf("Índice %d: %d\n", idx, valor)
   }
   
   // Apenas valores:
   for _, valor := range numeros {
       fmt.Println(valor)
   }

5. Range em map:
   idades := map[string]int{
       "João": 25,
       "Maria": 30,
   }
   for nome, idade := range idades {
       fmt.Printf("%s tem %d anos\n", nome, idade)
   }

6. Range em string:
   texto := "Olá"
   for idx, char := range texto {
       fmt.Printf("Posição %d: %c\n", idx, char)
   }

7. Break e Continue:
   for i := 0; i < 10; i++ {
       if i == 3 {
           continue  // Pula o 3
       }
       if i == 8 {
           break  // Para no 8
       }
       fmt.Println(i)
   }

DESAFIO (User Story):
Como vendedor,
eu quero calcular o total de vendas da semana e simular a redução de estoque,
para que eu tenha relatórios rápidos e saiba quando repor produtos.

Critérios de Aceitação:
- Usar range para somar vendas em slice
- Usar for para simular redução de estoque até zero
- Exibir total e "estoque esgotado"

INSTRUÇÕES:
1. Crie um slice com as vendas da semana (valores em reais)
2. Use range para somar todas as vendas
3. Simule um estoque inicial (ex: 100 unidades)
4. Use for para reduzir o estoque até zero, exibindo a cada iteração
5. Exiba o total de vendas e mensagem de estoque esgotado

OPCIONAL (Recomendado para prática avançada):
- Considere criar uma struct "RelatorioVendas" para agrupar os dados
- Exemplo: type RelatorioVendas struct { Vendas []float64; Total float64; Estoque int }
- Isso facilita passar os dados entre funções e organizar melhor o código
*/

import "fmt"

func main() {
	// TODO: Implemente a solução do desafio aqui
	
	// Dica: Use range para iterar sobre vendas e somar
	// Dica: Use for para reduzir estoque até zero
	// Dica: Use fmt.Printf para formatar saída
	
	// Exemplo de estrutura:
	// vendas := []float64{150.0, 200.0, 180.0, 250.0, 300.0}
	// 
	// totalVendas := 0.0
	// for _, venda := range vendas {
	//     totalVendas += venda
	// }
	// 
	// estoque := 100
	// for estoque > 0 {
	//     fmt.Printf("Estoque atual: %d unidades\n", estoque)
	//     estoque -= 10  // Reduz 10 unidades por iteração
	// }
	// 
	// fmt.Printf("\nTotal de vendas da semana: R$ %.2f\n", totalVendas)
	// fmt.Println("Estoque esgotado!")
}



