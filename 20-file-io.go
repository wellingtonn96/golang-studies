//go:build exercicio20 || ignore
// +build exercicio20 ignore

package main

/*
EXERCÍCIO 20: File I/O (Leitura e Escrita de Arquivos)
Link do exercício: Conceito essencial para desenvolvimento real

OBJETIVO:
Aprender a ler e escrever arquivos em Go.

CONCEITOS:
OS PACKAGE:
- os.Open(): Abre arquivo para leitura
- os.Create(): Cria arquivo para escrita
- os.OpenFile(): Abre com opções específicas
- file.Close(): Fecha arquivo (sempre feche!)

IO PACKAGE:
- io.ReadAll(): Lê todo conteúdo
- io.Copy(): Copia dados
- io.WriteString(): Escreve string

IOUTIL (Deprecated, mas ainda usado):
- ioutil.ReadFile(): Lê arquivo inteiro
- ioutil.WriteFile(): Escreve arquivo inteiro
- Use os.ReadFile() e os.WriteFile() no Go 1.16+

BUFIO:
- bufio.Scanner: Lê linha por linha
- bufio.Reader: Leitura bufferizada
- bufio.Writer: Escrita bufferizada

EXEMPLOS DE APLICAÇÃO:

1. Ler Arquivo Completo (Go 1.16+):
   dados, err := os.ReadFile("arquivo.txt")
   if err != nil {
       log.Fatal(err)
   }
   fmt.Println(string(dados))

2. Escrever Arquivo Completo:
   conteudo := []byte("Olá, Go!")
   err := os.WriteFile("saida.txt", conteudo, 0644)
   if err != nil {
       log.Fatal(err)
   }

3. Ler com Scanner (linha por linha):
   file, err := os.Open("arquivo.txt")
   if err != nil {
       log.Fatal(err)
   }
   defer file.Close()
   
   scanner := bufio.NewScanner(file)
   for scanner.Scan() {
       linha := scanner.Text()
       fmt.Println(linha)
   }
   
   if err := scanner.Err(); err != nil {
       log.Fatal(err)
   }

4. Escrever com Writer:
   file, err := os.Create("saida.txt")
   if err != nil {
       log.Fatal(err)
   }
   defer file.Close()
   
   writer := bufio.NewWriter(file)
   writer.WriteString("Linha 1\n")
   writer.WriteString("Linha 2\n")
   writer.Flush()  // Importante!

5. Verificar se Arquivo Existe:
   if _, err := os.Stat("arquivo.txt"); os.IsNotExist(err) {
       fmt.Println("Arquivo não existe")
   }

6. Ler JSON de Arquivo:
   dados, _ := os.ReadFile("dados.json")
   var pessoa Pessoa
   json.Unmarshal(dados, &pessoa)

7. Escrever JSON em Arquivo:
   pessoa := Pessoa{Nome: "João", Idade: 30}
   dados, _ := json.MarshalIndent(pessoa, "", "  ")
   os.WriteFile("pessoa.json", dados, 0644)

8. Append em Arquivo:
   file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
   if err != nil {
       log.Fatal(err)
   }
   defer file.Close()
   
   file.WriteString("Nova linha\n")

DESAFIO:
Como desenvolvedor de sistema de logs,
eu quero ler e escrever arquivos de log,
para que eu registre eventos do sistema.

Critérios de Aceitação:
- Ler arquivo de log existente
- Adicionar novas entradas
- Contar linhas do arquivo
- Exibir estatísticas

INSTRUÇÕES:
1. Crie um arquivo "log.txt" com algumas linhas de exemplo
2. Leia o arquivo e conte quantas linhas tem
3. Adicione novas entradas de log com timestamp
4. Leia novamente e exiba todas as linhas
5. Exiba estatísticas (total de linhas, última entrada)
*/

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// TODO: Implemente a solução do desafio aqui
	
	// Dica: Use os.ReadFile ou bufio.Scanner para ler
	// Dica: Use os.OpenFile com os.O_APPEND para adicionar
	// Dica: Use time.Now().Format() para timestamp
	// Dica: Sempre feche arquivos com defer
	
	// Exemplo de estrutura:
	// // Ler e contar linhas
	// file, err := os.Open("log.txt")
	// if err != nil {
	//     // Criar arquivo se não existir
	//     file, _ = os.Create("log.txt")
	// }
	// defer file.Close()
	// 
	// scanner := bufio.NewScanner(file)
	// contador := 0
	// for scanner.Scan() {
	//     contador++
	// }
	// 
	// // Adicionar nova entrada
	// logFile, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// defer logFile.Close()
	// 
	// timestamp := time.Now().Format("2006-01-02 15:04:05")
	// logFile.WriteString(fmt.Sprintf("[%s] Nova entrada de log\n", timestamp))
	// 
	// fmt.Printf("Total de linhas: %d\n", contador+1)
}

