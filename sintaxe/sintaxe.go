package sintaxe

func Verificar(exp string) bool {

	if len(exp) == 0 {
		return false
	}

	conectivo := map[rune]bool{
		'^': true,
		'v': true,
		'-': true,
		'=': true,
	}

	for i, el := range exp {

		if i == 0 { // primeira execucao

			// verifica se o primeiro caractere eh um conectivo
			if conectivo[el] {
				return false
			}

		} else { // demais execucoes

			// verifica se existe negacao junto com conectivo
			if exp[i-1] == '~' && conectivo[el] {
				return false
			}

			// verififica se tem dois simbolos seguidos
			if !conectivo[rune(exp[i-1])] && exp[i-1] != '~' && !conectivo[el] {
				return false
			}

			// dois conectivos seguidos
			if conectivo[rune(exp[i-1])] && conectivo[el] {
				return false
			}

		}

		// verifica se a expressao termina com negacao
		if i == len(exp)-1 && (exp[i] == '~' || conectivo[rune(exp[i])]) {
			return false
		}

	}

	return true
}
