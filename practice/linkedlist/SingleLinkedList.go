package main

import "fmt"

type ListNode struct {
	next *ListNode
	data int
}

type LinkedList struct {
	head   *ListNode
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:   &ListNode{next: nil, data: 0},
		length: 0,
	}
}

func (linkedList *LinkedList) InsertNode(node *ListNode) bool {
	if node == nil {
		return false
	}
	tempListNode := linkedList.head
	for tempListNode != nil && tempListNode.next != nil {
		tempListNode = tempListNode.next
	}
	tempListNode.next = node
	linkedList.length = linkedList.length + 1
	return true
}

func (linkedList *LinkedList) DeleteNode(node *ListNode) bool {

	if node == nil {
		return false
	}
	tempListNode := linkedList.head
	for tempListNode.next != node {
		if tempListNode.next == nil {
			return false
		}
		tempListNode = tempListNode.next
	}
	tempListNode.next = tempListNode.next.next
	linkedList.length = linkedList.length - 1
	return true

}

func (linkedList *LinkedList) PrintLinkedList() {
	tempListNode := linkedList.head
	fmt.Println("LinkedList length = %d", linkedList.length)
	for tempListNode.next != nil {
		fmt.Println("node %v", tempListNode.next.data)
		tempListNode = tempListNode.next
	}
}

func main() {

	linkedList := NewLinkedList()
	var tempListNode *ListNode
	for i := 1; i < 6; i++ {
		tempListNode = &ListNode{next: nil, data: i}
		linkedList.InsertNode(tempListNode)
	}
	linkedList.PrintLinkedList()
	linkedList.DeleteNode(&ListNode{next: nil, data: 100})
	linkedList.PrintLinkedList()
}
