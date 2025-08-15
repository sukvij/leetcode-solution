package array

import "fmt"

// it will print numbers also

func LongestConsecutive() int {
	nums := []int{100, 4, 200, 1, 3, 2}
	mp := make(map[int]bool)

	for _, val := range nums {
		mp[val] = true
	}

	for _, val := range nums {
		if canUse := mp[val]; !canUse {
			continue
		}
		ans := []int{}
		ans = append(ans, val)
		mp[val] = false

		left := val - 1
		for {
			if canUse, ok := mp[left]; !ok || !canUse {
				break
			}
			mp[left] = false
			ans = append(ans, left)
			left--
		}
		right := val + 1
		for {
			if canUse, ok := mp[right]; !ok || !canUse {
				break
			}
			mp[right] = false
			ans = append(ans, right)
			right++
		}
		fmt.Println(ans)
	}
	return 0
}
