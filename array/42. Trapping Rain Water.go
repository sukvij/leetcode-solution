package array

import "fmt"

func Trap(height []int) int {
	n := len(height)
	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	ans := 0
	for i := 1; i < n-1; i++ {
		l, r := leftMax[i-1], rightMax[i+1]
		minOfBoth := min(l, r)
		if minOfBoth > height[i] {
			ans += minOfBoth - height[i]
		}
		fmt.Println(ans)
	}
	return ans
}
