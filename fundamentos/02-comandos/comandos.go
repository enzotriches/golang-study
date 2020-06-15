package main

import "fmt"

func main() {
	fmt.Print("Olá ")
	comandos := `
	go fix // atualiza o código de uma versão antiga anterior a go1 para uma nova versão depois de go1
	go version // exibe informações sobre sua versão de Go
	go env //exibe as variáveis de ambiente relacionados a Go
	go list // lista todos os pacotes instalados
	go run // compila os arquivos temporários e executa a aplicação
	godoc // gera a doc do pacote ~godoc fmt Printf | ~godoc -http =:8080
	`
	fmt.Printf(comandos)
}
