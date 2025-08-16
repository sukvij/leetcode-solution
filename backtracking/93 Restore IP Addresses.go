package backtracking

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	ans := []string{}
	solveRestoreIpAddresses(s, 0, 0, 0, "", &ans)
	return ans
}

func solveRestoreIpAddresses(s string, from, to int, dots int, res string, ans *[]string) {
	if from == len(s) && to == len(s) {
		if dots == 4 {
			// fmt.Println()
			*ans = append(*ans, res[:len(res)-1])
		}
		return
	}

	if from == len(s) || to == len(s) {
		// fmt.Println(res)
		return
	}

	if dots == 4 {
		return
	}
	num, _ := strconv.Atoi(s[from : to+1])
	if num > 255 {
		return
	}

	x := strconv.Itoa(num)
	if len(x) == len(s[from:to+1]) {
		solveRestoreIpAddresses(s, to+1, to+1, dots+1, res+strconv.Itoa(num)+".", ans)
	}
	solveRestoreIpAddresses(s, from, to+1, dots, res, ans)
}

func RestoreIpAddresses() {
	fmt.Println(restoreIpAddresses("101023"))
}
