package two_pointers_medium

import "sort"

/*
	Лучшая структура это - slice
	Краевые случаи: массив не может быть пустым
	Сложность: O(n^2) по времени, O(1) по памяти (пространственная сложность)

	Задан целочисленный массив nums.
	Верните все тройки [nums[i], nums[j], nums[k]],
	где nums[i] + nums[j] + nums[k] == 0, а индексы i, j и k все различны.

	Вывод не должен содержать повторяющихся троек.
	Вы можете вернуть вывод и тройки в любом порядке.

	Input: nums = [-1,0,1,2,-1,-4]
    Output: [[-1,-1,2],[-1,0,1]]

    Input: nums = [0,1,1]
    Output: []

    Input: nums = [0,0,0]
    Output: [[0,0,0]]
*/

func threeSum(nums []int) [][]int {
	var result [][]int // альтернатива result := [][]int{}

	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
