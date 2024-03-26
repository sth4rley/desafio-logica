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

func Classificacao(tt []map[string]bool, main_exp string) string {
	t, f := 0, 0

	for _, el := range tt {
		if el[main_exp] {
			t++
		} else {
			f++
		}
	}

	if f == 0 {
		return "tautologia"
	}
	if t == 0 {
		return "contradicao"
	}
	return "contingencia"
}

func FND(tt []map[string]bool, main_exp string) string {
	fn := ""
	props := Props(main_exp)

	for _, el := range tt {
		if el[main_exp] { // para cada interpretacao verdadeira
			fn += "("
			for c := range props { // verifica cada propriedade
				if !el[string(c)] { // se a prop for falsa adiciona a negacao
					fn += "~"
				}
				fn += string(c) + "^"
			}
			fn = fn[:len(fn)-1] + ")v"
		}
	}

	if len(fn) > 0 {
		fn = fn[:len(fn)-1]
	}

	return fn
}

func FNC(tt []map[string]bool, main_exp string) string {
	fn := ""
	props := Props(main_exp)

	for _, el := range tt {
		if !el[main_exp] { // para cada interpretacao falsa
			fn += "("
			for c := range props { // verifica cada propriedade
				if el[string(c)] { // se a prop for verdadeira adiciona a negacao
					fn += "~"
				}
				fn += string(c) + "v"
			}
			fn = fn[:len(fn)-1] + ")^"
		}
	}

	if len(fn) > 0 {
		fn = fn[:len(fn)-1]
	}

	return fn
}

func Semantica(expression string) ([]map[string]bool, string, string, string) {
	combinations := GetCombinations(expression)
	truthtable := make([]map[string]bool, 0)
	main_exp := ""

	for _, combination := range combinations {
		part, mexp := Evaluate(expression, combination)
		truthtable = append(truthtable, part)
		main_exp = mexp
	}

	return truthtable, FNC(truthtable, main_exp), FND(truthtable, main_exp), Classificacao(truthtable, main_exp)
}

func Evaluate(expression string, values map[rune]bool) (map[string]bool, string) {

	subexpressions := make(map[string]bool)

	var s Stack
	var main_exp string

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

		main_exp = exp

	}

	return subexpressions, main_exp
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
