package stack

import (
	"fmt"
	"strings"
)

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	// fmt.Println(str)

	temp := []string{}
	num := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '+' || s[i] == '-' || s[i] == ')' {
			if num != "" {
				temp = append(temp, num)
				num = ""
			}
			temp = append(temp, string(s[i]))
			continue
		}
		num = num + string(s[i])
	}
	if num != "" {
		temp = append(temp, num)
	}
	fmt.Println(temp)

	ans := []string{}
	for i := 0; i < len(temp); i++ {
		if temp[i] == ")" {
			for {
				// index := len(ans) - 1
				// if
			}
		} else {
			ans = append(ans, temp[i])
		}
	}
	return 0
}

func Calculate() {
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}
