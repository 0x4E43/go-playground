package leetcode

import "strconv"

// Problem Link : https://leetcode.com/problems/two-sum/submissions/
func TwoSum(nums []int, target int) []int {
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

// Problem Link: https://leetcode.com/problems/palindrome-number/
func IsPalindrome(x int) bool {
	str := strconv.Itoa(x)
	if str[0] != str[len(str)-1] {
		return false
	}
	runes := []rune(str)
	// Reverse the order of the runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	if str == string(runes) {
		return true
	}
	return false
}
