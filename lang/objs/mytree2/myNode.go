package main

import (
	"fmt"
	"iproute.org/learngo/lang/objs/tree1"
)

// MyTreeNode 扩展原来的tree1.Node
type MyTreeNode struct {
	// 使用内嵌的方式来扩展
	*tree1.Node
}

func (mn *MyTreeNode) dfs(action func(v int)) {
	// 这句话很重要！！！
	if mn == nil || mn.Node == nil {
		return
	}

	left := MyTreeNode{mn.GetLeft()}
	left.dfs(action)

	action(mn.GetValue())

	right := MyTreeNode{mn.GetRight()}
	right.dfs(action)

}

func main() {

	root := tree1.Node{}
	root.SetValue(2)
	root.Add(1)
	root.Add(3)

	mn := MyTreeNode{&root}
	mn.dfs(func(v int) {
		fmt.Println("value = ", v)
	})

}
