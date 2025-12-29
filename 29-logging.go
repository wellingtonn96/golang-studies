//go:build exercicio29 || ignore
// +build exercicio29 ignore

package main

/*
EXERCÍCIO 29: Logging
Link do exercício: Conceito essencial para desenvolvimento real

OBJETIVO:
Aprender a usar logging em Go para debugging e monitoramento.

CONCEITOS:
LOG PACKAGE:
- log.Print(): Log padrão
- log.Printf(): Log formatado
- log.Println(): Log com nova linha
- log.Fatal(): Log e exit
- log.Panic(): Log e panic

LOG LEVELS:
- Info: Informações gerais
- Warning: Avisos
- Error: Erros
- Debug: Debugging (precisa implementar)

CUSTOM LOGGERS:
- log.New(): Cria logger customizado
- SetOutput(): Define destino
- SetPrefix(): Define prefixo
- SetFlags(): Define formato

EXEMPLOS DE APLICAÇÃO:

1. Log Básico:
   import "log"
   
   log.Print("Mensagem de log")
   log.Printf("Valor: %d", 42)
   log.Println("Mensagem com nova linha")

2. Log para Arquivo:
   file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
   if err != nil {
       log.Fatal(err)
   }
   defer file.Close()
   
   log.SetOutput(file)
   log.Println("Esta mensagem vai para o arquivo")

3. Logger Customizado:
   logger := log.New(os.Stdout, "APP: ", log.Ldate|log.Ltime)
   logger.Println("Mensagem customizada")

4. Múltiplos Loggers:
   infoLogger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
   errorLogger := log.New(os.Stderr, "ERROR: ", log.LstdFlags)
   
   infoLogger.Println("Operação bem-sucedida")
   errorLogger.Println("Erro na operação")

5. Log Flags:
   log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
   log.Println("Mensagem")
   // 2024/01/15 10:30:45 main.go:15: Mensagem

6. Log Levels (implementação customizada):
   type LogLevel int
   
   const (
       DEBUG LogLevel = iota
       INFO
       WARNING
       ERROR
   )
   
   func logLevel(level LogLevel, format string, args ...interface{}) {
       prefix := []string{"[DEBUG]", "[INFO]", "[WARNING]", "[ERROR]"}[level]
       log.Printf("%s %s", prefix, fmt.Sprintf(format, args...))
   }

7. Structured Logging (JSON):
   type LogEntry struct {
       Timestamp string `json:"timestamp"`
       Level     string `json:"level"`
       Message   string `json:"message"`
   }
   
   entry := LogEntry{
       Timestamp: time.Now().Format(time.RFC3339),
       Level:     "INFO",
       Message:    "Operação realizada",
   }
   jsonBytes, _ := json.Marshal(entry)
   log.Println(string(jsonBytes))

8. Log com Context:
   func logWithContext(ctx context.Context, level, msg string) {
       requestID := ctx.Value("requestID")
       log.Printf("[%s] RequestID: %v - %s", level, requestID, msg)
   }

DESAFIO:
Como desenvolvedor de sistema de monitoramento,
eu quero implementar logging estruturado com diferentes níveis,
para que eu tenha visibilidade completa do sistema.

Critérios de Aceitação:
- Criar logger customizado
- Implementar níveis de log (DEBUG, INFO, WARNING, ERROR)
- Logar para arquivo e console
- Incluir timestamp e contexto

INSTRUÇÕES:
1. Crie estrutura para níveis de log
2. Crie função que aceita nível e mensagem
3. Configure logger para escrever em arquivo e console
4. Implemente diferentes níveis de log
5. Teste com diferentes mensagens e níveis
*/

import (
	"fmt"
	"log"
	"os"
	"time"
)

// TODO: Defina níveis de log e implemente logging customizado

func main() {
	// TODO: Implemente sistema de logging
	
	// Dica: Use log.New para criar loggers customizados
	// Dica: Use io.MultiWriter para múltiplos destinos
	// Dica: Use log.SetFlags para formato
	// Dica: Use time.Now() para timestamps
	
	// Exemplo de estrutura:
	// type LogLevel string
	// 
	// const (
	//     DEBUG   LogLevel = "DEBUG"
	//     INFO    LogLevel = "INFO"
	//     WARNING LogLevel = "WARNING"
	//     ERROR   LogLevel = "ERROR"
	// )
	// 
	// func setupLogger() *log.Logger {
	//     file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//     multiWriter := io.MultiWriter(os.Stdout, file)
	//     logger := log.New(multiWriter, "", log.Ldate|log.Ltime|log.Lshortfile)
	//     return logger
	// }
	// 
	// func logMessage(logger *log.Logger, level LogLevel, msg string) {
	//     timestamp := time.Now().Format("2006-01-02 15:04:05")
	//     logger.Printf("[%s] [%s] %s", timestamp, level, msg)
	// }
	// 
	// logger := setupLogger()
	// logMessage(logger, INFO, "Sistema iniciado")
	// logMessage(logger, WARNING, "Recurso com uso alto")
	// logMessage(logger, ERROR, "Falha na conexão")
}



