package main

import (
	"fmt"
	"iproute.org/learngo/lang/objs/tree1"
)

// MyTreeNode 扩展原来的tree1.Node
type MyTreeNode struct {
	node *tree1.Node
}

func (mn *MyTreeNode) dfs(action func(v int)) {
	if mn.node == nil {
		return
	}

	left := MyTreeNode{mn.node.GetLeft()}
	left.dfs(action)

	action(mn.node.GetValue())

	right := MyTreeNode{mn.node.GetRight()}
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
