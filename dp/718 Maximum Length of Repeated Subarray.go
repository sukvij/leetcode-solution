package dp

import "fmt"

func findLength(nums1 []int, nums2 []int) int {
	n1 := len(nums1)
	n2 := len(nums2)
	dp := make([][][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([][]int, n2+1)
		for j := 0; j <= n2; j++ {
			dp[i][j] = []int{-1, -1}
		}
	}
	return longestCommonPrefix(nums1, nums2, len(nums1)-1, len(nums2)-1, 0, &dp)
}

func longestCommonPrefix(nums1, nums2 []int, n1, n2 int, included int, dp *[][][]int) int {
	if n1 < 0 || n2 < 0 {
		return 0
	}
	// three choices
	if (*dp)[n1][n2][included] != -1 {
		return (*dp)[n1][n2][included]
	}
	ans := 0
	if included == 0 {
		ans = max(ans, longestCommonPrefix(nums1, nums2, n1, n2-1, 0, dp))
		ans = max(ans, longestCommonPrefix(nums1, nums2, n1-1, n2, 0, dp))
		ans = max(ans, longestCommonPrefix(nums1, nums2, n1-1, n2-1, 0, dp))
		if nums1[n1] == nums2[n2] {
			ans = max(ans, 1+longestCommonPrefix(nums1, nums2, n1-1, n2-1, 1, dp))
		}
	} else {
		if nums1[n1] == nums2[n2] {
			ans = max(ans, 1+longestCommonPrefix(nums1, nums2, n1-1, n2-1, 0, dp))
		} else {
			(*dp)[n1][n2][included] = 0
			return 0
		}
	}
	(*dp)[n1][n2][included] = ans
	return ans
}

func FindLength() {
	fmt.Println(findLength([]int{0, 1, 1, 1, 1}, []int{1, 0, 1, 0, 1}))
}
