//go:build exercicio03 || ignore
// +build exercicio03 ignore

package main

import "fmt"

/*
EXERCÍCIO 03: Tipos de Dados Primitivos e Compostos
Link do exercício: roadmap-com-desafios-golang.md - Etapa 1: Básico - Tipos de Dados Primitivos e Compostos

OBJETIVO:
Conhecer e utilizar os tipos de dados primitivos e compostos do Go.

CONCEITOS:
TIPOS PRIMITIVOS:
- int, int8, int16, int32, int64: Números inteiros
- uint, uint8, uint16, uint32, uint64: Números inteiros sem sinal
- float32, float64: Números decimais
- string: Cadeia de caracteres
- bool: Valores booleanos (true/false)
- byte: Alias para uint8
- rune: Alias para int32 (representa um código Unicode)

TIPOS COMPOSTOS:
- Array: [tamanho]tipo - Tamanho fixo
- Slice: []tipo - Tamanho dinâmico (mais usado)
- Map: map[chave]valor - Estrutura chave-valor
- Struct: Estrutura personalizada (coleção de campos nomeados)

DEFINIÇÃO DE STRUCT:
- type NomeStruct struct { campo1 tipo1; campo2 tipo2 }
- Permite agrupar dados relacionados
- Campos podem ser de tipos diferentes
- Inicialização: NomeStruct{campo1: valor1, campo2: valor2}
- Acesso: instancia.campo
- Structs aninhadas: campo pode ser outra struct

EXEMPLOS DE APLICAÇÃO:

 1. Tipos Primitivos:
    var idade int = 30
    var salario float64 = 5000.50
    var nome string = "Maria"
    var ativo bool = true

 2. Array (tamanho fixo):
    var numeros [3]int = [3]int{1, 2, 3}
    var nomes [2]string = [2]string{"João", "Maria"}

 3. Slice (tamanho dinâmico - mais comum):
    numeros := []int{1, 2, 3}
    nomes := []string{"João", "Maria", "Pedro"}
    var vazio []int  // Slice vazio

 4. Map (chave-valor):
    idades := map[string]int{
    "João": 25,
    "Maria": 30,
    }
    detalhes := map[string]interface{}{
    "nome": "Produto",
    "preco": 99.90,
    }

 5. Acessar elementos:
    primeiro := numeros[0]        // Slice/Array
    idadeJoao := idades["João"]   // Map
    idades["Pedro"] = 28          // Adicionar ao Map

 6. Struct (estrutura personalizada):
    type Pessoa struct {
        Nome  string
        Idade int
        Ativo bool
    }

    // Inicialização com nomes de campos:
    pessoa := Pessoa{
        Nome:  "João",
        Idade: 30,
        Ativo: true,
    }

    // Acesso aos campos:
    fmt.Println(pessoa.Nome)   // João
    pessoa.Idade = 31          // Modificar campo

    // Struct aninhada:
    type Endereco struct {
        Rua    string
        Cidade string
    }

    type Cliente struct {
        Nome     string
        Endereco Endereco  // Struct aninhada
    }

    cliente := Cliente{
        Nome: "Maria",
        Endereco: Endereco{
            Rua:    "Rua A, 123",
            Cidade: "São Paulo",
        },
    }
    fmt.Println(cliente.Endereco.Rua)  // Rua A, 123

DESAFIO (User Story):
Como analista de vendas,
eu quero armazenar dados de um cliente usando os tipos corretos de Golang,
para que eu possa validar se o cliente está ativo e decidir se envio uma campanha promocional.

Critérios de Aceitação:
- Armazenar ID (int), nome (string), ativo (bool), endereços (slice) e detalhes (map)
- Verificar booleano ativo
- Exibir mensagem "Cliente ativo: sim/não"

INSTRUÇÕES:
1. Crie variáveis para: ID (int), nome (string), ativo (bool)
2. Crie um slice de strings para endereços
3. Crie um map[string]string para detalhes adicionais
4. Verifique se o cliente está ativo e exiba a mensagem apropriada

OPCIONAL (Recomendado para prática avançada):
- Considere criar uma struct "Cliente" para agrupar todos esses dados
- Isso torna o código mais organizado e fácil de manter
- Exemplo: type Cliente struct { ID int; Nome string; Ativo bool; ... }
*/
func main() {
	// Estruturas
	type Address struct {
		Street  string
		Zipcode string
	}

    type Details struct {
        telefone string
        email    string
    }

	type Cliente struct {
		ID        int
		Nome      string
		Active    bool
		Address   Address
		Enderecos []string
		Details Details
	}



	// Instância do cliente
	cliente := Cliente{
		ID:     1,
		Nome:   "Wellington",
		Active: true,
		Address: Address{
			Street:  "Rua Guilhermino",
			Zipcode: "08390541",
		},
		Details: Details{
            telefone: "123456789",
            email:    "cliente@email.com",
		},
	}

	// Verifica se o cliente está ativo
	if cliente.Active == true {
		fmt.Println("Cliente ativo: sim")
	} else {
		fmt.Println("Cliente ativo: não")
	}

	// Prints adicionais (exercício completo)
	fmt.Println("ID:", cliente.ID)
	fmt.Println("Nome:", cliente.Nome)
	fmt.Println("Endereço principal:", cliente.Address.Street, "-", cliente.Address.Zipcode)
	fmt.Println("Endereços:", cliente.Enderecos)
	fmt.Println("Detalhes:", cliente.Details.telefone, "-", cliente.Details.email)
}
