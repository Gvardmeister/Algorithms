package Yandex

/*
	Лучшая структура - это slice + hashmap
	Краевые случаи:
	Сложность: O(n) по времени, O(k) по памяти (пространственная сложность)

   Дана строка, например "aafbaaaaffc"
   Вывести для каждого символа в ней максимальное количество
   непрерывных повторений этого символа в строке.

   Для данной строки, например, результат будет:
   a:4
   b:1
   f:2
   c:1
*/

func maxConsecutiveCounts(str string) map[rune]int {
	strRune := []rune(str)
	m := make(map[rune]int)

	if len(strRune) == 0 {
		return m
	}

	currChar := strRune[0]
	currCount := 1

	for _, v := range strRune[1:] {
		if v == currChar {
			currCount++
		} else {
			if currCount > m[currChar] {
				m[currChar] = currCount
			}

			currChar = v
			currCount = 1
		}
	}

	if currCount > m[currChar] {
		m[currChar] = currCount
	}

	return m
}
