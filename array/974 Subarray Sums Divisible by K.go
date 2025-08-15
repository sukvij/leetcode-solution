package array

import "fmt"

func subarraysDivByK(nums []int, k int) {
	mp := make(map[int][]int)
	currSum := 0
	mp[0] = append(mp[0], -1)

	for index, val := range nums {
		currSum += (val + k*10000)
		currSum = currSum % k
		extra := currSum % k

		if lastIndex, ok := mp[extra]; ok {
			for _, haha := range lastIndex {
				fmt.Println(haha+1, index)
			}
		}

		mp[extra] = append(mp[extra], index)
	}
}

func SubarrayDivByK() {
	subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5)
}
