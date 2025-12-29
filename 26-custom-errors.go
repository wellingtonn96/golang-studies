//go:build exercicio26 || ignore
// +build exercicio26 ignore

package main

/*
EXERCÍCIO 26: Custom Errors e Error Wrapping
Link do exercício: Conceito essencial para desenvolvimento real

OBJETIVO:
Aprender a criar erros customizados e usar error wrapping para rastreamento.

CONCEITOS:
CUSTOM ERRORS:
- Implementar interface error (método Error() string)
- Adicionar campos adicionais
- Criar tipos de erro específicos

ERROR WRAPPING (Go 1.13+):
- fmt.Errorf com %w: Envolve erro original
- errors.Unwrap(): Desenrola erro
- errors.Is(): Verifica se erro é de tipo específico
- errors.As(): Converte para tipo específico

SENTINEL ERRORS:
- Erros pré-definidos como variáveis
- Comparação direta: err == ErrNotFound
- Útil para erros conhecidos

EXEMPLOS DE APLICAÇÃO:

1. Error Customizado Simples:
   type MeuErro struct {
       Mensagem string
       Codigo    int
   }
   
   func (e *MeuErro) Error() string {
       return fmt.Sprintf("Erro %d: %s", e.Codigo, e.Mensagem)
   }
   
   func exemplo() error {
       return &MeuErro{Mensagem: "Algo deu errado", Codigo: 404}
   }

2. Sentinel Errors:
   var (
       ErrNotFound   = errors.New("não encontrado")
       ErrInvalid    = errors.New("inválido")
       ErrUnauthorized = errors.New("não autorizado")
   )
   
   func buscar(id int) error {
       if id < 0 {
           return ErrInvalid
       }
       if id > 100 {
           return ErrNotFound
       }
       return nil
   }
   
   err := buscar(200)
   if err == ErrNotFound {
       fmt.Println("Item não encontrado")
   }

3. Error Wrapping:
   func processarArquivo(nome string) error {
       dados, err := os.ReadFile(nome)
       if err != nil {
           return fmt.Errorf("erro ao processar arquivo %s: %w", nome, err)
       }
       // processar dados...
       return nil
   }

4. Unwrap:
   err := processarArquivo("arquivo.txt")
   if err != nil {
       erroOriginal := errors.Unwrap(err)
       fmt.Println("Erro original:", erroOriginal)
   }

5. errors.Is():
   err := processarArquivo("arquivo.txt")
   if errors.Is(err, os.ErrNotExist) {
       fmt.Println("Arquivo não existe")
   }

6. errors.As():
   err := exemplo()
   var meuErro *MeuErro
   if errors.As(err, &meuErro) {
       fmt.Printf("Código do erro: %d\n", meuErro.Codigo)
   }

7. Múltiplos Wraps:
   func camada1() error {
       return camada2()
   }
   
   func camada2() error {
       return fmt.Errorf("camada2: %w", camada3())
   }
   
   func camada3() error {
       return errors.New("erro original")
   }

8. Error com Context:
   type ErroValidacao struct {
       Campo   string
       Mensagem string
   }
   
   func (e *ErroValidacao) Error() string {
       return fmt.Sprintf("Campo '%s': %s", e.Campo, e.Mensagem)
   }
   
   func validar(usuario Usuario) error {
       if usuario.Email == "" {
           return &ErroValidacao{
               Campo:    "email",
               Mensagem: "email é obrigatório",
           }
       }
       return nil
   }

DESAFIO:
Como desenvolvedor de API,
eu quero criar erros customizados e rastrear erros através de camadas,
para que eu tenha melhor debugging e tratamento de erros.

Critérios de Aceitação:
- Criar tipos de erro customizados
- Usar error wrapping
- Verificar tipos de erro com errors.Is/As
- Criar sentinel errors

INSTRUÇÕES:
1. Crie sentinel errors: ErrUsuarioNaoEncontrado, ErrEmailInvalido
2. Crie struct ErroValidacao com campos Campo e Mensagem
3. Crie função validarEmail que retorna erro customizado
4. Crie função buscarUsuario que retorna erro wrapped
5. Use errors.Is para verificar tipo de erro
6. Use errors.As para acessar campos do erro customizado
*/

import (
	"errors"
	"fmt"
	"strings"
)

// TODO: Defina sentinel errors e tipos de erro customizados
// var ErrUsuarioNaoEncontrado = errors.New("usuário não encontrado")
// var ErrEmailInvalido = errors.New("email inválido")
//
// type ErroValidacao struct {
//     Campo    string
//     Mensagem string
// }
//
// func (e *ErroValidacao) Error() string {
//     return fmt.Sprintf("Validação falhou no campo '%s': %s", e.Campo, e.Mensagem)
// }

func main() {
	// TODO: Implemente as funções e testes
	
	// Dica: Use fmt.Errorf com %w para wrapping
	// Dica: Use errors.Is para verificar sentinel errors
	// Dica: Use errors.As para converter para tipo customizado
	
	// Exemplo de estrutura:
	// func validarEmail(email string) error {
	//     if email == "" {
	//         return &ErroValidacao{
	//             Campo:    "email",
	//             Mensagem: "email não pode ser vazio",
	//         }
	//     }
	//     if !strings.Contains(email, "@") {
	//         return fmt.Errorf("validação de email: %w", ErrEmailInvalido)
	//     }
	//     return nil
	// }
	// 
	// func buscarUsuario(id int) error {
	//     if id < 0 {
	//         return fmt.Errorf("buscar usuário %d: %w", id, ErrUsuarioNaoEncontrado)
	//     }
	//     return nil
	// }
	// 
	// // Testes:
	// err := validarEmail("")
	// var erroValidacao *ErroValidacao
	// if errors.As(err, &erroValidacao) {
	//     fmt.Printf("Erro de validação: %s\n", erroValidacao.Mensagem)
	// }
	// 
	// err = validarEmail("sem-arroba")
	// if errors.Is(err, ErrEmailInvalido) {
	//     fmt.Println("Email inválido detectado")
	// }
	// 
	// err = buscarUsuario(-1)
	// if errors.Is(err, ErrUsuarioNaoEncontrado) {
	//     fmt.Println("Usuário não encontrado")
	// }
}



