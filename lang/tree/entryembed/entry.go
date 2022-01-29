package main

import (
	"fmt"

	"github.com/ghjan/learngo/lang/tree"
	"golang.org/x/tools/container/intsets"
)

//Embedding内嵌 可以直接访问Node下面的成员变量和方法
//就好象Node下面的成员平铺到了上一级
type myTreeNode struct {
	*tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse() {
	fmt.Println("This method is shadowed.")

}

func main() {
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	fmt.Println("\n- testNode tranverse 二叉树中序遍历----")
	testNode(root)
	fmt.Println("\n- myTreeNode postOrder 二叉树后序遍历----")
	root.postOrder()
	fmt.Println("\n- testSparser----")
	testSparse()

}

func testNode(root myTreeNode) {

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
	var pRoot2 = &myTreeNode{}
	pRoot2.SetValue(200)
	pRoot2.Print()
	pRoot2 = &root
	pRoot2.SetValue(300)
	pRoot2.Print()
	fmt.Println("\n-Traverse---")
	pRoot.SetValue(3)
	pRoot.Right.Left.SetValue(4)
	fmt.Print("pRoot.Traverse():")
	pRoot.Traverse()
	fmt.Print("pRoot.Node.Traverse():")
	pRoot.Node.Traverse()

	nodeCount := 0
	fmt.Println("\n-TraverseFunc， 自定义一个操作不再是print而已---")
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("nodeCount:", nodeCount)

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Printf("Max node value:%d\n", maxNode)
}

func testSparse() {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(100000))

}
