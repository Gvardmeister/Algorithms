package array_hashing

import (
	"strings"
	"unicode/utf8"
)

/*
	Lvl - easy

	Лучшая структура - это HashMap
	Краевые случаи: разная длина строк, строки пусты или пробел - false
	Сложность: O(n) по времени, O(n) по памяти
*/

func IsAnagram(s string, t string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	t = strings.ToLower(strings.TrimSpace(t))

	if utf8.RuneCountInString(s) != utf8.RuneCountInString(t) {
		return false
	}

	runnersS := []rune(s)
	runnersT := []rune(t)

	resultMap := make(map[rune]int, len(s+t))
	for i := 0; i < len(runnersS); i++ {
		resultMap[runnersS[i]]++
		resultMap[runnersT[i]]--
	}

	for _, count := range resultMap {
		if count != 0 {
			return false
		}
	}

	return true
}

/*
	Перенести в main(), проверено с другими тест-кейсами

	s := "racecar"
	t := "carrace"

	s2 := "jar"
	t2 := "jam"

	result := array_hashing.IsAnagram(s, t)      // true
	resultTwo := array_hashing.IsAnagram(s2, t2) // false

	fmt.Println(result)
	fmt.Println(resultTwo)

*/
