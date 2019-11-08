package main

import (
	"fmt"
	"github.com/20100204/xuexi/struct/tree"
)

type MyNode struct {
	Node *tree.Node
}
type ab tree.Node

func (ab *ab) postOrder() {
	if ab == nil {
		return
	}
	ab.Left.Traverse()
	ab.Right.Traverse()
	ab.print()
}
func (ab *ab) print() {
	if ab == nil {
		return
	}
	fmt.Println(ab.Value)
}
func (myNode *MyNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	leftNode := &MyNode{myNode.Node.Left}
	leftNode.postOrder()
	rightNode := &MyNode{myNode.Node.Right}
	rightNode.postOrder()
	myNode.Node.Print()

}
func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	//root.Left.Right.SetValue(10)
	root.Left.Left.Print()
	//root.Traverse()
	root.SetValue(10)
	fmt.Println()
	myNode := MyNode{&root}
	myNode.postOrder()
	abc := ab(root)
	abc.postOrder()

}
