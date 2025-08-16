package stack

import "fmt"

// type Cell struct {
// 	FromUpper, FromLeft int
// }

func maximalRectangle(matrix [][]byte) int {
	// n, m := len(matrix), len(matrix[0])
	// dp := make([][]Cell, n)
	// for i := 0; i < n; i++ {
	// 	dp[i] = make([]Cell, m)
	// }
	// fmt.Println(dp)
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < m; j++ {
	// 		if matrix[i][j] == '0' {
	// 			dp[i][j].FromLeft = 0
	// 			dp[i][j].FromUpper = 0
	// 			continue
	// 		}
	// 		if i == 0 && j == 0 {
	// 			dp[i][j].FromLeft = 1
	// 			dp[i][j].FromUpper = 1
	// 			continue
	// 		}
	// 		if i == 0 {
	// 			dp[i][j].FromLeft = max(1, 1+dp[i][j-1].FromLeft)
	// 			dp[i][j].FromUpper = 1
	// 			// dp[i][j] = Cell{FromUpper: 1, FromLeft: max(1, dp[i][j-1].FromLeft)}
	// 		} else if j == 0 {
	// 			dp[i][j].FromLeft = 1
	// 			dp[i][j].FromUpper = max(1, 1+dp[i-1][j].FromUpper)
	// 			// dp[i][j] = Cell{FromUpper: , FromLeft: 1}
	// 		} else {

	// 			maxUpper, maxLeft := max(1, 1+dp[i-1][j].FromUpper), max(1, 1+dp[i][j-1].FromLeft)
	// 			dp[i][j].FromLeft = maxLeft
	// 			dp[i][j].FromUpper = maxUpper
	// 			// dp[i][j] = Cell{FromUpper: maxUpper, FromLeft: maxLeft}
	// 		}
	// 	}
	// }
	// fmt.Println(dp)
	return 0
}

func MaximalRectangle() {
	fmt.Println(maximalRectangle(nil))
}

// 1,1  0,0  1,1  0,0  0,0
// 2,1  0,0  2,1  1,2  1,3
// 3,1  1,2  3,3  2,4  2,5
// {4 1} {0 0} {0 0} {3 1} {0 0}
