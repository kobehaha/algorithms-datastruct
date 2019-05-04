package main

import "fmt"

type ListNode struct {
	next *ListNode
	data int
}

type CircleLinkedList struct {
	head   *ListNode
    tail   *ListNode
	length int
}

func NewCircleLinkedList() *CircleLinkedList {
    head := &ListNode{next: nil, data:0}
    tail := &ListNode{next: nil , data: 0}
    head.next = tail
    tail.next = head
	return &CircleLinkedList{
		head:   head,
        tail:   tail,
		length: 0,
	}
}

func (circleLinkedList *CircleLinkedList) InsertNode(node *ListNode) bool {
	if node == nil {
		return false
	}
	tempListNode := circleLinkedList.head
	for tempListNode.next != circleLinkedList.tail  {
		tempListNode = tempListNode.next
	}
    tempListNode.next = node
    node.next = circleLinkedList.tail
	circleLinkedList.length = circleLinkedList.length + 1
	return true
}

func (circleLinkedList *CircleLinkedList) DeleteNode(node *ListNode) bool {

	if node == nil {
		return false
	}
	tempListNode := circleLinkedList.head
	for tempListNode.next != circleLinkedList.tail {
        if (tempListNode.next == node){
	        tempListNode.next = node.next
	        circleLinkedList.length = circleLinkedList.length - 1
            return true
        }
		tempListNode = tempListNode.next
	}
    return false

}

func (circleLinkedList *CircleLinkedList) PrintLinkedList() {
	tempListNode := circleLinkedList.head
	fmt.Println("LinkedList length = %d", circleLinkedList.length)
	for tempListNode.next != circleLinkedList.tail {
		fmt.Println("node %v", tempListNode.next.data)
		tempListNode = tempListNode.next
	}
}

func main() {

	circleLinkedList := NewCircleLinkedList()
	var tempListNode *ListNode
	for i := 1; i < 6; i++ {
		tempListNode = &ListNode{next: nil, data: i}
		circleLinkedList.InsertNode(tempListNode)
	}
	circleLinkedList.PrintLinkedList()
	circleLinkedList.DeleteNode(tempListNode)
	circleLinkedList.DeleteNode(&ListNode{next: nil, data: 100})
	circleLinkedList.PrintLinkedList()
}
