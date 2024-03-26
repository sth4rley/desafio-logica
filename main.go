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

	fmt.Printf("\nInsira uma expressão lógica:\n> ")
	fmt.Scan(&exp)
	fmt.Printf("\n")

	notacao_polonesa := infix_to_postfix(exp)

	if !logica.Check(notacao_polonesa) {
		fmt.Printf(colorRed("Erro: A expressão inserida é inválida.\n"))
	} else {
		truthtable, fnc, fnd, classificacao := logica.Semantica(notacao_polonesa)

		fmt.Println("Notacao infixa:", exp)
		fmt.Println("Notacao polonesa:", notacao_polonesa)
		fmt.Println("Forma Normal Conjuntiva (FNC):", fnc)
		fmt.Println("Forma Normal Disjuntiva (FND):", fnd)
		fmt.Println("Classificacao:", classificacao)

		tui.ShowTruthtable(truthtable)
	}

}

func infix_to_postfix(exp string) string {
	var pOp []rune
	var precedences = map[rune]int{'~': 3, '^': 2, 'v': 2, '=': 2, '-': 2}

	newExp := ""
	for _, el := range exp {
		if el == '(' {
			pOp = append(pOp, el)
		} else if el == ')' {
			for pOp[len(pOp)-1] != '(' {
				newExp += string(pOp[len(pOp)-1])
				pOp = pOp[:len(pOp)-1]
			}
			pOp = pOp[:len(pOp)-1] // remove o '('
		} else if precedence, ok := precedences[el]; ok {
			for len(pOp) > 0 && precedences[pOp[len(pOp)-1]] >= precedence {
				newExp += string(pOp[len(pOp)-1])
				pOp = pOp[:len(pOp)-1]
			}
			pOp = append(pOp, el)
		} else {
			newExp += string(el)
		}
	}

	for len(pOp) > 0 {
		newExp += string(pOp[len(pOp)-1])
		pOp = pOp[:len(pOp)-1]
	}
	return newExp
}
