package stack

import "fmt"

type Cell struct {
	Index, Val int
}

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	stk := []Cell{}
	stk = append(stk, Cell{Index: 0, Val: heights[0]})
	n := len(heights)
	ans := heights[0]
	for i := 1; i < n; i++ {
		if heights[i] >= stk[len(stk)-1].Val {
			stk = append(stk, Cell{Index: i, Val: heights[i]})
			// ans = max(ans, stk[0].Val*(i-stk[0].Index+1))
		} else {
			// we need to pop things
			for len(stk) > 0 {
				top := stk[len(stk)-1]
				if top.Val <= heights[i] {
					break
				}
				width := (i - top.Index)
				ans = max(ans, width*top.Val)
				stk = stk[:len(stk)-1]
			}
			stk = append(stk, Cell{Index: i, Val: heights[i]})
		}
		fmt.Println(ans, "after ", heights[i])
	}
	fmt.Println(stk)
	if len(stk) == 1 {
		ans = max(ans, stk[0].Val)
	}
	if len(stk) > 1 {
		ans = max(ans, stk[0].Val*(stk[len(stk)-1].Index-stk[0].Index+1))
	}
	return ans
}

func LargestRectangleArea() {
	fmt.Println(largestRectangleArea([]int{0, 9}))
}

// 2,1,5,6,2,3
// 5 6
