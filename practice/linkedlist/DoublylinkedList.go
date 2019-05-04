package main

import "fmt"

type ListNode struct {
    pre  *ListNode
	next *ListNode
	data int
}

type DoublyLinkedList struct {
	head   *ListNode
	length int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head:   &ListNode{pre:nil, next: nil, data: 0},
		length: 0,
	}
}

func (doublyLinkedList *DoublyLinkedList) InsertNode(node *ListNode) bool {
	if node == nil {
		return false
	}
	tempListNode := doublyLinkedList.head
	for tempListNode != nil && tempListNode.next != nil {
		tempListNode = tempListNode.next
	}
	tempListNode.next = node
    node.pre = tempListNode
	doublyLinkedList.length = doublyLinkedList.length + 1
	return true
}

func (doublyLinkedList *DoublyLinkedList) GetPreNode(node *ListNode) *ListNode {
   if node == nil {
     return nil 
   }
   return node.pre


}

func (doublyLinkedList *DoublyLinkedList) DeleteNode(node *ListNode) bool {

	if node == nil {
		return false
	}
	tempListNode := doublyLinkedList.head
	for tempListNode.next != node {
		if tempListNode.next == nil {
			return false
		}

		tempListNode = tempListNode.next
	}

    if tempListNode.next.next == nil {
      tempListNode.next = nil
	  doublyLinkedList.length = doublyLinkedList.length - 1
      return true
    }


	tempListNode.next = tempListNode.next.next
    tempListNode.next.next.pre = tempListNode.next
	doublyLinkedList.length = doublyLinkedList.length - 1
	return true

}


func (doublyLinkedList *DoublyLinkedList) PrintLinkedList() {
	tempListNode := doublyLinkedList.head
	fmt.Println("LinkedList length = %d", doublyLinkedList.length)
	for tempListNode.next != nil {
		fmt.Println("node %v", tempListNode.next.data)
		tempListNode = tempListNode.next
	}
}

func main() {

	doublyLinkedList := NewDoublyLinkedList()
	var tempListNode *ListNode
	for i := 1; i < 6; i++ {
		tempListNode = &ListNode{next: nil, pre: nil, data: i}
		doublyLinkedList.InsertNode(tempListNode)
	}
	doublyLinkedList.PrintLinkedList()
	doublyLinkedList.DeleteNode(tempListNode)
	doublyLinkedList.DeleteNode(tempListNode.pre)
	doublyLinkedList.DeleteNode(&ListNode{next: nil, data: 100})
	doublyLinkedList.PrintLinkedList()
}
