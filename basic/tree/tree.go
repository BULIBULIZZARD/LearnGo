package main

import (
	"./entry"
)

func main() {

	var root tree.Node
	root = Node{value:3}
	root.Left = &Node{}
	root.right = &Node{5,nil,nil}
	root.right.left = new (Node)
	root.left.right = createNode(2)
	root.right.left.setValue(4)
	root.traverse()

}
