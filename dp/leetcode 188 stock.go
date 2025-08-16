package dp

import "fmt"

func maxProfit(k int, prices []int) int {
	n := len(prices)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k+2)
		for j := 0; j <= k+1; j++ {
			dp[i][j] = []int{-1000000, -1000000}
		}
	}
	return solve(prices, 0, 1, k, &dp)
}

func solve(prices []int, index int, turn int, total int, dp *[][][]int) int {
	if index >= len(prices) || total == 0 {
		return 0
	}

	if (*dp)[index][total][turn] != -1000000 {
		return (*dp)[index][total][turn]
	}

	ans := solve(prices, index+1, turn, total, dp)
	if turn == 1 {
		ans = max(ans, -prices[index]+solve(prices, index+1, 0, total, dp))
	} else {
		ans = max(ans, prices[index]+solve(prices, index+1, 1, total-1, dp))
	}
	(*dp)[index][total][turn] = ans
	return ans
}

func Maxprofile() {
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}
