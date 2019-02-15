package tree

import "fmt"

type Node struct {
	Value int
	Left , Right *Node
}

func (node Node) Print()  {
	fmt.Println(node.value)
}
func (node *Node) SetValue(value int)  {
	node.Value= value
}
func (node *Node) Traverse(){
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func CreateNode(value int )*Node{
	return &Node{Value:value}
}