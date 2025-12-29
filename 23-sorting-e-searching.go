//go:build exercicio23 || ignore
// +build exercicio23 ignore

package main

/*
EXERCÍCIO 23: Sorting e Searching
Link do exercício: Conceito essencial para desenvolvimento real

OBJETIVO:
Aprender a ordenar e buscar em slices usando a biblioteca padrão.

CONCEITOS:
SORT PACKAGE:
- sort.Ints(): Ordena slice de ints
- sort.Float64s(): Ordena slice de float64s
- sort.Strings(): Ordena slice de strings
- sort.Slice(): Ordena slice genérico com função de comparação
- sort.Sort(): Ordena implementando sort.Interface

SEARCHING:
- sort.SearchInts(): Busca binária em ints ordenados
- sort.Search(): Busca binária genérica
- sort.SearchStrings(): Busca binária em strings

CUSTOM SORTING:
- Implementar sort.Interface (Len, Less, Swap)
- Usar sort.Slice com função de comparação

EXEMPLOS DE APLICAÇÃO:

1. Ordenar Ints:
   numeros := []int{3, 1, 4, 1, 5, 9, 2, 6}
   sort.Ints(numeros)
   fmt.Println(numeros)  // [1 1 2 3 4 5 6 9]

2. Ordenar Strings:
   nomes := []string{"Zé", "Ana", "Carlos", "Bruno"}
   sort.Strings(nomes)
   fmt.Println(nomes)  // [Ana Bruno Carlos Zé]

3. Ordenar com Função Customizada:
   pessoas := []struct {
       Nome  string
       Idade int
   }{
       {"João", 30},
       {"Maria", 25},
       {"Pedro", 35},
   }
   
   sort.Slice(pessoas, func(i, j int) bool {
       return pessoas[i].Idade < pessoas[j].Idade
   })
   // Ordenado por idade

4. Ordenar Decrescente:
   numeros := []int{3, 1, 4, 1, 5, 9, 2, 6}
   sort.Slice(numeros, func(i, j int) bool {
       return numeros[i] > numeros[j]
   })

5. Busca Binária:
   numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
   pos := sort.SearchInts(numeros, 5)
   fmt.Println(pos)  // 4 (índice)
   
   // Verificar se existe:
   if pos < len(numeros) && numeros[pos] == 5 {
       fmt.Println("Encontrado!")
   }

6. Busca Binária Genérica:
   numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
   pos := sort.Search(len(numeros), func(i int) bool {
       return numeros[i] >= 5
   })
   fmt.Println(pos)

7. Implementar sort.Interface:
   type Pessoa struct {
       Nome  string
       Idade int
   }
   
   type PorIdade []Pessoa
   
   func (p PorIdade) Len() int           { return len(p) }
   func (p PorIdade) Less(i, j int) bool { return p[i].Idade < p[j].Idade }
   func (p PorIdade) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
   
   pessoas := PorIdade{{"João", 30}, {"Maria", 25}}
   sort.Sort(pessoas)

8. Ordenar Múltiplos Campos:
   sort.Slice(pessoas, func(i, j int) bool {
       if pessoas[i].Idade != pessoas[j].Idade {
           return pessoas[i].Idade < pessoas[j].Idade
       }
       return pessoas[i].Nome < pessoas[j].Nome
   })

DESAFIO:
Como desenvolvedor de sistema de ranking,
eu quero ordenar e buscar em listas de jogadores,
para que eu exiba rankings e encontre jogadores rapidamente.

Critérios de Aceitação:
- Ordenar jogadores por pontuação (decrescente)
- Buscar jogador por nome
- Ordenar por múltiplos critérios (pontuação, depois nome)
- Exibir top 10 jogadores

INSTRUÇÕES:
1. Crie struct Jogador com Nome e Pontuacao
2. Crie slice de jogadores
3. Ordene por pontuação (maior para menor)
4. Se pontuação igual, ordene por nome
5. Implemente busca por nome
6. Exiba top 10 jogadores
*/

import (
	"fmt"
	"sort"
	"strings"
)

// TODO: Defina struct Jogador
// type Jogador struct { ... }

func main() {
	// TODO: Implemente a solução do desafio aqui
	
	// Dica: Use sort.Slice com função de comparação
	// Dica: Para busca, use loop ou sort.Search
	// Dica: Para top N, use slice[:N]
	
	// Exemplo de estrutura:
	// jogadores := []Jogador{
	//     {"João", 1500},
	//     {"Maria", 2000},
	//     {"Pedro", 1500},
	//     {"Ana", 1800},
	// }
	// 
	// sort.Slice(jogadores, func(i, j int) bool {
	//     if jogadores[i].Pontuacao != jogadores[j].Pontuacao {
	//         return jogadores[i].Pontuacao > jogadores[j].Pontuacao
	//     }
	//     return jogadores[i].Nome < jogadores[j].Nome
	// })
	// 
	// top10 := jogadores
	// if len(top10) > 10 {
	//     top10 = top10[:10]
	// }
	// 
	// for i, j := range top10 {
	//     fmt.Printf("%d. %s - %d pontos\n", i+1, j.Nome, j.Pontuacao)
	// }
	// 
	// // Buscar por nome:
	// nomeBuscado := "Maria"
	// for _, j := range jogadores {
	//     if strings.EqualFold(j.Nome, nomeBuscado) {
	//         fmt.Printf("Encontrado: %s com %d pontos\n", j.Nome, j.Pontuacao)
	//         break
	//     }
	// }
}



