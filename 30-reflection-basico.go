//go:build exercicio30 || ignore
// +build exercicio30 ignore

package main

/*
EXERCÍCIO 30: Reflection Básico
Link do exercício: Conceito avançado para casos específicos

OBJETIVO:
Aprender o básico de reflection para casos onde tipos são desconhecidos em tempo de compilação.

CONCEITOS:
REFLECT PACKAGE:
- reflect.TypeOf(): Obtém tipo
- reflect.ValueOf(): Obtém valor
- Kind(): Tipo subjacente
- Interface(): Converte de volta para interface{}

CASOS DE USO:
- Serialização genérica
- Validação de structs
- Debugging
- Frameworks e bibliotecas

AVISOS:
- Reflection é lento
- Perde type safety
- Use apenas quando necessário
- Prefira interfaces quando possível

EXEMPLOS DE APLICAÇÃO:

1. TypeOf Básico:
   import "reflect"

   var x int = 42
   t := reflect.TypeOf(x)
   fmt.Println(t)  // int

   var s string = "hello"
   t2 := reflect.TypeOf(s)
   fmt.Println(t2)  // string

2. ValueOf Básico:
   var x int = 42
   v := reflect.ValueOf(x)
   fmt.Println(v.Int())  // 42

3. Kind:
   var x int = 42
   v := reflect.ValueOf(x)
   fmt.Println(v.Kind())  // reflect.Int

4. Inspecionar Struct:
   type Pessoa struct {
       Nome  string
       Idade int
   }

   p := Pessoa{Nome: "João", Idade: 30}
   v := reflect.ValueOf(p)
   t := reflect.TypeOf(p)

   for i := 0; i < v.NumField(); i++ {
       field := v.Field(i)
       fieldType := t.Field(i)
       fmt.Printf("%s: %v (tipo: %s)\n", fieldType.Name, field.Interface(), fieldType.Type)
   }

5. Modificar Valores (ponteiro necessário):
   var x int = 42
   v := reflect.ValueOf(&x).Elem()
   v.SetInt(100)
   fmt.Println(x)  // 100

6. Chamar Método via Reflection:
   type Calculadora struct{}

   func (c Calculadora) Soma(a, b int) int {
       return a + b
   }

   c := Calculadora{}
   v := reflect.ValueOf(c)
   metodo := v.MethodByName("Soma")
   args := []reflect.Value{
       reflect.ValueOf(10),
       reflect.ValueOf(20),
   }
   resultado := metodo.Call(args)
   fmt.Println(resultado[0].Int())  // 30

7. Verificar se é Nil:
   var x *int = nil
   v := reflect.ValueOf(x)
   fmt.Println(v.IsNil())  // true

8. Type Assertion com Reflection:
   var x interface{} = 42
   v := reflect.ValueOf(x)
   if v.Kind() == reflect.Int {
       fmt.Println(v.Int())
   }

DESAFIO:
Como desenvolvedor de framework de validação,
eu quero usar reflection para validar structs dinamicamente,
para que eu crie validações genéricas.

Critérios de Aceitação:
- Inspecionar campos de struct
- Validar tipos de campos
- Acessar valores via reflection
- Exibir informações da struct

INSTRUÇÕES:
1. Crie struct com múltiplos campos de diferentes tipos
2. Crie função que inspeciona struct usando reflection
3. Exiba nome, tipo e valor de cada campo
4. Valide se campos obrigatórios estão preenchidos
5. Exiba informações completas da struct
*/

// TODO: Defina struct para inspecionar
// type Usuario struct {
//     Nome  string
//     Email string
//     Idade int
//     Ativo bool
// }

func main() {
	// TODO: Implemente inspeção de struct com reflection
	
	// Dica: Use reflect.TypeOf e reflect.ValueOf
	// Dica: Use NumField() e Field() para iterar campos
	// Dica: Use Kind() para verificar tipo
	// Dica: Use Interface() para obter valor
	
	// Exemplo de estrutura:
	// func inspecionarStruct(s interface{}) {
	//     v := reflect.ValueOf(s)
	//     t := reflect.TypeOf(s)
	//     
	//     if v.Kind() == reflect.Ptr {
	//         v = v.Elem()
	//         t = t.Elem()
	//     }
	//     
	//     if v.Kind() != reflect.Struct {
	//         fmt.Println("Não é uma struct")
	//         return
	//     }
	//     
	//     fmt.Printf("Struct: %s\n", t.Name())
	//     fmt.Println("Campos:")
	//     
	//     for i := 0; i < v.NumField(); i++ {
	//         field := v.Field(i)
	//         fieldType := t.Field(i)
	//         
	//         fmt.Printf("  %s: %v (tipo: %s, kind: %s)\n",
	//             fieldType.Name,
	//             field.Interface(),
	//             fieldType.Type,
	//             field.Kind(),
	//         )
	//         
	//         // Validação básica
	//         if field.Kind() == reflect.String && field.String() == "" {
	//             fmt.Printf("    ⚠️ Campo %s está vazio\n", fieldType.Name)
	//         }
	//     }
	// }
	// 
	// usuario := Usuario{
	//     Nome:  "João",
	//     Email: "joao@email.com",
	//     Idade: 30,
	//     Ativo: true,
	// }
	// 
	// inspecionarStruct(usuario)
}



