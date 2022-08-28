package main

import (
	"fmt"
	"iproute.org/learngo/lang/objs/tree1"
)

func printNode(node *tree1.Node) {
	fmt.Println(node.GetValue())
}

func main() {
	root := tree1.Node{}
	root.SetValue(2)
	root.Print(printNode)

	root.Add(1)
	root.Add(3)

	root.DFS(printNode)
}
