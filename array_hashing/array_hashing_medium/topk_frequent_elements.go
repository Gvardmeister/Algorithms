package array_hashing_medium

/*
	Lvl - medium

	Лучшая структура данных - это HashMap, Buckets(слайс слайсов)
	Краевые случаи: числа могут быть уникальными, одинаковыми, и пустой слайс
	Сложность: O(n) по времени, O(n+m) по памяти, но по факту O(n), так как m < n

	Дан массив целых чисел nums и целое число k, верните k наиболее частых элементов в массиве.
	Тестовые случаи сгенерированы таким образом, что ответ всегда является единственным.
	Вы можете вернуть результат в любом порядке.

	Input: nums = [1,2,2,3,3,3], k = 2
	Output: [2,3]

	Input: nums = [7,7], k = 1
	Output: [7]
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
