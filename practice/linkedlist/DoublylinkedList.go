package main

import "fmt"

type DoubleLinkedNode struct {
    pre  *DoubleLinkedNode
	next *DoubleLinkedNode
	data int
}

type DoublyLinkedList struct {
	head   *DoubleLinkedNode
	length int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head:   &DoubleLinkedNode{pre:nil, next: nil, data: 0},
		length: 0,
	}
}

func (doublyLinkedList *DoublyLinkedList) InsertNode(node *DoubleLinkedNode) bool {
	if node == nil {
		return false
	}
	tempDoubleLinkedNode := doublyLinkedList.head
	for tempDoubleLinkedNode != nil && tempDoubleLinkedNode.next != nil {
		tempDoubleLinkedNode = tempDoubleLinkedNode.next
	}
	tempDoubleLinkedNode.next = node
    node.pre = tempDoubleLinkedNode
	doublyLinkedList.length = doublyLinkedList.length + 1
	return true
}

func (doublyLinkedList *DoublyLinkedList) GetPreNode(node *DoubleLinkedNode) *DoubleLinkedNode {
   if node == nil {
     return nil 
   }
   return node.pre


}

func (doublyLinkedList *DoublyLinkedList) DeleteNode(node *DoubleLinkedNode) bool {

	if node == nil {
		return false
	}
	tempDoubleLinkedNode := doublyLinkedList.head
	for tempDoubleLinkedNode.next != node {
		if tempDoubleLinkedNode.next == nil {
			return false
		}

		tempDoubleLinkedNode = tempDoubleLinkedNode.next
	}

    if tempDoubleLinkedNode.next.next == nil {
      tempDoubleLinkedNode.next = nil
	  doublyLinkedList.length = doublyLinkedList.length - 1
      return true
    }


	tempDoubleLinkedNode.next = tempDoubleLinkedNode.next.next
    tempDoubleLinkedNode.next.next.pre = tempDoubleLinkedNode.next
	doublyLinkedList.length = doublyLinkedList.length - 1
	return true

}


func (doublyLinkedList *DoublyLinkedList) PrintLinkedList() {
	tempDoubleLinkedNode := doublyLinkedList.head
	fmt.Println("LinkedList length = %d", doublyLinkedList.length)
	for tempDoubleLinkedNode.next != nil {
		fmt.Println("node %v", tempDoubleLinkedNode.next.data)
		tempDoubleLinkedNode = tempDoubleLinkedNode.next
	}
}

func main() {

	doublyLinkedList := NewDoublyLinkedList()
	var tempDoubleLinkedNode *DoubleLinkedNode
	for i := 1; i < 6; i++ {
		tempDoubleLinkedNode = &DoubleLinkedNode{next: nil, pre: nil, data: i}
		doublyLinkedList.InsertNode(tempDoubleLinkedNode)
	}
	doublyLinkedList.PrintLinkedList()
	doublyLinkedList.DeleteNode(tempDoubleLinkedNode)
	doublyLinkedList.DeleteNode(tempDoubleLinkedNode.pre)
	doublyLinkedList.DeleteNode(&DoubleLinkedNode{next: nil, data: 100})
	doublyLinkedList.PrintLinkedList()
}
