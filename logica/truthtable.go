package logica

import (
	"sort"
)

var LogicOperators = map[rune]bool{
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

		if LogicOperators[chr] {

			first := s.pop()

			if chr == '~' { // se for o operador de negacao

				exp = /*"(" +*/ "~" + first.expression /* + ")"*/
				res = !first.value

			} else { // se for outro operador

				second := s.pop()

				if chr == 'v' {
					exp = "(" + second.expression + "v" + first.expression + ")"
					res = first.value || second.value
				}

				if chr == '^' {
					exp = "(" + second.expression + "^" + first.expression + ")"
					res = first.value && second.value
				}

				if chr == '-' {
					exp = "(" + second.expression + "-" + first.expression + ")"
					res = !(second.value && !first.value)
				}

				if chr == '=' {
					exp = "(" + second.expression + "=" + first.expression + ")"
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
	props := Props(exp)

	propList := make([]rune, 0, len(props))
	for prop := range props {
		propList = append(propList, prop)
	}
	sort.Slice(propList, func(i, j int) bool { return propList[i] < propList[j] })

	result := make([]map[rune]bool, 0)

	var generateComb func(int, map[rune]bool)
	generateComb = func(index int, current map[rune]bool) {

		if index == len(propList) {
			newMap := make(map[rune]bool)
			for k, v := range current {
				newMap[k] = v
			}
			result = append(result, newMap)
			return
		}

		current[propList[index]] = true
		generateComb(index+1, current)

		current[propList[index]] = false
		generateComb(index+1, current)
	}

	generateComb(0, make(map[rune]bool))

	return result
}
