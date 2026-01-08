package stack

/*
	Лучшая структура - это стек, реализация через slice []rune
	Краевые случаи: строка не может быть пустой, другие символы не допустимы
	Сложность: O(n) по времени, O(n) по памяти

	Вам дана строка s, состоящая из следующих символов: „(“, „)“, „{“, „}“, „[“ и „]“.
	Входная строка s является допустимой, если и только если:

	Каждая открытая скобка закрывается скобкой того же типа.
	Открытые скобки закрываются в правильном порядке.
	Каждая закрытая скобка имеет соответствующую открытую скобку того же типа.
	Верните true, если s является допустимой строкой, и false в противном случае.

    Input: s = "[]"
    Output: true

    Input: s = "([{}])"
    Output: true

    Input: s = "[(])"
    Output: false
*/

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}

	var stack []rune

	for _, v := range s {
		switch v {
		case '(', '{', '[':
			stack = append(stack, v)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}

			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}

			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}

			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	return len(stack) == 0
}
