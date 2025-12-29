//go:build exercicio27 || ignore
// +build exercicio27 ignore

package main

/*
EXERCÍCIO 27: Encoding (Base64, Hex)
Link do exercício: Conceito essencial para desenvolvimento real

OBJETIVO:
Aprender a codificar e decodificar dados em diferentes formatos.

CONCEITOS:
ENCODING/BASE64:
- base64.StdEncoding: Encoding padrão
- base64.URLEncoding: Encoding seguro para URLs
- EncodeToString(): Codifica para string
- DecodeString(): Decodifica de string

ENCODING/HEX:
- hex.EncodeToString(): Codifica para hexadecimal
- hex.DecodeString(): Decodifica de hexadecimal

CASOS DE USO:
- Armazenar dados binários em texto
- URLs e cookies
- Assinaturas e tokens
- Logs e debugging

EXEMPLOS DE APLICAÇÃO:

1. Base64 Encoding:
   import "encoding/base64"
   
   dados := []byte("Olá, Go!")
   encoded := base64.StdEncoding.EncodeToString(dados)
   fmt.Println(encoded)  // T2zDoSwgR28h
   
   decoded, err := base64.StdEncoding.DecodeString(encoded)
   fmt.Println(string(decoded))  // Olá, Go!

2. Base64 URL Encoding:
   dados := []byte("dados+especiais")
   encoded := base64.URLEncoding.EncodeToString(dados)
   // Seguro para usar em URLs (sem + e /)

3. Hex Encoding:
   import "encoding/hex"
   
   dados := []byte("Hello")
   encoded := hex.EncodeToString(dados)
   fmt.Println(encoded)  // 48656c6c6f
   
   decoded, err := hex.DecodeString(encoded)
   fmt.Println(string(decoded))  // Hello

4. Encoding de Struct para Base64:
   import "encoding/json"
   
   pessoa := Pessoa{Nome: "João", Idade: 30}
   jsonBytes, _ := json.Marshal(pessoa)
   encoded := base64.StdEncoding.EncodeToString(jsonBytes)
   
   // Decodificar:
   decoded, _ := base64.StdEncoding.DecodeString(encoded)
   var pessoa2 Pessoa
   json.Unmarshal(decoded, &pessoa2)

5. Comparar Encodings:
   dados := []byte("teste@exemplo.com")
   
   base64 := base64.StdEncoding.EncodeToString(dados)
   hex := hex.EncodeToString(dados)
   
   fmt.Printf("Base64: %s (tamanho: %d)\n", base64, len(base64))
   fmt.Printf("Hex: %s (tamanho: %d)\n", hex, len(hex))

6. Encoding de Arquivo:
   dados, _ := os.ReadFile("imagem.jpg")
   encoded := base64.StdEncoding.EncodeToString(dados)
   // Agora pode ser armazenado como texto

DESAFIO:
Como desenvolvedor de sistema de autenticação,
eu quero codificar e decodificar tokens e dados sensíveis,
para que eu armazene informações de forma segura.

Critérios de Aceitação:
- Codificar dados em Base64
- Codificar dados em Hex
- Decodificar ambos os formatos
- Comparar tamanhos dos encodings

INSTRUÇÕES:
1. Crie função que codifica string em Base64
2. Crie função que decodifica Base64
3. Crie função que codifica string em Hex
4. Crie função que decodifica Hex
5. Teste com diferentes strings
6. Compare tamanhos dos encodings
*/

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	// TODO: Implemente as funções de encoding/decoding
	
	// Dica: Use base64.StdEncoding para Base64
	// Dica: Use hex.EncodeToString/DecodeString para Hex
	// Dica: Sempre trate erros de decoding
	
	// Exemplo de estrutura:
	// func encodeBase64(dados []byte) string {
	//     return base64.StdEncoding.EncodeToString(dados)
	// }
	// 
	// func decodeBase64(encoded string) ([]byte, error) {
	//     return base64.StdEncoding.DecodeString(encoded)
	// }
	// 
	// func encodeHex(dados []byte) string {
	//     return hex.EncodeToString(dados)
	// }
	// 
	// func decodeHex(encoded string) ([]byte, error) {
	//     return hex.DecodeString(encoded)
	// }
	// 
	// // Testes:
	// texto := "Dados sensíveis: senha123"
	// dados := []byte(texto)
	// 
	// base64Encoded := encodeBase64(dados)
	// hexEncoded := encodeHex(dados)
	// 
	// fmt.Printf("Original: %s (tamanho: %d bytes)\n", texto, len(dados))
	// fmt.Printf("Base64: %s (tamanho: %d bytes)\n", base64Encoded, len(base64Encoded))
	// fmt.Printf("Hex: %s (tamanho: %d bytes)\n", hexEncoded, len(hexEncoded))
	// 
	// // Decodificar:
	// base64Decoded, _ := decodeBase64(base64Encoded)
	// hexDecoded, _ := decodeHex(hexEncoded)
	// 
	// fmt.Printf("Base64 decodificado: %s\n", string(base64Decoded))
	// fmt.Printf("Hex decodificado: %s\n", string(hexDecoded))
}

