package array_hashing

/*
	Lvl - easy

	Лучшая структура - это HashMap
	Краевые случаи:
	Сложность: O(n) по времени, O(n) по памяти
*/

func TwoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	rm := make(map[int]int, len(nums))
	for key, value := range nums {
		comp := target - value

		if idx, ok := rm[comp]; ok {
			return []int{idx, key}
		}

		rm[value] = key
	}

	return []int{}
}

/*
	Перенести в main(), проверено с другими тест-кейсами

	nums := []int{3, 4, 5, 6}
	target := 7
	numsTwo := []int{5, 5}
	targetTwo := 10

	result := array_hashing.TwoSum(nums, target)
	resultTwo := array_hashing.TwoSum(numsTwo, targetTwo)

	fmt.Println(result)
	fmt.Println(resultTwo)
*/
