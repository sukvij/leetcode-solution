package slidingwindow

import (
	"fmt"
	"sort"
)

func abs(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}

func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})
	ans := 100000000
	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]
			if abs(currentSum, target) < abs(ans, target) {
				ans = currentSum
			}
			if currentSum < target {
				left++
			} else if currentSum > target {
				right--
			} else {
				return currentSum
			}
		}
	}
	return ans
}

func ThreeSumClosest() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
