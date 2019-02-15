package tree

import "fmt"

type Node struct {
	Value int
	Left , Right *Node
}

func (node Node) Print()  {
	fmt.Println(node.Value)
}
func (node *Node) SetValue(value int)  {
	node.Value= value
}
func (node *Node) Traverse(){
	node.TraverseFunc(func (node *Node){
		node.Print()
	})
	fmt.Println()
}
func (node *Node) TraverseFunc(f func (*Node)){
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
func CreateNode(value int)*Node{
	return &Node{Value:value}
}