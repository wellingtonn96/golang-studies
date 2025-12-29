//go:build exercicio04 || ignore
// +build exercicio04 ignore

package main

/*
EXERCÍCIO 04: Operadores
Link do exercício: roadmap-com-desafios-golang.md - Etapa 1: Básico - Operadores

OBJETIVO:
Dominar os operadores aritméticos, de comparação e lógicos do Go.

CONCEITOS:
OPERADORES ARITMÉTICOS:
- + : Adição
- - : Subtração
- * : Multiplicação
- / : Divisão
- % : Módulo (resto da divisão)
- ++ : Incremento
- -- : Decremento

OPERADORES DE COMPARAÇÃO:
- == : Igual a
- != : Diferente de
- < : Menor que
- <= : Menor ou igual
- > : Maior que
- >= : Maior ou igual

OPERADORES LÓGICOS:
- && : E (AND)
- || : OU (OR)
- ! : NÃO (NOT)

OPERADORES DE ATRIBUIÇÃO:
- = : Atribuição simples
- += : Adição e atribuição
- -= : Subtração e atribuição
- *= : Multiplicação e atribuição
- /= : Divisão e atribuição
- %= : Módulo e atribuição

EXEMPLOS DE APLICAÇÃO:

1. Operadores Aritméticos:
   soma := 5 + 3        // 8
   subtracao := 10 - 4  // 6
   multiplicacao := 3 * 4  // 12
   divisao := 15 / 3    // 5
   resto := 10 % 3      // 1
   x := 5
   x++                  // x agora é 6

2. Operadores de Comparação:
   igual := 5 == 5      // true
   diferente := 5 != 3  // true
   menor := 3 < 5       // true
   maior := 10 > 5      // true

3. Operadores Lógicos:
   resultado := true && false  // false
   resultado2 := true || false // true
   resultado3 := !true         // false

   // Combinações:
   condicao := (idade >= 18) && (temCNH == true)
   desconto := (valor > 100) || (clienteVIP == true)

4. Operadores de Atribuição:
   x := 10
   x += 5   // x = 15
   x -= 3   // x = 12
   x *= 2   // x = 24
   x /= 4   // x = 6

DESAFIO (User Story):
Como dono de uma cafeteria,
eu quero calcular descontos automaticamente com base no valor da compra e no status VIP do cliente,
para que eu ofereça promoções personalizadas e aumente a fidelidade.

Critérios de Aceitação:
- Aplicar 10% de desconto se compra > 50
- Aplicar desconto extra de 5% se VIP
- Usar operadores para lógica e cálculos

INSTRUÇÕES:
1. Defina o valor da compra e se o cliente é VIP
2. Calcule o desconto base (10% se valor > 50)
3. Adicione desconto extra (5%) se for VIP
4. Calcule o valor final e exiba os detalhes
*/

func main() {
	// TODO: Implemente a solução do desafio aqui

	// Dica: Use operadores de comparação para verificar condições
	// Dica: Use operadores aritméticos para cálculos
	// Dica: Use operadores lógicos para combinar condições

	// Exemplo de estrutura:
	// valorCompra := 75.0
	// clienteVIP := true
	// desconto := 0.0
	//
	// if valorCompra > 50 {
	//     desconto += valorCompra * 0.10
	// }
	// if clienteVIP {
	//     desconto += valorCompra * 0.05
	// }
	//
	// valorFinal := valorCompra - desconto
	// fmt.Printf("Valor original: R$ %.2f\nDesconto: R$ %.2f\nValor final: R$ %.2f\n", ...)
}
