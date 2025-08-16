package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
	Root *TreeNode
}

func Constructor() Codec {
	return Codec{Root: nil}
}

func solveSerial(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	ans := "["
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			ans += "null,"
			return
		}
		ans += strconv.Itoa(node.Val) + ","
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)

	ans = ans[:len(ans)-1] + "]"
	return ans
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// this.Root = root

	return solveSerial(root)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	// fmt.Println("data is ", data[1:len(data)-1])
	// str := strings.Split(data[1:len(data)-1], ",")
	// fmt.Println("after spliting ", str, data, len(str), str[0])
	// if len(str) == 0 || str == nil {
	// 	return nil
	// }
	if data == "[]" {
		return nil
	}
	return makeTree(strings.Split(data[1:len(data)-1], ","))
}

func makeTree(str []string) *TreeNode {
	// fmt.Println("str is ", str)
	if len(str) == 0 {
		// fmt.Println("zero length")
		return nil
	}
	index := 0
	var buildTree func() *TreeNode
	buildTree = func() *TreeNode {
		if index >= len(str) {
			return nil
		}
		if str[index] == "null" {
			index++
			return nil
		}
		val, _ := strconv.Atoi(str[index])
		node := &TreeNode{Val: val}
		index++
		node.Left = buildTree()
		node.Right = buildTree()
		return node
	}
	return buildTree()
}

func SerDeser() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}
	serial := Constructor()
	// fmt.Println(serial.serialize(root))
	// node := serial.deserialize(serial.serialize(root))
	str := serial.serialize(nil)
	fmt.Println(str)
	node := serial.deserialize(str)
	fmt.Println(serial.serialize(node))
}
