//go:build exercicio00 || ignore
// +build exercicio00 ignore

package main

/*
EXERCÍCIO 08: Manipulação de Slices, Maps e Strings
Link do exercício: roadmap-com-desafios-golang.md - Etapa 1: Básico - Manipulação de Slices, Maps e Strings

OBJETIVO:
Aprender a manipular slices, maps e strings usando a biblioteca padrão do Go.

CONCEITOS:
SLICES:
- append(slice, elementos...): Adiciona elementos ao slice
- len(slice): Retorna o tamanho
- cap(slice): Retorna a capacidade
- slice[inicio:fim]: Cria sub-slice
- make([]tipo, tamanho, capacidade): Cria slice com tamanho/capacidade

MAPS:
- map[chave]valor: Estrutura chave-valor
- map[chave] = valor: Adiciona/atualiza
- delete(map, chave): Remove elemento
- valor, existe := map[chave]: Verifica existência
- len(map): Retorna número de elementos

STRINGS:
- strings package: Funções úteis para strings
- len(string): Tamanho em bytes
- strings.Contains(s, substr): Verifica se contém
- strings.ToUpper(s) / ToLower(s): Maiúsculas/minúsculas
- strings.TrimSpace(s): Remove espaços
- strings.Split(s, sep): Divide string
- strings.Join(slice, sep): Junta slice em string

EXEMPLOS DE APLICAÇÃO:

1. Manipulação de Slices:
   numeros := []int{1, 2, 3}
   numeros = append(numeros, 4)        // [1 2 3 4]
   numeros = append(numeros, 5, 6)     // [1 2 3 4 5 6]

   tamanho := len(numeros)              // 6
   capacidade := cap(numeros)           // pode ser maior que tamanho

   sub := numeros[1:4]                  // [2 3 4]
   primeiro := numeros[0]               // 1
   ultimo := numeros[len(numeros)-1]   // 6

2. Manipulação de Maps:
   idades := make(map[string]int)
   idades["João"] = 25                 // Adiciona
   idades["Maria"] = 30
   idades["João"] = 26                 // Atualiza

   idade, existe := idades["Pedro"]     // 0, false
   if existe {
       fmt.Println(idade)
   }

   delete(idades, "Maria")             // Remove
   total := len(idades)                 // 1

3. Manipulação de Strings:
   import "strings"

   texto := "  Olá Mundo  "
   maiuscula := strings.ToUpper(texto)        // "  OLÁ MUNDO  "
   minuscula := strings.ToLower(texto)        // "  olá mundo  "
   semEspacos := strings.TrimSpace(texto)     // "Olá Mundo"

   contem := strings.Contains(texto, "Mundo") // true

   palavras := strings.Split("a,b,c", ",")    // ["a", "b", "c"]
   juntado := strings.Join(palavras, "-")     // "a-b-c"

   tamanho := len(texto)                      // 13 (bytes)

4. Validação de Email (exemplo prático):
   email := "usuario@email.com"
   if strings.Contains(email, "@") && strings.Contains(email, ".") {
       fmt.Println("Email válido")
   }

DESAFIO (User Story):
Como profissional de marketing,
eu quero processar uma lista de emails de clientes para extrair nomes e contar válidos,
para que eu possa preparar uma campanha segmentada.

Critérios de Aceitação:
- Usar slice para emails
- Usar map para contagens
- Manipular strings para validação (contains "@")

INSTRUÇÕES:
1. Crie um slice com uma lista de emails (alguns válidos, alguns inválidos)
2. Use um map para contar emails válidos e inválidos
3. Valide cada email verificando se contém "@"
4. Extraia o nome (parte antes do @) de cada email válido
5. Exiba estatísticas e lista de nomes extraídos
*/

func main() {
	// TODO: Implemente a solução do desafio aqui
	
	// Dica: Use slice para armazenar emails
	// Dica: Use map para contar válidos/inválidos
	// Dica: Use strings.Contains para validar
	// Dica: Use strings.Split para extrair nome
	
	// Exemplo de estrutura:
	// emails := []string{
	//     "joao@email.com",
	//     "maria@email.com",
	//     "email-invalido",
	//     "pedro@email.com",
	// }
	// 
	// contagem := map[string]int{
	//     "validos":   0,
	//     "invalidos": 0,
	// }
	// 
	// nomes := []string{}
	// 
	// for _, email := range emails {
	//     if strings.Contains(email, "@") {
	//         contagem["validos"]++
	//         partes := strings.Split(email, "@")
	//         nomes = append(nomes, partes[0])
	//     } else {
	//         contagem["invalidos"]++
	//     }
	// }
	// 
	// fmt.Printf("Emails válidos: %d\n", contagem["validos"])
	// fmt.Printf("Emails inválidos: %d\n", contagem["invalidos"])
	// fmt.Printf("Nomes extraídos: %v\n", nomes)
}



