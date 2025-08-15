package array

func GroupAnagrams() [][]string {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ans := [][]string{}
	mp := make(map[string][]string)

	for _, str := range strs {
		temp := ""
		arr := make([]int, 26)
		for _, char := range str {
			arr[int(char-'a')]++
		}
		for key, idx := range arr {
			for i := 0; i < idx; i++ {
				temp = temp + string(rune(key)+'a')
			}
		}
		mp[temp] = append(mp[temp], str)
	}
	// fmt.Println(mp)
	for _, val := range mp {
		ans = append(ans, val)
	}
	return ans
}
