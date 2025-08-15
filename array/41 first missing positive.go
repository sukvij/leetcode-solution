package array

import (
	"math"
)

func firstMissingPositive(nums []int) int {
	var smallPositive, largestPositive int64 = math.MaxInt64, math.MinInt64
	var total int64 = 0
	for _, val := range nums {
		if val > 0 {
			total += int64(val)
			smallPositive = min(smallPositive, int64(val))
			largestPositive = max(largestPositive, int64(val))
		}
	}

	if total == 0 {
		return -1
	}
	elements := (largestPositive - smallPositive) + 1
	var targetSum int64 = elements*(smallPositive-1) + (elements*(elements+1))/2
	// fmt.Println(targetSum-total, total, targetSum, elements)
	return int(targetSum - total)
}

func FirstMissingPositive() {
	firstMissingPositive([]int{3, 4, -1, 1})
}
