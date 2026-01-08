package array_hashing_easy

/*
	Lvl - easy

	Лучшая структура - это HashMap, так как нет явного HashSet
	Краевые случаи: пустой массив, элементов массива 1 - возращаем false
	Сложность: O(n) по времени, O(n) по памяти

	Учитывая массив целых чисел nums, верните true,
	если какое-либо значение встречается в массиве более одного раза,
	в противном случае верните false.

	Input: nums = [1, 2, 3, 3]
	Output: true

	Input: nums = [1, 2, 3, 4]
	Output: false
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
