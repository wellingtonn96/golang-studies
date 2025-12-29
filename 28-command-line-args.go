//go:build exercicio28 || ignore
// +build exercicio28 ignore

package main

/*
EXERCÍCIO 28: Command Line Arguments e Flags
Link do exercício: Conceito essencial para desenvolvimento real

OBJETIVO:
Aprender a processar argumentos de linha de comando e flags.

CONCEITOS:
OS.ARGS:
- os.Args: Slice com argumentos
- os.Args[0]: Nome do programa
- os.Args[1:]: Argumentos passados

FLAG PACKAGE:
- flag.String(): Flag do tipo string
- flag.Int(): Flag do tipo int
- flag.Bool(): Flag do tipo bool
- flag.Parse(): Processa flags
- flag.Usage: Função de ajuda customizada

EXEMPLOS DE APLICAÇÃO:

1. os.Args Básico:
   import "os"
   
   if len(os.Args) < 2 {
       fmt.Println("Uso: programa <nome>")
       os.Exit(1)
   }
   
   nome := os.Args[1]
   fmt.Printf("Olá, %s!\n", nome)

2. Processar Múltiplos Args:
   for i, arg := range os.Args {
       fmt.Printf("Args[%d] = %s\n", i, arg)
   }

3. Flags Básicos:
   import "flag"
   
   nome := flag.String("nome", "Mundo", "Nome para saudação")
   idade := flag.Int("idade", 0, "Idade da pessoa")
   ativo := flag.Bool("ativo", false, "Se está ativo")
   
   flag.Parse()
   
   fmt.Printf("Nome: %s\n", *nome)
   fmt.Printf("Idade: %d\n", *idade)
   fmt.Printf("Ativo: %v\n", *ativo)

4. Flags com Var:
   var nome string
   var idade int
   
   flag.StringVar(&nome, "nome", "Mundo", "Nome")
   flag.IntVar(&idade, "idade", 0, "Idade")
   flag.Parse()
   
   fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)

5. Flags Obrigatórios:
   nome := flag.String("nome", "", "Nome (obrigatório)")
   flag.Parse()
   
   if *nome == "" {
       fmt.Println("Erro: --nome é obrigatório")
       flag.Usage()
       os.Exit(1)
   }

6. Custom Usage:
   flag.Usage = func() {
       fmt.Fprintf(os.Stderr, "Uso: %s [opções]\n", os.Args[0])
       fmt.Fprintln(os.Stderr, "Opções:")
       flag.PrintDefaults()
   }

7. Subcommands:
   import "flag"
   
   comando := flag.NewFlagSet("comando", flag.ExitOnError)
   nome := comando.String("nome", "", "Nome")
   
   if len(os.Args) < 2 {
       fmt.Println("Uso: programa <comando>")
       os.Exit(1)
   }
   
   switch os.Args[1] {
   case "criar":
       comando.Parse(os.Args[2:])
       fmt.Printf("Criando: %s\n", *nome)
   case "listar":
       fmt.Println("Listando...")
   default:
       fmt.Println("Comando desconhecido")
   }

8. Environment Variables:
   import "os"
   
   valor := os.Getenv("VARIAVEL")
   if valor == "" {
       valor = "valor padrão"
   }
   
   os.Setenv("NOVA_VAR", "valor")

DESAFIO:
Como desenvolvedor de CLI tool,
eu quero processar argumentos e flags de linha de comando,
para que usuários possam configurar o programa facilmente.

Critérios de Aceitação:
- Processar argumentos posicionais
- Processar flags opcionais
- Validar argumentos obrigatórios
- Exibir ajuda quando necessário

INSTRUÇÕES:
1. Crie programa que aceita argumento posicional (nome do arquivo)
2. Adicione flags: --output (string), --verbose (bool), --count (int)
3. Valide se arquivo foi fornecido
4. Exiba informações processadas
5. Implemente custom usage
*/

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// TODO: Implemente processamento de argumentos e flags
	
	// Dica: Use os.Args para argumentos posicionais
	// Dica: Use flag package para flags
	// Dica: Use flag.Parse() antes de usar flags
	// Dica: Use flag.Usage para ajuda customizada
	
	// Exemplo de estrutura:
	// output := flag.String("output", "saida.txt", "Arquivo de saída")
	// verbose := flag.Bool("verbose", false, "Modo verboso")
	// count := flag.Int("count", 1, "Número de iterações")
	// 
	// flag.Usage = func() {
	//     fmt.Fprintf(os.Stderr, "Uso: %s [opções] <arquivo>\n", os.Args[0])
	//     fmt.Fprintln(os.Stderr, "Opções:")
	//     flag.PrintDefaults()
	// }
	// 
	// flag.Parse()
	// 
	// if len(flag.Args()) < 1 {
	//     fmt.Println("Erro: arquivo é obrigatório")
	//     flag.Usage()
	//     os.Exit(1)
	// }
	// 
	// arquivo := flag.Args()[0]
	// 
	// fmt.Printf("Processando arquivo: %s\n", arquivo)
	// fmt.Printf("Saída: %s\n", *output)
	// fmt.Printf("Verbose: %v\n", *verbose)
	// fmt.Printf("Count: %d\n", *count)
}



