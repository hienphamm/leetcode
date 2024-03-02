package leetcode

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		if expectedNum, isExist := hashMap[target-num]; isExist {
			return []int{i, expectedNum}
		}
		hashMap[num] = i
	}

	return []int{}
}
