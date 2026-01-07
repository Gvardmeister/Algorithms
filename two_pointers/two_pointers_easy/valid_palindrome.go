package two_pointers_easy

import (
	"strings"
	"unicode"
)

/*
	Lvl - easy

	Лучшая структура - это Slice rune и 2 указателя
	Краевые случаи: строка не может быть пустой
	Сложность: O(n) по памяти, O(n) пространственная сложность (по памяти)
*/

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	strRune := []rune(str)

	if len(strRune) == 0 {
		return true
	}

	left := 0
	right := len(strRune) - 1

	for left < right {
		for left < right && !unicode.IsLetter(strRune[left]) &&
			!unicode.IsDigit(strRune[left]) {
			left++
		}

		for left < right && !unicode.IsLetter(strRune[right]) &&
			!unicode.IsDigit(strRune[right]) {
			right--
		}

		if strRune[left] != strRune[right] {
			return false
		}

		left++
		right--
	}

	return true
}

/*
	Перенести в main(), проверено с другими тест-кейсами

	s := "Was it a car or a cat I saw?" // true
	sTwo := "tab a cat" // false

	fmt.Println(IsPalindrome(s))
	fmt.Println(IsPalindrome(sTwo))
*/
