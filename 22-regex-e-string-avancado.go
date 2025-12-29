//go:build exercicio22 || ignore
// +build exercicio22 ignore

package main

/*
EXERCÍCIO 22: Regex e String Avançado
Link do exercício: Conceito essencial para desenvolvimento real

OBJETIVO:
Aprender a usar expressões regulares e manipulação avançada de strings.

CONCEITOS:
REGEXP PACKAGE:
- regexp.MustCompile(): Compila regex (panic se inválida)
- regexp.Compile(): Compila regex (retorna error)
- MatchString(): Verifica se string corresponde
- FindString(): Encontra primeira correspondência
- FindAllString(): Encontra todas correspondências
- ReplaceAllString(): Substitui correspondências

STRINGS PACKAGE AVANÇADO:
- strings.Builder: Construção eficiente de strings
- strings.HasPrefix/HasSuffix: Verifica início/fim
- strings.Index/LastIndex: Encontra posição
- strings.Replace: Substitui ocorrências
- strings.Fields: Divide por espaços
- strings.Join: Junta strings

EXEMPLOS DE APLICAÇÃO:

1. Regex Básica:
   import "regexp"
   
   padrao := regexp.MustCompile(`\d+`)
   match := padrao.MatchString("abc123")
   fmt.Println(match)  // true

2. Encontrar Correspondências:
   padrao := regexp.MustCompile(`\d+`)
   texto := "Tenho 5 maçãs e 10 laranjas"
   
   primeira := padrao.FindString(texto)      // "5"
   todas := padrao.FindAllString(texto, -1)  // ["5", "10"]

3. Grupos de Captura:
   padrao := regexp.MustCompile(`(\d{2})/(\d{2})/(\d{4})`)
   texto := "Data: 15/01/2024"
   
   matches := padrao.FindStringSubmatch(texto)
   // matches[0]: correspondência completa
   // matches[1]: dia
   // matches[2]: mês
   // matches[3]: ano

4. Substituir com Regex:
   padrao := regexp.MustCompile(`\d+`)
   texto := "Preço: R$ 100"
   novo := padrao.ReplaceAllString(texto, "XXX")
   fmt.Println(novo)  // "Preço: R$ XXX"

5. Validar Email:
   emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
   email := "usuario@email.com"
   valido := emailRegex.MatchString(email)
   fmt.Println(valido)

6. Validar CPF:
   cpfRegex := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$`)
   cpf := "123.456.789-00"
   valido := cpfRegex.MatchString(cpf)

7. Strings.Builder:
   var builder strings.Builder
   builder.WriteString("Olá")
   builder.WriteString(" ")
   builder.WriteString("Mundo")
   resultado := builder.String()  // "Olá Mundo"

8. Strings Avançado:
   texto := "  Olá Mundo  "
   
   // Verificar prefixo/sufixo:
   strings.HasPrefix(texto, "Olá")  // false (tem espaços)
   strings.HasSuffix(texto, "Mundo")  // false
   
   // Encontrar posição:
   pos := strings.Index(texto, "Mundo")  // 5
   
   // Dividir:
   palavras := strings.Fields(texto)  // ["Olá", "Mundo"]
   
   // Substituir:
   novo := strings.Replace(texto, "Mundo", "Go", -1)

9. Extrair Números de String:
   padrao := regexp.MustCompile(`\d+\.?\d*`)
   texto := "Preço: R$ 99.90, Desconto: 10%"
   numeros := padrao.FindAllString(texto, -1)
   // ["99.90", "10"]

DESAFIO:
Como desenvolvedor de sistema de validação,
eu quero validar e extrair informações de textos,
para que eu processe dados de formulários corretamente.

Critérios de Aceitação:
- Validar formato de email
- Extrair números de telefone
- Validar data no formato DD/MM/YYYY
- Substituir informações sensíveis

INSTRUÇÕES:
1. Crie uma função que valida email usando regex
2. Crie uma função que extrai números de telefone (formato: (XX) XXXXX-XXXX)
3. Crie uma função que valida data no formato brasileiro
4. Crie uma função que mascara emails (substitui parte do email por ***)
5. Teste todas as funções com diferentes entradas
*/

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// TODO: Implemente as funções de validação e extração
	
	// Dica: Use regexp.MustCompile para compilar padrões
	// Dica: Use MatchString para validar
	// Dica: Use FindString/FindAllString para extrair
	// Dica: Use ReplaceAllString para substituir
	
	// Exemplo de estrutura:
	// func validarEmail(email string) bool {
	//     padrao := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	//     return padrao.MatchString(email)
	// }
	// 
	// func extrairTelefone(texto string) []string {
	//     padrao := regexp.MustCompile(`\(\d{2}\)\s?\d{4,5}-?\d{4}`)
	//     return padrao.FindAllString(texto, -1)
	// }
	// 
	// func validarData(data string) bool {
	//     padrao := regexp.MustCompile(`^\d{2}/\d{2}/\d{4}$`)
	//     return padrao.MatchString(data)
	// }
	// 
	// func mascararEmail(email string) string {
	//     partes := strings.Split(email, "@")
	//     if len(partes) != 2 {
	//         return email
	//     }
	//     nome := partes[0]
	//     dominio := partes[1]
	//     if len(nome) > 2 {
	//         nome = nome[:2] + "***"
	//     }
	//     return nome + "@" + dominio
	// }
	// 
	// // Testes:
	// fmt.Println(validarEmail("user@email.com"))  // true
	// fmt.Println(validarEmail("invalido"))       // false
	// 
	// texto := "Contato: (11) 98765-4321 ou (21) 1234-5678"
	// telefones := extrairTelefone(texto)
	// fmt.Println(telefones)
	// 
	// fmt.Println(validarData("15/01/2024"))  // true
	// fmt.Println(validarData("2024-01-15"))   // false
	// 
	// fmt.Println(mascararEmail("usuario@email.com"))  // us***@email.com
}



