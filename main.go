package main

import (
	"desafio/logica"
	"desafio/tui"
	"fmt"
)

func colorRed(text string) string {
	return "\033[31m" + text + "\033[0m"
}

func ntop2(exp string) string {
	// notacao normal para polonesa
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

func ntop(exp string) string {
	// notacao normal para polonesa
	var pOp []rune
	newExp := ""
	for _, el := range exp {
		if logica.LogicOperators[el] {
			pOp = append(pOp, el)
		} else {
			// se for parenteses
			if el == '(' || el == ')' {
				if el == ')' {
					// desempilha o operador
					newExp += string(pOp[len(pOp)-1])
					pOp = pOp[:len(pOp)-1]
				}
			} else { // se nao for (variaveis)
				newExp += string(el)
			}
		}
	}
	// se sobrou algum operador
	for len(pOp) > 0 {
		newExp += string(pOp[len(pOp)-1])
		pOp = pOp[:len(pOp)-1]
	}
	return newExp
}

func main() {

	var exp string

	fmt.Printf("\nInsira uma expressão lógica (notação polonesa):\n> ")
	fmt.Scan(&exp)
	fmt.Printf("\n")

	newExp := ntop2(exp)

	fmt.Println(exp + " <-> " + newExp)

	if !logica.Check(newExp) {
		fmt.Printf(colorRed("Erro: A expressão inserida é inválida.\n"))
	} else {
		truthtable := logica.TruthTable(newExp)
		t := 0
		f := 0

		// truthtable pode ja retornar a formula principal

		// verificar as ultimas formulas
		for _, itrp := range truthtable {
			exp2 := ""
			for chave := range itrp {
				exp2 = chave
			}
			if itrp[exp2] {
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
