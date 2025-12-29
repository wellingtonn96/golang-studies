// go:build exercicio02 || ignore
// +build exercicio02 ignore

//go:build ignore
// +build ignore

package main

import "fmt"

/*
EXERCÍCIO 02: Variáveis e Constantes
Link do exercício: roadmap-com-desafios-golang.md - Etapa 1: Básico - Variáveis e Constantes

OBJETIVO:
Aprender a declarar variáveis e constantes em Go.

CONCEITOS:
- var: Declaração explícita de variável com tipo
- := : Declaração curta com inferência de tipo (só dentro de funções)
- const: Declaração de constante (valor imutável)
- Tipos básicos: int, string, bool, float64

EXEMPLOS DE APLICAÇÃO:

1. Declaração longa (com tipo explícito):
   var idade int = 25
   var nome string = "João"

2. Declaração curta (inferência de tipo):
   idade := 25        // Go infere que é int
   nome := "João"     // Go infere que é string
   ativo := true      // Go infere que é bool

3. Declaração múltipla:
   var a, b int = 10, 20
   x, y := 5, 10

4. Constantes:
   const PI = 3.14159
   const NOME_PRODUTO = "Notebook"

5. Constantes tipadas:
   const MAX_USUARIOS int = 100

DESAFIO (User Story):
Como gerente de estoque em uma loja online,
eu quero rastrear o preço de um produto que pode mudar durante promoções,
mas o nome do produto deve ser constante,
para que eu consiga calcular o total de vendas sem risco de alterar acidentalmente o nome.

Critérios de Aceitação:
- Usar declaração curta (:=) para preço mutável
- Usar const para nome imutável
- Calcular e exibir total para 5 itens (preço * 5)

INSTRUÇÕES:
1. Declare uma constante para o nome do produto
2. Declare uma variável para o preço (usando :=)
3. Calcule o total para 5 unidades
4. Exiba o nome do produto, preço unitário e total
*/

func main() {
	// Implementação do desafio
	const nomeProduto = "Meu produto"
	preco := 29.90
	total := preco * 5
	fmt.Println(nomeProduto, preco, total)
}
