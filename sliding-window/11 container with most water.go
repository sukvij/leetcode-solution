package slidingwindow

import "fmt"

func maxArea(height []int) int {

	left, right := 0, len(height)-1
	ans := 0
	for left < right {
		ans = max(ans, (right-left)*min(height[left], height[right]))
		if height[left] >= height[right] {
			right--
		} else {
			left++
		}
	}
	return ans
}

func MaxArea() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
