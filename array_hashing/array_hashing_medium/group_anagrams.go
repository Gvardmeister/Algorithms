package array_hashing_medium

/*
	Lvl - medium

	Лучшая структура данных - это HashMap
	Краевые случаи: изначальный массив не может быть пустым
	Сложность: O(n*k) по времени, O(n*k) по памяти

	Имея массив строк strs, сгруппируйте все анаграммы в подсписки.
	Вы можете вернуть результат в любом порядке.
	Анаграмма — это строка, которая содержит точно такие же символы,
	как и другая строка, но порядок символов может быть другим.

	Input: strs = ["act","pots","tops","cat","stop","hat"]
	Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]

	Input: strs = ["x"]
	Output: [["x"]]

	Input: strs = [""]
	Output: [[""]]
*/

func GroupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	groups := make(map[[26]int][]string)

	for _, str := range strs {
		var arr [26]int

		for i := 0; i < len(str); i++ {
			arr[str[i]-'a']++
		}

		groups[arr] = append(groups[arr], str)
	}

	result := make([][]string, 0, len(groups))

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

/*
	Перенести в main(), проверено с другими тест-кейсами

	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	strsTwo := []string{"x"}
	var strsThree []string

	result := array_hashing.GroupAnagrams(strs)
	resultTwo := array_hashing.GroupAnagrams(strsTwo)
	resultThree := array_hashing.GroupAnagrams(strsThree)

	fmt.Println(result)
	fmt.Println(resultTwo)
	fmt.Println(resultThree)
*/
