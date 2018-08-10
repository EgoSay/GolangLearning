package main

import (
	"Learning/tree"
	"fmt"
)

func main() {
	root := tree.Node{Value: 3}
	//root.Print()
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Left.Right = new(tree.Node)
	root.Right.Left = tree.CreateNode(2)
	root.Left.Right.SetValue(4)
	//root.Left.Right.Print()
	root.PostOrder()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println(nodeCount)
}
