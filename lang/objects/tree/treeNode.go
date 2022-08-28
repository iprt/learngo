package main

import (
	"fmt"
)

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	root := treeNode{}

	fmt.Println(root)

	node := createNode(3)
	fmt.Println(node)

	root.left = node

	fmt.Println(root)

}
