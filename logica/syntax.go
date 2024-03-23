package logica

type stack []rune

func (s *stack) push(item rune) {
	*s = append(*s, item)
}

func (s *stack) pop() rune {
	if len(*s) != 0 {
		aux := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return aux
	}
	return ' '
}

func (s *stack) top() rune {
	return (*s)[len(*s)-1]
}

func Check(expression string) bool {

	var s stack

	for _, chr := range expression {

		if LogicOperators[chr] {

			if chr == '~' {
				if len(s) < 1 {
					return false
				}
				s.pop()
			} else {
				if len(s) < 2 {
					return false
				}
				s.pop()
				s.pop()
			}

		} else {
			if chr < 'A' || chr > 'Z' {
				return false
			}
		}

		s.push('X')

	}

	return len(s) == 1
}
