package easy

/*
Учитывая массив целых чисел nums и целочисленное целевое значение,
верните индексы двух чисел так, чтобы в сумме они составляли целевое значение.
Вы можете предположить, что каждый вход будет иметь ровно одно решение,
и вы не можете использовать один и тот же элемент дважды.
Вы можете вернуть ответ в любом порядке.
*/

/*
Input: []int{2,7,11,15}, 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Input: []int{3,2,4}, 6
Output: [1,2]

Input: []int{3,3}, 6
Output: [0,1]
*/

func TwoSumMap(nums []int, target int) []int {

	m := make(map[int]int)

	for i, val := range nums {
		if res, ok := m[target-val]; ok {
			return []int{res, i}
		}
		m[val] = i
	}

	return []int{}
}

func TwoSumBruteForce(nums []int, target int) []int {

	arr := []int{0, 1}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				arr[0] = i
				arr[1] = j
			}
		}
	}

	return arr
}
