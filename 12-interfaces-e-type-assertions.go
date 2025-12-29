//go:build exercicio12 || ignore
// +build exercicio12 ignore

package main

/*
EXERCÍCIO 12: Interfaces e Type Assertions
Link do exercício: roadmap-com-desafios-golang.md - Etapa 2: Intermediário - Interfaces e Type Assertions

OBJETIVO:
Aprender a usar interfaces para polimorfismo e type assertions para verificação de tipos.

CONCEITOS:
INTERFACES:
- type Nome interface { métodos }: Define contrato de métodos
- Tipo implementa interface implicitamente (duck typing)
- Interface vazia: interface{} (aceita qualquer tipo)
- Uma interface pode ter múltiplos métodos

IMPLEMENTAÇÃO IMPLÍCITA:
- Não precisa declarar "implements"
- Se tipo tem os métodos, implementa a interface
- Permite polimorfismo

TYPE ASSERTION:
- valor.(Tipo): Converte interface para tipo específico
- valor, ok := interface.(Tipo): Verifica se conversão é possível
- Útil para acessar métodos/campos específicos

EXEMPLOS DE APLICAÇÃO:

1. Interface Básica:
   type Forma interface {
       Area() float64
   }
   
   type Quadrado struct {
       Lado float64
   }
   
   func (q Quadrado) Area() float64 {
       return q.Lado * q.Lado
   }
   
   var f Forma = Quadrado{Lado: 4}
   fmt.Println(f.Area())  // 16

2. Múltiplas Implementações:
   type Circulo struct {
       Raio float64
   }
   
   func (c Circulo) Area() float64 {
       return 3.14159 * c.Raio * c.Raio
   }
   
   formas := []Forma{
       Quadrado{Lado: 4},
       Circulo{Raio: 3},
   }
   
   for _, forma := range formas {
       fmt.Println(forma.Area())
   }

3. Type Assertion:
   var f Forma = Quadrado{Lado: 4}
   
   // Assertion simples (pode causar panic se falhar):
   quadrado := f.(Quadrado)
   fmt.Println(quadrado.Lado)  // 4
   
   // Assertion segura:
   quadrado, ok := f.(Quadrado)
   if ok {
       fmt.Println("É um quadrado:", quadrado.Lado)
   }

4. Type Switch:
   func descrever(f Forma) {
       switch v := f.(type) {
       case Quadrado:
           fmt.Printf("Quadrado com lado %.2f\n", v.Lado)
       case Circulo:
           fmt.Printf("Círculo com raio %.2f\n", v.Raio)
       default:
           fmt.Println("Forma desconhecida")
       }
   }

5. Interface com Múltiplos Métodos:
   type Animal interface {
       FazerSom() string
       Mover() string
   }
   
   type Cachorro struct {
       Nome string
   }
   
   func (c Cachorro) FazerSom() string {
       return "Au au!"
   }
   
   func (c Cachorro) Mover() string {
       return "Correndo"
   }

6. Interface Vazia (interface{}):
   var qualquer interface{}
   qualquer = 42
   qualquer = "texto"
   qualquer = []int{1, 2, 3}
   
   // Type assertion para usar:
   if texto, ok := qualquer.(string); ok {
       fmt.Println(texto)
   }

DESAFIO (User Story):
Como gestor de equipe,
eu quero usar interfaces para definir comportamentos comuns em structs de funcionários,
para que eu tenha flexibilidade em hierarquias.

Critérios de Aceitação:
- Definir interface com método
- Implementar em struct
- Usar type assertion para acessar campos específicos

INSTRUÇÕES:
1. Crie uma interface "Funcionario" com método "CalcularSalario() float64"
2. Crie structs "Desenvolvedor" e "Gerente" que implementam a interface
3. Cada struct deve ter campos específicos (ex: Desenvolvedor tem "Linguagem", Gerente tem "Equipe")
4. Crie uma função que aceita interface Funcionario e usa type assertion para exibir informações específicas
5. Teste com instâncias de ambos os tipos
*/

import "fmt"

// TODO: Defina a interface e structs aqui
// type Funcionario interface { ... }
// type Desenvolvedor struct { ... }
// type Gerente struct { ... }

func main() {
	// TODO: Implemente a solução do desafio aqui
	
	// Dica: Implemente CalcularSalario() em cada struct
	// Dica: Use type assertion para acessar campos específicos
	// Dica: Use type switch para diferentes tipos
	
	// Exemplo de estrutura:
	// func (d Desenvolvedor) CalcularSalario() float64 {
	//     return 5000.0  // Salário base
	// }
	// 
	// func (g Gerente) CalcularSalario() float64 {
	//     return 8000.0
	// }
	// 
	// func exibirDetalhes(f Funcionario) {
	//     salario := f.CalcularSalario()
	//     
	//     switch v := f.(type) {
	//     case Desenvolvedor:
	//         fmt.Printf("Desenvolvedor - Linguagem: %s, Salário: R$ %.2f\n", v.Linguagem, salario)
	//     case Gerente:
	//         fmt.Printf("Gerente - Equipe: %d pessoas, Salário: R$ %.2f\n", v.Equipe, salario)
	//     }
	// }
	// 
	// dev := Desenvolvedor{Linguagem: "Go"}
	// ger := Gerente{Equipe: 5}
	// 
	// exibirDetalhes(dev)
	// exibirDetalhes(ger)
}

