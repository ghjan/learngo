package main

import (
	"fmt"

	"github.com/ghjan/learngo/tree"
	"golang.org/x/tools/container/intsets"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	fmt.Println("\n- testNode tranverse 二叉树中序遍历----")
	testNode(root)
	fmt.Println("\n- myTreeNode postOrder 二叉树后序遍历----")
	mynode := myTreeNode{&root}
	mynode.postOrder()
	fmt.Println("\n- testSparser----")
	testSparse()

}

func testNode(root tree.Node) {

	fmt.Println(root)
	//nodes := []Node{
	//  {value: 3},
	//  {},
	//  {6, nil, &root},
	//}
	//fmt.Println(nodes)

	root.Print()
	fmt.Println()
	root.Right.Left.Print()
	fmt.Println()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	root.SetValue(100)
	root.Print()
	fmt.Println()
	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	fmt.Println("\n-nil---")
	var pRoot2 *tree.Node
	pRoot2.SetValue(200)
	pRoot2.Print()
	pRoot2 = &root
	pRoot2.SetValue(300)
	pRoot2.Print()
	fmt.Println("\n-Traverse---")
	pRoot.SetValue(3)
	pRoot.Right.Left.SetValue(4)
	pRoot.Traverse()

	nodeCount := 0
	fmt.Println("\n-TraverseFunc， 自定义一个操作不再是print而已---")
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("nodeCount:", nodeCount)
}

func testSparse() {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(100000))

}
