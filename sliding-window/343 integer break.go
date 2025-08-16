package slidingwindow

import "fmt"

func integerBreak(n int) int {
	dp := make([]int, n+3)
	dp[1] = 0
	dp[2] = 1
	dp[3] = 2

	for i := 4; i <= n; i++ {
		mxx := 0
		for j := 1; j < i; j++ {
			first, second := j, i-j
			x := max(dp[first], first)
			y := max(dp[second], second)
			mxx = max(mxx, x*y)
		}
		dp[i] = mxx
	}
	// fmt.Println(dp)
	return dp[n]
}

func IntegerBreak() {
	fmt.Println(integerBreak(10))
}
