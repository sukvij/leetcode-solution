package dp

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	mp := make(map[string]bool)
	for _, word := range wordDict {
		mp[word] = true
	}
	n := len(s)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	ans := wordBreakOneByOne(s, mp, 0, 0, &dp)
	return ans > 0
}

func wordBreakOneByOne(s string, mp map[string]bool, from, to int, dp *[][]int) int {

	if from == len(s) && to == len(s) {
		return 1
	}
	if from >= len(s) || to >= len(s) {
		return 0
	}
	if (*dp)[from][to] != -1 {
		return (*dp)[from][to]
	}
	ans := 0
	ans = ans + wordBreakOneByOne(s, mp, from, to+1, dp)
	if ans > 0 {
		return ans
	}
	if _, ok := mp[s[from:to+1]]; ok {
		ans = ans + wordBreakOneByOne(s, mp, to+1, to+1, dp)
	}
	(*dp)[from][to] = ans
	return ans
}

func WordBreak() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}
