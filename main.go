package main

import (
	"desafio/sintaxe"
	"fmt"
	"os"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

func main() {
	var exp string
	fmt.Print("Insira uma expressao logica: ")

	if _, err := fmt.Scan(&exp); err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		os.Exit(1)
	}

	if sintaxe.Verificar(exp) {
		fmt.Println(Green + "A expressao " + exp + " eh valida" + Reset)
	} else {
		fmt.Println(Red + "A expressao " + exp + " eh invalida" + Reset)
	}

	fmt.Println()
}
