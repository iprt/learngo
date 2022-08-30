package main

import (
	"fmt"
)

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

/**
自定义的工厂函数

注意：返回的是局部变量的地址
*/
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

/**
初步面向对象

个人理解， who do(param) return

给结构体定义了方法？？？

func (接收者) 方法名(参数) 返回值

*/
func (node *treeNode) print() {
	if node == nil {
		fmt.Println("node is nil")
		return
	}
	fmt.Println("value is", node.value)
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("node is nil")
		return
	}
	node.value = value
}

func (node *treeNode) getValue() int {
	return node.value
}

/**
遍历
*/
func (node *treeNode) dfs(action func(node *treeNode)) {
	if node == nil {
		return
	}

	node.left.dfs(action)
	action(node)
	node.right.dfs(action)

}

func main() {
	/*
			2
		  /  \
		1      3
	*/

	root := treeNode{value: 3, left: nil, right: nil}

	fmt.Println(root)

	// node := createNode(3)
	// fmt.Println(node)
	//
	// root.left = node
	// fmt.Println(root)
	root.left = &treeNode{1, nil, nil}
	root.right = createNode(3)

	fmt.Println("hello world")

	fmt.Println("--- print")

	root.print()

	root.setValue(200)

	root.print()

	root.left.setValue(100)

	root.right.setValue(300)

	root.left.print()
	root.right.print()

	root.left.left.print()

	fmt.Println("dfs ... ")

	root.dfs(func(node *treeNode) {
		if node == nil {
			return
		}
		fmt.Println("node's value is:", node.getValue())
	})

	count := 0
	root.dfs(func(node *treeNode) {
		count++
	})

	fmt.Println("size is ", count)

}
