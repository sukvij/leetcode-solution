package tree

import "fmt"

type Node struct {
	IsEnding bool
	Links    map[rune]*Node
}

type Trie struct {
	Root *Node
}

func NewNode() *Node {
	return &Node{IsEnding: false, Links: make(map[rune]*Node)}
}

func (trie *Trie) InsertWord(word string) {
	root := trie.Root
	for _, char := range word {
		if root.Links[char] == nil {
			root.Links[char] = NewNode()
		}
		root = root.Links[char]
	}
	root.IsEnding = true
}

func (trie *Trie) SearchWord(word string) bool {
	root := trie.Root
	for _, char := range word {
		if root.Links[char] == nil {
			return false
		}
	}
	return true
}

func findWords(board [][]byte, words []string) []string {
	// construct trie using words
	trie := &Trie{Root: NewNode()}
	for _, word := range words {
		trie.InsertWord(word)
	}

	n, m := len(board), len(board[0])
	mp := make(map[string]bool)

	root := trie.Root
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dfs(i, j, board, root, "", mp)
		}
	}

	ans := []string{}
	for key, _ := range mp {
		ans = append(ans, key)
	}
	return ans
}

func dfs(i, j int, board [][]byte, root *Node, word string, mp map[string]bool) {
	// if root.IsEnding {
	// 	fmt.Println("haha ", word)
	// }
	n, m := len(board), len(board[0])
	if i < 0 || j < 0 || i >= n || j >= m {
		return
	}
	char := board[i][j]
	if char == '@' {
		return
	}
	nextWord := word + string(char)
	nextRoot := root.Links[rune(char)]
	if nextRoot == nil {
		return
	}
	if nextRoot.IsEnding {
		mp[nextWord] = true
		// fmt.Println("founded ", nextWord)
	}
	board[i][j] = '@'

	dfs(i+1, j, board, nextRoot, word+string(char), mp)
	dfs(i-1, j, board, nextRoot, word+string(char), mp)
	dfs(i, j+1, board, nextRoot, word+string(char), mp)
	dfs(i, j-1, board, nextRoot, word+string(char), mp)

	board[i][j] = char
}

func FindWords() {
	boards := [][]byte{}
	boards = append(boards, []byte{'o', 'a', 'a', 'n'})
	boards = append(boards, []byte{'e', 't', 'a', 'e'})
	boards = append(boards, []byte{'i', 'h', 'k', 'r'})
	boards = append(boards, []byte{'i', 'f', 'l', 'v'})

	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(boards, words))
}
