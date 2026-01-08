package Yandex

import "strings"

/*
	Лучшая структура - это HashMap
	Краевые случаи: разная длина строк, строки пусты или пробел - false
	Сложность: O(n) по времени, O(n) по памяти

	Даны две строки.
	Написать функцию, которая проверяет, являются ли одна анаграммой другой,
	то есть содержат одни и те же символы, но возможно в разном порядке.
	Например, abcdef и abdcfe - анаграммы, а abcdef и abcddef - нет.
*/

func isAnagram(s string, t string) bool {
	s = strings.ToLower(s)
	t = strings.ToLower(t)

	m := make(map[rune]int)
	sRunes := []rune(s)
	tRunes := []rune(t)

	for i := 0; i < len(sRunes); i++ {
		m[sRunes[i]]++
		m[tRunes[i]]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
