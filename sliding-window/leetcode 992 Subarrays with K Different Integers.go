package slidingwindow

import "fmt"

// func subarraysWithKDistinct(nums []int, k int) int {
// 	mp := make(map[int]int)
// 	ans := 0
// 	left := 0
// 	for index, val := range nums {
// 		mp[val]++
// 		if len(mp) > k {
// 			left = makeNumsProper(nums, left, index, k, mp)
// 		}
// 		if len(mp) == k {
// 			ans += findAllSolutions(nums, left, index, k)
// 		}
// 	}
// 	return ans
// }

// func makeNumsProper(nums []int, left, index, k int, mp map[int]int) int {
// 	for i := left; i < index; i++ {
// 		val := nums[i]
// 		mp[val]--
// 		if mp[val] == 0 {
// 			delete(mp, val)
// 		}
// 		if len(mp) == k {
// 			return i + 1
// 		}
// 	}
// 	return 0
// }

// func findAllSolutions(nums []int, low, high int, k int) int {
// 	total, till := 0, -1
// 	mp := make(map[int]int)
// 	for i := high; i >= low; i++ {
// 		val := mp[nums[i]]
// 		mp[val]++
// 		if len(mp) == k {
// 			till = i
// 			break
// 		}
// 	}
// 	// mean ans will be from low to till
// 	total += till - low + 1
// 	return total
// }

func SubarraysWithKDistinct() {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2))
}

func subarraysWithKDistinct(nums []int, k int) int {
	return countSubarraysWithMostKDistinct(nums, k) - countSubarraysWithMostKDistinct(nums, k-1)
}

func countSubarraysWithMostKDistinct(nums []int, k int) int {
	if k < 1 {
		return 0
	}

	freq := make(map[int]int)
	left := 0
	res := 0
	for right := 0; right < len(nums); right++ {
		val := nums[right]
		freq[val]++
		for len(freq) > k {
			temp := nums[left]
			freq[temp]--
			if freq[temp] == 0 {
				delete(freq, temp)
			}
			left++
		}
		res += right - left + 1
	}
	return res
}
