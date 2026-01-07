package array_hashing

/*
	Lvl - medium

	Лучшая структура данных - это HashMap, Buckets(слайс слайсов)
	Краевые случаи: числа могут быть уникальными, одинаковыми, и пустой слайс
	Сложность: O(n) по времени, O(n+m) по памяти, но по факту O(n), так как m < n
*/

func TopKElements(nums []int, k int) []int {
	rm := make(map[int]int, len(nums))
	for _, value := range nums {
		rm[value]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, freq := range rm {
		buckets[freq] = append(buckets[freq], num)
	}

	result := make([]int, 0, k)
	for i := len(buckets) - 1; i > 0; i-- {
		for _, num := range buckets[i] {
			result = append(result, num)

			if len(result) == k {
				return result
			}
		}
	}

	return result
}

/*
	Перенести в main(), проверено с другими тест-кейсами

	nums := []int{1, 2, 2, 3, 3, 3}
	k := 2
	numsTwo := []int{7, 7}
	kTwo := 1

	fmt.Println(array_hashing.TopKElements(nums, k))
	fmt.Println(array_hashing.TopKElements(numsTwo, kTwo))
*/
