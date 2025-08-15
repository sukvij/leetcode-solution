package slidingwindow

import (
	"fmt"
)

func minWindow(s string, t string) string {
	lower, upper := make([]int, 26), make([]int, 26)
	for _, val := range t {
		if val >= 'a' && val <= 'z' {
			lower[int(val-'a')]++
		} else if val >= 'A' && val <= 'Z' {
			upper[int(val-'A')]++
		}
	}

	compLower, compUpper := make([]int, 26), make([]int, 26)
	left, length, minLeft, minRight := -1, 1000000, -1, -1

	for index, val := range s {
		if val >= 'a' && val <= 'z' {
			compLower[int(val-'a')]++
		} else if val >= 'A' && val <= 'Z' {
			compUpper[int(val-'A')]++
		}

		// we will compare both strings
		for {
			founded := true
			for i := 0; i < 26; i++ {
				if lower[i] > compLower[i] || upper[i] > compUpper[i] {
					founded = false
					break
				}
			}
			if !founded {
				break
			}

			// mean founded bro
			left++

			// fmt.Println(left, index, s[left:index+1])
			if length > index-left+1 {
				length = index - left + 1
				minLeft = left
				minRight = index
			}
			x := s[left]
			if x >= 'a' && x <= 'z' {
				compLower[int(x-'a')]--
			} else if x >= 'A' && x <= 'Z' {
				compUpper[int(x-'A')]--
			}
		}
	}
	if length >= 1000000 {
		return ""
	}
	return s[minLeft : minRight+1]
}

func MinWindow() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
