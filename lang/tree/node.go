package tree

import (
	"fmt"
)

//Node 节点
//go语言仅支持封装，不支持继承和多态
//go就没有class，只有struct
//面向接口
type Node struct {
	Value       int
	Left, Right *Node
}

//CreateNode 自定义工厂函数
func CreateNode(Value int) *Node {
	//局部变量的地址返回也是可以被调用者使用的，这和c不同
	return &Node{Value: Value}
}

//Print 接收者 receiver
func (node *Node) Print() {
	if node == nil {
		fmt.Print(" ")
		return
	}
	fmt.Print(node.Value, " ")
}

//SetValue 为结构定义方法
//只有使用指针才能改变结构内容
//nil指针也可以调用方法
func (node *Node) SetValue(Value int) {
	if node == nil {
		fmt.Println("setting Value to nil node.ignored!")
		return
	}
	node.Value = Value
}

//TraverseWithChannel 使用chan的遍历
func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
