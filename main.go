package main

import (
	"desafio/logica"
	"desafio/tui"
	"fmt"
)

func colorRed(text string) string {
	return "\033[31m" + text + "\033[0m"
}

func main() {

	var exp string

	fmt.Printf("\nInsira uma expressão lógica (notação polonesa):\n> ")
	fmt.Scan(&exp)
	fmt.Printf("\n")

	if !logica.Check(exp) {
		fmt.Printf(colorRed("Erro: A expressão inserida é inválida.\n"))
	} else {
		truthtable := logica.TruthTable(exp)
		t := 0
		f := 0

		for _, itrp := range truthtable {
			if itrp[exp] {
				t++
			} else {
				f++
			}
		}

		tui.ShowTruthtable(truthtable)

		fmt.Println("true: ", t)
		fmt.Println("false: ", f)
		if t == 0 {
			fmt.Println("contradicao")
		} else if f == 0 {
			fmt.Println("tautologia")
		} else {
			fmt.Println("contigencia")
		}

	}

}
