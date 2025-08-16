package slidingwindow

import "fmt"

func removeDuplicates(nums []int) int {
	left := -100000000
	n, count := len(nums), 0

	for i := 0; i < n; {
		if nums[i] == left {
			nums[i] = 100000000
			count++
			i++
		} else {
			left = nums[i]
			if i == n-1 {
				break
			}
			if nums[i+1] != nums[i] {
				i++
			} else {
				i += 2
			}

		}
	}
	fmt.Println(nums)
	left = 0
	right := 0
	for left < n && right < n {
		if nums[left] != 100000000 {
			left++
			right++
		} else {
			// find right
			for right < n {
				if nums[right] == 100000000 {
					right++
				} else {
					break
				}
			}
			// found left and right index
			// swap till
			for left < right && right < n {
				nums[left], nums[right] = nums[right], nums[left]
				left++
				right++
				if right < n && nums[right] == 100000000 {
					for right < n {
						if nums[right] != 100000000 {
							break
						}
						right++
					}
				}
			}
		}
	}
	fmt.Println(nums)
	return len(nums) - count
}

func RemoveDuplicates() {
	fmt.Println(removeDuplicates([]int{1, 2, 2, 2, 3, 4, 4, 4, 5, 6, 6, 6}))
}
