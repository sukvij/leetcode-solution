package array

import "fmt"

func subarraySumPositiveElements(nums []int, k int) int {
	currSum := nums[0]
	left, right := 0, 0
	for left <= right {
		if currSum == k {
			fmt.Println(left, right)
			currSum -= nums[left]
			left++
		} else if currSum < k {
			right++
			if right >= len(nums) {
				break
			}
			currSum += nums[right]
		} else {
			currSum -= nums[left]
			left++
		}
	}
	return 0
}

func SubSum() {
	subarraySumNegativeElements([]int{1, 1, 1}, 2)
}

func subarraySumNegativeElements(arr []int, target int) {
	mp := make(map[int]int)
	currSum := 0
	mp[currSum] = 0
	for currIndex, val := range arr {
		currSum += val
		extra := currSum - target
		if index, ok := mp[extra]; ok {
			fmt.Println("from to to ", index+1, currIndex)
		}
		mp[currSum] = currIndex
	}
}
