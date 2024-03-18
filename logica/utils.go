package logica

type Node struct {
	expression string
	value      bool
}

type Stack []Node

func (s *Stack) push(item Node) {
	*s = append(*s, item)
}

func (s *Stack) pop() Node {
	if len(*s) != 0 {
		aux := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return aux
	}
	return Node{}
}

func (s *Stack) top() Node {
	return (*s)[len(*s)-1]
}
