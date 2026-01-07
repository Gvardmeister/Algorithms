package array_hashing

/*
	Lvl - easy

	Лучшая структура - это HashMap, так как нет явного HashSet
	Краевые случаи: пустой массив, элементов массива 1 - возращаем false
	Сложность: O(n) по времени, O(n) по памяти
*/

func HasDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	resultMap := make(map[int]bool)
	for _, value := range nums {
		if _, ok := resultMap[value]; ok {
			return true
		}

		resultMap[value] = true
	}

	return false
}

/*
	Перенести в main(), проверено с другими тест-кейсами

	nums := []int{1, 2, 3, 3}       // true
	numsTwo := []int{1, 2, 3, 4}    // false

	resultOne := array_hashing.HasDuplicate(nums)
	resultTwo := array_hashing.HasDuplicate(numsTwo)

	fmt.Println(resultOne)
	fmt.Println(resultTwo)
*/
