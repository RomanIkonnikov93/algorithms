package easy

/*
Учитывая целочисленный массив nums, вернуть true,
если какое-либо значение встречается в массиве не менее двух раз,
и вернуть false, если каждый элемент различен.
*/

/*
Input: []int{1,2,3,1}
Output: true

Input: []int{1,2,3,4}
Output: false

Input: []int{1,1,1,3,3,4,3,2,4,2}
Output: true
*/

func ContainsDuplicateMap(nums []int) bool {

	m := make(map[int]int)

	for _, val := range nums {
		m[val] = val
	}

	if len(m) < len(nums) {
		return true
	} else {
		return false
	}
}

func ContainsDuplicateBruteForce(nums []int) bool {

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
