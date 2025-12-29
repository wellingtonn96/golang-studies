//go:build exercicio11 || ignore
// +build exercicio11 ignore

package main

/*
EXERCÍCIO 11: Métodos e Receivers (Value vs Pointer)
Link do exercício: roadmap-com-desafios-golang.md - Etapa 2: Intermediário - Métodos e Receivers

OBJETIVO:
Aprender a adicionar métodos a tipos usando receivers (value e pointer).

CONCEITOS:
MÉTODOS:
- func (receiver Tipo) NomeMetodo() retorno { }
- Permite adicionar comportamento a tipos
- Receiver pode ser value ou pointer

VALUE RECEIVER:
- func (t Tipo) Metodo(): Recebe cópia do valor
- Não modifica o original
- Use quando método apenas lê dados

POINTER RECEIVER:
- func (t *Tipo) Metodo(): Recebe ponteiro
- Pode modificar o original
- Use quando método modifica estado
- Mais eficiente para structs grandes

ESCOLHA ENTRE VALUE E POINTER:
- Value: Métodos de leitura, tipos pequenos, imutabilidade
- Pointer: Métodos de mutação, structs grandes, consistência (use sempre pointer ou sempre value)

EXEMPLOS DE APLICAÇÃO:

1. Value Receiver:
   type Circulo struct {
       Raio float64
   }
   
   func (c Circulo) Area() float64 {
       return 3.14159 * c.Raio * c.Raio
   }
   
   circ := Circulo{Raio: 5}
   area := circ.Area()  // 78.54
   // circ não foi modificado

2. Pointer Receiver:
   type Contador struct {
       Valor int
   }
   
   func (c *Contador) Incrementar() {
       c.Valor++
   }
   
   cont := Contador{Valor: 0}
   cont.Incrementar()
   fmt.Println(cont.Valor)  // 1 (modificado)

3. Método que Modifica (deve usar pointer):
   type Retangulo struct {
       Largura float64
       Altura  float64
   }
   
   func (r *Retangulo) Escalar(fator float64) {
       r.Largura *= fator
       r.Altura *= fator
   }
   
   ret := Retangulo{Largura: 10, Altura: 5}
   ret.Escalar(2)
   fmt.Println(ret.Largura, ret.Altura)  // 20 10

4. Múltiplos Métodos:
   type ContaBancaria struct {
       Saldo float64
   }
   
   func (c *ContaBancaria) Depositar(valor float64) {
       c.Saldo += valor
   }
   
   func (c *ContaBancaria) Sacar(valor float64) bool {
       if valor <= c.Saldo {
           c.Saldo -= valor
           return true
       }
       return false
   }
   
   func (c ContaBancaria) ConsultarSaldo() float64 {
       return c.Saldo
   }

5. Método em Tipo Não-Struct:
   type MeuInt int
   
   func (m MeuInt) Dobrar() MeuInt {
       return m * 2
   }
   
   numero := MeuInt(5)
   resultado := numero.Dobrar()  // 10

DESAFIO (User Story):
Como desenvolvedor de CRM,
eu quero adicionar métodos a structs de clientes para cálculos personalizados,
para que eu modifique estados eficientemente com receivers.

Critérios de Aceitação:
- Método com value receiver para leitura
- Método com pointer receiver para mutação
- Chamar e verificar mudanças

INSTRUÇÕES:
1. Crie uma struct "Cliente" com campos: Nome, Saldo, Ativo
2. Crie um método com value receiver "ExibirInfo()" que retorna string formatada
3. Crie um método com pointer receiver "AdicionarSaldo(valor)" que adiciona ao saldo
4. Crie um método com pointer receiver "Desativar()" que muda Ativo para false
5. Crie uma instância, chame os métodos e exiba os resultados
*/

import "fmt"

// TODO: Defina a struct Cliente e seus métodos aqui
// type Cliente struct { ... }
// func (c Cliente) ExibirInfo() string { ... }
// func (c *Cliente) AdicionarSaldo(valor float64) { ... }
// func (c *Cliente) Desativar() { ... }

func main() {
	// TODO: Implemente a solução do desafio aqui
	
	// Dica: Use value receiver para métodos de leitura
	// Dica: Use pointer receiver para métodos que modificam estado
	// Dica: Go permite chamar métodos mesmo quando você tem ponteiro ou valor
	
	// Exemplo de estrutura:
	// cliente := Cliente{
	//     Nome:  "João Silva",
	//     Saldo: 1000.0,
	//     Ativo: true,
	// }
	// 
	// fmt.Println(cliente.ExibirInfo())
	// 
	// cliente.AdicionarSaldo(500.0)
	// fmt.Println("Após adicionar saldo:")
	// fmt.Println(cliente.ExibirInfo())
	// 
	// cliente.Desativar()
	// fmt.Println("Após desativar:")
	// fmt.Println(cliente.ExibirInfo())
}

