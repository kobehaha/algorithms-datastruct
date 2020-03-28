package main

import "fmt"

type CircleListNode struct {
	next *CircleListNode
	data int
}

type CircleLinkedList struct {
	head   *CircleListNode
    tail   *CircleListNode
	length int
}

func NewCircleLinkedList() *CircleLinkedList {
    head := &CircleListNode{next: nil, data:0}
    tail := &CircleListNode{next: nil , data: 0}
    head.next = tail
    tail.next = head
	return &CircleLinkedList{
		head:   head,
        tail:   tail,
		length: 0,
	}
}

func (circleLinkedList *CircleLinkedList) InsertNode(node *CircleListNode) bool {
	if node == nil {
		return false
	}
	tempCircleListNode := circleLinkedList.head
	for tempCircleListNode.next != circleLinkedList.tail  {
		tempCircleListNode = tempCircleListNode.next
	}
    tempCircleListNode.next = node
    node.next = circleLinkedList.tail
	circleLinkedList.length = circleLinkedList.length + 1
	return true
}

func (circleLinkedList *CircleLinkedList) DeleteNode(node *CircleListNode) bool {

	if node == nil {
		return false
	}
	tempCircleListNode := circleLinkedList.head
	for tempCircleListNode.next != circleLinkedList.tail {
        if (tempCircleListNode.next == node){
	        tempCircleListNode.next = node.next
	        circleLinkedList.length = circleLinkedList.length - 1
            return true
        }
		tempCircleListNode = tempCircleListNode.next
	}
    return false

}

func (circleLinkedList *CircleLinkedList) PrintLinkedList() {
	tempCircleListNode := circleLinkedList.head
	fmt.Println("LinkedList length = %d", circleLinkedList.length)
	for tempCircleListNode.next != circleLinkedList.tail {
		fmt.Println("node %v", tempCircleListNode.next.data)
		tempCircleListNode = tempCircleListNode.next
	}
}

func main() {

	circleLinkedList := NewCircleLinkedList()
	var tempCircleListNode *CircleListNode
	for i := 1; i < 6; i++ {
		tempCircleListNode = &CircleListNode{next: nil, data: i}
		circleLinkedList.InsertNode(tempCircleListNode)
	}
	circleLinkedList.PrintLinkedList()
	circleLinkedList.DeleteNode(tempCircleListNode)
	circleLinkedList.DeleteNode(&CircleListNode{next: nil, data: 100})
	circleLinkedList.PrintLinkedList()
}
