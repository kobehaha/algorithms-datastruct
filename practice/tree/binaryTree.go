package main

type Node struct {
	data  int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

type Stack struct {
	data  []int
	count int
	lenth int
}

// preOder
// preOder(r) = print r->preOder(r->left)->preOder(r->right)
func (tree *Tree) PreOrderRecursive(node *Node) bool {

	if node == nil {
		return false
	}

	fmt.Println(node.data)
	if node.left != nil {
		tree.PreOrderRecursive(node.left)
	}
	if node.right != nil {
		tree.PreOrderRecursive(node.right)
	}
	return true

}

func (tree *Tree) PreOderStack(node *Node) bool {
	if node == nil {
		return nil
	}
	tempNode := node
	// todo: change size
	stack := &Stack{data: make([]int, 100), count: 0, lenth: 100}
	for tempNode != nil || stack.length() != 0 {
		for tempNode != nil {
			fmt.Println(tempNode.data)
			stack.Push(tempNode)
			tempNode = tempNode.left
		}
		tempNode = stack.Pop()
		tempNode = tempNode.right
	}

	return true

}

// midOrder
// midOrder = midOrder(r-left) print r midOrder(r->left)
func (tree *Tree) MidOrderRecursive(node *Node) {

	if node == nil {
		return false
	}

	if node.left != nil {
		tree.MidOrderRecursive(node.left)
	}

	fmt.Println(node.data)

	if node.right != nil {
		tree.MidOrderRecursive(node.right)
	}
	return true

}

func (tree *Tree) MidOrderStack(node *Node) bool {
	if node == nil {
		return false
	}
	stack := &Stack{data: make([]int, 100), count: 0, lenth: 100}
	tempNode := node
	for tempNode != nil || statck.length() != 0 {
		for tempNode != nil {
			stack.Push(tempNode)
			tempNode = tempNode.left
		}
		tempNode = stack.Pop()
		fmt.Println(tempNode.data)
		tempNode = tempNode.right
	}
	return true
}

func (tree *Tree) PostOrderRecursive(node *Node) bool {
	if node == nil {
		return false
	}
	tree.PostOrderRecursive(node.left)
	tree.PostOrderRecursive(node.right)
	fmt.Println(node.data)
	return true
}

func (tree *Tree) PostOrderStack(node *Node) bool {
	if node == nil {
		return false
	}
	last := nil
	tempNode := node
	stack := &Stack{data: make([]int, 100), count: 0, lenth: 100}
	for tempNode != nil || statck.length() != 0 {
		for tempNode != nil {
			stack.Push(tempNode)
			tempNode := tempNode.left
		}

		last = stack.Peek()
		if tempNode.right == nil || tempNode.right == last {
			fmt.Println(tempNode.node)
			stack.Pop()
			last = tempNode
			tempNode = nil
		} else {
			tempNode = tempNode.right
		}

	}

	return true

}
