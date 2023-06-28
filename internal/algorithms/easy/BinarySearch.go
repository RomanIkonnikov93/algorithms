package easy

/*
Дан массив целых чисел nums, отсортированный в порядке возрастания,
и целочисленная цель, напишите функцию для поиска цели в nums.
Если цель существует, верните ее индекс. В противном случае вернуть -1.
Вы должны написать алгоритм со сложностью выполнения O(log n).
*/

/*
Input: []int{-1,0,3,5,9,12}, 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Input: []int{-1,0,3,5,9,12}, 2
Output: -1
Explanation: 2 does not exist in nums so return -1
*/

func BinarySearch(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
