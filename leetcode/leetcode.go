package leetcode

func twoSum(nums []int, target int) []int {
	var ret = make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]+nums[j] == target && i != j {
				ret[0] = i
				ret[1] = j
				return ret
			}
		}
	}
	return ret
}
