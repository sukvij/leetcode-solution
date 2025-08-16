package dp

import (
	"fmt"
	"sort"
)

func FindChanges(coins []int, amount int, index int, dp *[][]int) int {
	if amount == 0 {
		return 1
	}
	if amount < 0 {
		return 0
	}
	if index >= len(coins) {
		return 0
	}
	if (*dp)[index][amount] != -1 {
		return (*dp)[index][amount]
	}
	ans := FindChanges(coins, amount, index+1, dp)
	ans += FindChanges(coins, amount-coins[index], index, dp)
	(*dp)[index][amount] = ans
	return ans
}

func change(amount int, coins []int) int {
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] >= coins[j]
	})
	// fmt.Println(coins)
	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(coins)+1; i++ {
		dp[i] = make([]int, amount+1)
		for j := 0; j < amount+1; j++ {
			dp[i][j] = -1
		}
	}
	return FindChanges(coins, amount, 0, &dp)
}

func Change2() {
	fmt.Println(change(5, []int{1, 2, 5}))
}
