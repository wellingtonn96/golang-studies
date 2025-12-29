//go:build exercicio17 || ignore
// +build exercicio17 ignore

package main

/*
EXERCÍCIO 17: Biblioteca Padrão (net/http, encoding/json)
Link do exercício: roadmap-com-desafios-golang.md - Etapa 3: Avançado - Biblioteca Padrão

OBJETIVO:
Aprender a criar servidores HTTP e trabalhar com JSON usando a biblioteca padrão.

CONCEITOS:
NET/HTTP:
- http.HandleFunc(): Registra handler para rota
- http.ListenAndServe(): Inicia servidor
- http.Request: Requisição recebida
- http.ResponseWriter: Resposta a enviar
- Métodos: GET, POST, PUT, DELETE, etc.

ENCODING/JSON:
- json.Marshal(): Struct → JSON ([]byte)
- json.Unmarshal(): JSON ([]byte) → Struct
- json.NewEncoder(): Encoder para streams
- json.NewDecoder(): Decoder para streams
- Tags `json:"nome"`: Controla serialização

EXEMPLOS DE APLICAÇÃO:

1. Servidor HTTP Básico:
   func handler(w http.ResponseWriter, r *http.Request) {
       fmt.Fprintf(w, "Olá, %s!", r.URL.Path[1:])
   }
   
   http.HandleFunc("/", handler)
   http.ListenAndServe(":8080", nil)

2. Handler com JSON:
   type Resposta struct {
       Mensagem string `json:"mensagem"`
       Status   int    `json:"status"`
   }
   
   func handler(w http.ResponseWriter, r *http.Request) {
       w.Header().Set("Content-Type", "application/json")
       resp := Resposta{Mensagem: "Sucesso", Status: 200}
       json.NewEncoder(w).Encode(resp)
   }

3. Marshal/Unmarshal:
   type Pessoa struct {
       Nome  string `json:"nome"`
       Idade int    `json:"idade"`
   }
   
   // Struct → JSON
   p := Pessoa{Nome: "João", Idade: 30}
   jsonBytes, _ := json.Marshal(p)
   fmt.Println(string(jsonBytes))  // {"nome":"João","idade":30}
   
   // JSON → Struct
   jsonStr := `{"nome":"Maria","idade":25}`
   var p2 Pessoa
   json.Unmarshal([]byte(jsonStr), &p2)
   fmt.Println(p2.Nome)  // Maria

4. Handler POST com JSON:
   func criarUsuario(w http.ResponseWriter, r *http.Request) {
       if r.Method != "POST" {
           http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
           return
       }
       
       var usuario Usuario
       if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
           http.Error(w, err.Error(), http.StatusBadRequest)
           return
       }
       
       // Processar usuário...
       w.Header().Set("Content-Type", "application/json")
       json.NewEncoder(w).Encode(usuario)
   }

5. Múltiplas Rotas:
   http.HandleFunc("/", homeHandler)
   http.HandleFunc("/api/usuarios", usuariosHandler)
   http.HandleFunc("/api/produtos", produtosHandler)

6. Query Parameters:
   func handler(w http.ResponseWriter, r *http.Request) {
       nome := r.URL.Query().Get("nome")
       idade := r.URL.Query().Get("idade")
       fmt.Fprintf(w, "Nome: %s, Idade: %s", nome, idade)
   }
   // Acesse: /?nome=João&idade=30

7. Headers Customizados:
   func handler(w http.ResponseWriter, r *http.Request) {
       w.Header().Set("X-Custom-Header", "valor")
       w.WriteHeader(http.StatusOK)
       fmt.Fprint(w, "Resposta")
   }

DESAFIO (User Story):
Como dono de loja virtual,
eu quero criar um server HTTP simples com JSON,
para que eu exponha APIs de produtos.

Critérios de Aceitação:
- Usar net/http para endpoint
- Marshal/unmarshal JSON com tags
- Testar com curl

INSTRUÇÕES:
1. Crie uma struct "Produto" com campos: ID, Nome, Preco, Estoque
2. Crie um slice de produtos (simulando banco de dados)
3. Crie handler GET /api/produtos que retorna lista de produtos em JSON
4. Crie handler GET /api/produtos/{id} que retorna produto específico
5. Crie handler POST /api/produtos para adicionar novo produto
6. Inicie servidor na porta 8080
7. Teste com: curl http://localhost:8080/api/produtos
*/

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// TODO: Defina a struct Produto aqui
// type Produto struct { ... }

func main() {
	// TODO: Implemente a solução do desafio aqui
	
	// Dica: Use http.HandleFunc para rotas
	// Dica: Use json.Marshal/Unmarshal para JSON
	// Dica: Use r.URL.Path para extrair ID
	// Dica: Use w.Header().Set("Content-Type", "application/json")
	
	// Exemplo de estrutura:
	// var produtos []Produto
	// produtos = append(produtos, Produto{ID: 1, Nome: "Notebook", Preco: 2500.0, Estoque: 10})
	// 
	// http.HandleFunc("/api/produtos", func(w http.ResponseWriter, r *http.Request) {
	//     if r.Method == "GET" {
	//         w.Header().Set("Content-Type", "application/json")
	//         json.NewEncoder(w).Encode(produtos)
	//     } else if r.Method == "POST" {
	//         var novo Produto
	//         json.NewDecoder(r.Body).Decode(&novo)
	//         novo.ID = len(produtos) + 1
	//         produtos = append(produtos, novo)
	//         w.Header().Set("Content-Type", "application/json")
	//         json.NewEncoder(w).Encode(novo)
	//     }
	// })
	// 
	// fmt.Println("Servidor rodando em http://localhost:8080")
	// http.ListenAndServe(":8080", nil)
}

