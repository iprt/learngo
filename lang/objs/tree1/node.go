package tree1

type Node struct {
	value       int
	left, right *Node
}

func CreateNode(value int) *Node {
	return &Node{value: value}
}

func (node *Node) GetLeft() *Node {
	return node.left
}

func (node *Node) GetRight() *Node {
	return node.right
}
func (node *Node) SetValue(v int) {
	node.value = v
}

func (node *Node) GetValue() int {
	return node.value
}

func (node *Node) Add(v int) *Node {
	if node == nil {
		return &Node{value: v}
	}
	if node.value == v {

		return node
	} else if v < node.value {
		node.left = node.left.Add(v)
		return node
	} else {
		node.right = node.right.Add(v)
		return node
	}
}

func (node *Node) Print(action func(*Node)) {
	action(node)
}

func (node *Node) DFS(action func(*Node)) {
	if nil == node {
		return
	}
	node.left.DFS(action)
	action(node)
	node.right.DFS(action)
}
