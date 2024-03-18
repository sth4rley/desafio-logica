package logica

import (
	"sort"
)

var logicOperators = map[rune]bool{
	'~': true,
	'v': true,
	'^': true,
	'-': true,
	'=': true,
}

func TruthTable(expression string) []map[string]bool {

	combinations := GetCombinations(expression)

	truthtable := make([]map[string]bool, 0)

	for _, combination := range combinations {
		truthtable = append(truthtable, Evaluate(expression, combination))
	}

	return truthtable
}

func Evaluate(expression string, values map[rune]bool) map[string]bool {

	subexpressions := make(map[string]bool)

	var s Stack

	for _, chr := range expression {

		var exp string
		var res bool

		if logicOperators[chr] {

			first := s.pop()

			if chr == '~' { // se for o operador de negacao

				exp = first.expression + "~"
				res = !first.value

			} else { // se for outro operador

				second := s.pop()

				if chr == 'v' {
					exp = second.expression + first.expression + "v"
					res = first.value || second.value
				}

				if chr == '^' {
					exp = second.expression + first.expression + "^"
					res = first.value && second.value
				}

				if chr == '-' {
					exp = second.expression + first.expression + "-"
					res = !(second.value && !first.value)
				}

				if chr == '=' {
					exp = second.expression + first.expression + "="
					res = first.value == second.value
				}

			}

		} else {

			exp = string(chr)
			res = values[chr]

		}

		s.push(Node{exp, res})
		subexpressions[exp] = res

	}

	return subexpressions
}

func Props(expression string) map[rune]bool {
	result := make(map[rune]bool)

	for _, chr := range expression {
		if chr >= 'A' && chr <= 'Z' {
			result[chr] = true
		}
	}

	return result
}

func GetCombinations(exp string) []map[rune]bool {
	// Obtém as proposições únicas na expressão
	props := Props(exp)
	
	// Cria uma lista ordenada de proposições
	propList := make([]rune, 0, len(props))
	for prop := range props {
		propList = append(propList, prop)
	}
	sort.Slice(propList, func(i, j int) bool { return propList[i] < propList[j] })

	// Inicializa uma lista para armazenar as combinações
	result := make([]map[rune]bool, 0)

	// Função recursiva para gerar as combinações
	var generateComb func(int, map[rune]bool)
	generateComb = func(index int, current map[rune]bool) {
	
		// Quando todas as proposições foram processadas, adiciona a combinação resultante à lista de resultados
		if index == len(propList) {
			newMap := make(map[rune]bool)
			for k, v := range current {
				newMap[k] = v
			}
			result = append(result, newMap)
			return
		}

		// Define a proposição atual como verdadeira e gera as combinações para o próximo índice
		current[propList[index]] = true
		generateComb(index+1, current)
		
		// Define a proposição atual como falsa e gera as combinações para o próximo índice
		current[propList[index]] = false
		generateComb(index+1, current)
	}

	// Inicia o processo de geração de combinações com o índice inicial e um mapa vazio
	generateComb(0, make(map[rune]bool))

	return result
}
