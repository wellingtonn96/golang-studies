//go:build exercicio05 || ignore
// +build exercicio05 ignore

package main

/*
EXERCÍCIO 05: Controle de Fluxo (if/else, switch)
Link do exercício: roadmap-com-desafios-golang.md - Etapa 1: Básico - Controle de Fluxo

OBJETIVO:
Aprender a controlar o fluxo de execução com condicionais.

CONCEITOS:
IF/ELSE:
- if: Executa código se condição for true
- else: Executa código se condição for false
- else if: Adiciona condições adicionais
- if com inicialização: if var := valor; condição { }

SWITCH:
- switch: Compara valor com múltiplos casos
- case: Define um caso específico
- default: Caso padrão quando nenhum case corresponde
- switch sem expressão: Equivale a múltiplos if/else
- fallthrough: Continua para o próximo case

EXEMPLOS DE APLICAÇÃO:

1. If/Else básico:
   idade := 18
   if idade >= 18 {
       fmt.Println("Maior de idade")
   } else {
       fmt.Println("Menor de idade")
   }

2. If com inicialização:
   if idade := 20; idade >= 18 {
       fmt.Println("Maior de idade")
   }
   // idade não está disponível aqui

3. If/Else if/Else:
   nota := 85
   if nota >= 90 {
       fmt.Println("A")
   } else if nota >= 80 {
       fmt.Println("B")
   } else if nota >= 70 {
       fmt.Println("C")
   } else {
       fmt.Println("Reprovado")
   }

4. Switch básico:
   dia := 3
   switch dia {
   case 1:
       fmt.Println("Domingo")
   case 2:
       fmt.Println("Segunda")
   case 3:
       fmt.Println("Terça")
   default:
       fmt.Println("Outro dia")
   }

5. Switch sem expressão (múltiplas condições):
   idade := 25
   switch {
   case idade < 18:
       fmt.Println("Menor")
   case idade >= 18 && idade < 65:
       fmt.Println("Adulto")
   default:
       fmt.Println("Idoso")
   }

6. Switch com múltiplos valores:
   letra := "a"
   switch letra {
   case "a", "e", "i", "o", "u":
       fmt.Println("Vogal")
   default:
       fmt.Println("Consoante")
   }

DESAFIO (User Story):
Como gerente de RH,
eu quero classificar funcionários por cargo e departamento para calcular bônus,
para que eu possa automatizar parte do processo de folha de pagamento.

Critérios de Aceitação:
- Usar if/else para faixas salariais
- Usar switch para bônus por departamento
- Exibir cargo, salário e bônus

INSTRUÇÕES:
1. Defina cargo, salário e departamento de um funcionário
2. Use if/else para classificar faixas salariais (ex: < 3000, 3000-5000, > 5000)
3. Use switch para calcular bônus por departamento (ex: TI: 10%, Vendas: 15%, RH: 5%)
4. Calcule o bônus e exiba todas as informações

OPCIONAL (Recomendado para prática avançada):
- Considere criar uma struct "Funcionario" para agrupar os dados
- Exemplo: type Funcionario struct { Cargo string; Salario float64; Departamento string }
- Isso torna o código mais organizado e facilita adicionar mais campos depois
*/

func main() {
	// TODO: Implemente a solução do desafio aqui

	// Dica: Use if/else para classificar salário
	// Dica: Use switch para calcular bônus por departamento
	// Dica: Calcule valor do bônus = salario * (percentual / 100)

	// Exemplo de estrutura:
	// cargo := "Desenvolvedor"
	// salario := 4500.0
	// departamento := "TI"
	//
	// var faixaSalarial string
	// if salario < 3000 {
	//     faixaSalarial = "Júnior"
	// } else if salario <= 5000 {
	//     faixaSalarial = "Pleno"
	// } else {
	//     faixaSalarial = "Sênior"
	// }
	//
	// var percentualBonus float64
	// switch departamento {
	// case "TI":
	//     percentualBonus = 10.0
	// case "Vendas":
	//     percentualBonus = 15.0
	// case "RH":
	//     percentualBonus = 5.0
	// default:
	//     percentualBonus = 0.0
	// }
	//
	// bonus := salario * (percentualBonus / 100)
	// fmt.Printf("Cargo: %s\nSalário: R$ %.2f\nFaixa: %s\nDepartamento: %s\nBônus: R$ %.2f (%.1f%%)\n", ...)

	type Employee struct {
		job        string
		salary     float64
		department string
		bonus      int
	}

	employee := []Employee{
		{
			job:        "Programador",
			salary:     3000.00,
			department: "TI",
			bonus:      10,
		},
		{
			job:        "Auxiliar administrativo",
			salary:     5000.00,
			department: "RH",
			bonus:      10,
		},
		{
			job:        "Programador",
			salary:     2000.000,
			department: "Vendas",
			bonus:      20,
		},
	}

}
