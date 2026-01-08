package array_hashing_easy

import (
	"strings"
	"unicode/utf8"
)

/*
	Lvl - easy

	Лучшая структура - это HashMap
	Краевые случаи: разная длина строк, строки пусты или пробел - false
	Сложность: O(n) по времени, O(n) по памяти

	Даны две строки.
	Написать функцию, которая проверяет, являются ли одна анаграммой другой,
	то есть содержат одни и те же символы, но возможно в разном порядке.
	Например, abcdef и abdcfe - анаграммы, а abcdef и abcddef - нет.
*/

func IsAnagram(s string, t string) bool {
	// Можно добавить strings.TrimSpace()

	s = strings.ToLower(s)
	t = strings.ToLower(s)

	if utf8.RuneCountInString(s) != utf8.RuneCountInString(t) {
		return false
	}

	// Альтернатива if len(runnersS) != len(runnersT) {return false}

	resultMap := make(map[rune]int)
	runnersS := []rune(s)
	runnersT := []rune(t)

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
