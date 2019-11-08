package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node *Node) Print() {
	if node == nil {
		fmt.Println("Setting value to nil node.Ignored")
		return
	}
	fmt.Println(node.Value)
}
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node.Ignored.")
		return
	}
	node.Value = value
}
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
