package main

import "fmt"

type Node struct {
	next *Node
	data int
}

type LinkedList struct {
	head   *Node
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:   &Node{next: nil, data: 0},
		length: 0,
	}
}

func (linkedList *LinkedList) InsertNode(node *Node) bool {
	if node == nil {
		return false
	}
	tempNode := linkedList.head
	for tempNode != nil && tempNode.next != nil {
		tempNode = tempNode.next
	}
	tempNode.next = node
	linkedList.length = linkedList.length + 1
	return true
}

func (linkedList *LinkedList) DeleteNode(node *Node) bool {

	if node == nil {
		return false
	}
	tempNode := linkedList.head
	for tempNode.next != node {
		if tempNode.next == nil {
			return false
		}
		tempNode = tempNode.next
	}
	tempNode.next = tempNode.next.next
	linkedList.length = linkedList.length - 1
	return true

}

func (linkedList *LinkedList) PrintLinkedList() {
	tempNode := linkedList.head
	fmt.Println("LinkedList length = %d", linkedList.length)
	for tempNode.next != nil {
		fmt.Println("node %v", tempNode.next.data)
		tempNode = tempNode.next
	}
}

func main() {

	linkedList := NewLinkedList()
	var tempNode *Node
	for i := 1; i < 6; i++ {
		tempNode = &Node{next: nil, data: i}
		linkedList.InsertNode(tempNode)
	}
	linkedList.PrintLinkedList()
	linkedList.DeleteNode(&Node{next: nil, data: 100})
	linkedList.PrintLinkedList()
}
