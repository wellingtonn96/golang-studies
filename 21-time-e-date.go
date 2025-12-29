//go:build exercicio21 || ignore
// +build exercicio21 ignore

package main

/*
EXERCÍCIO 21: Time e Date Operations
Link do exercício: Conceito essencial para desenvolvimento real

OBJETIVO:
Aprender a trabalhar com datas e horas em Go.

CONCEITOS:
TIME PACKAGE:
- time.Now(): Data/hora atual
- time.Date(): Cria data específica
- time.Parse(): Converte string para time.Time
- time.Format(): Formata time.Time para string
- time.Duration: Representa duração
- time.Sleep(): Pausa execução

LAYOUTS:
- Go usa data de referência: "Mon Jan 2 15:04:05 MST 2006"
- "2006-01-02": Formato ISO (YYYY-MM-DD)
- "02/01/2006": Formato brasileiro (DD/MM/YYYY)
- "15:04:05": Formato 24h (HH:MM:SS)

DURAÇÕES:
- time.Second, time.Minute, time.Hour
- time.Millisecond, time.Microsecond, time.Nanosecond
- time.Duration pode ser somado/subtraído

EXEMPLOS DE APLICAÇÃO:

1. Data/Hora Atual:
   agora := time.Now()
   fmt.Println(agora)  // 2024-01-15 10:30:45.123456789 -0300 -03

2. Formatar Data:
   agora := time.Now()
   fmt.Println(agora.Format("2006-01-02"))           // 2024-01-15
   fmt.Println(agora.Format("02/01/2006"))           // 15/01/2024
   fmt.Println(agora.Format("15:04:05"))              // 10:30:45
   fmt.Println(agora.Format("02/01/2006 15:04:05"))  // 15/01/2024 10:30:45

3. Criar Data Específica:
   data := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
   fmt.Println(data)

4. Parse de String:
   dataStr := "2024-01-15"
   data, err := time.Parse("2006-01-02", dataStr)
   if err != nil {
       log.Fatal(err)
   }
   fmt.Println(data)

5. Operações com Datas:
   agora := time.Now()
   amanha := agora.Add(24 * time.Hour)
   ontem := agora.Add(-24 * time.Hour)
   
   diferenca := amanha.Sub(agora)
   fmt.Println(diferenca)  // 24h0m0s

6. Comparar Datas:
   data1 := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
   data2 := time.Date(2024, 1, 16, 0, 0, 0, 0, time.UTC)
   
   if data1.Before(data2) {
       fmt.Println("data1 é anterior")
   }
   
   if data2.After(data1) {
       fmt.Println("data2 é posterior")
   }

7. Extrair Componentes:
   agora := time.Now()
   ano := agora.Year()
   mes := agora.Month()
   dia := agora.Day()
   hora := agora.Hour()
   minuto := agora.Minute()
   segundo := agora.Second()
   
   fmt.Printf("%d/%d/%d %d:%d:%d\n", dia, mes, ano, hora, minuto, segundo)

8. Timezone:
   utc := time.Now().UTC()
   local := time.Now().Local()
   
   loc, _ := time.LoadLocation("America/Sao_Paulo")
   sp := time.Now().In(loc)
   fmt.Println(sp)

9. Timer e Ticker:
   // Timer (executa uma vez):
   timer := time.NewTimer(2 * time.Second)
   <-timer.C
   fmt.Println("Timer expirado")
   
   // Ticker (executa repetidamente):
   ticker := time.NewTicker(1 * time.Second)
   for i := 0; i < 5; i++ {
       <-ticker.C
       fmt.Println("Tick", i)
   }
   ticker.Stop()

DESAFIO:
Como desenvolvedor de sistema de agendamento,
eu quero calcular diferenças entre datas e formatar timestamps,
para que eu gerencie eventos e prazos.

Critérios de Aceitação:
- Criar data de evento
- Calcular dias até o evento
- Formatar em diferentes formatos
- Verificar se evento já passou

INSTRUÇÕES:
1. Defina uma data de evento futuro
2. Calcule quantos dias faltam até o evento
3. Formate a data em formato brasileiro (DD/MM/YYYY HH:MM)
4. Verifique se o evento já passou
5. Exiba todas as informações formatadas
*/

import (
	"fmt"
	"time"
)

func main() {
	// TODO: Implemente a solução do desafio aqui
	
	// Dica: Use time.Date() para criar data específica
	// Dica: Use time.Now() para data atual
	// Dica: Use Sub() para calcular diferença
	// Dica: Use Format() para formatar
	// Dica: Use Before()/After() para comparar
	
	// Exemplo de estrutura:
	// evento := time.Date(2024, 12, 25, 18, 0, 0, 0, time.Local)
	// agora := time.Now()
	// 
	// diferenca := evento.Sub(agora)
	// dias := int(diferenca.Hours() / 24)
	// 
	// formatoBR := evento.Format("02/01/2006 15:04")
	// 
	// fmt.Printf("Evento: %s\n", formatoBR)
	// fmt.Printf("Dias restantes: %d\n", dias)
	// 
	// if evento.After(agora) {
	//     fmt.Println("Evento ainda não ocorreu")
	// } else {
	//     fmt.Println("Evento já passou")
	// }
}

