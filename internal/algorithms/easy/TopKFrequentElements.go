package easy

/*
Если задан целочисленный массив nums и целое число k, то верните k наиболее часто встречающихся элементов.
Ответ может быть получен в любом порядке.
*/

/*
Input: []int{1,1,1,2,2,3}, k = 2
Output: []{1,2}
Example 2:

Input: nums = []int{1}, k = 1
Output: [1]
*/

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	res := []int{}
	for k > 0 {
		max := 0
		val := 0
		for k, v := range m {
			if v > max {
				max = v
				val = k
			}
		}
		res = append(res, val)
		k--
		delete(m, val)
	}

	return res
}
