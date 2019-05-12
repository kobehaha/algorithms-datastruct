package main

import "fmt"

type PriorityQueue struct {
	data     []Node
	capacity int
	count    int
}

type Node struct {
	value    int
	priority int
}

func NewPriorityQueue(capacity int) *PriorityQueue {
	return &PriorityQueue{
		data:     make([]Node, capacity+1, capacity+1),
		capacity: capacity,
		count:    0,
	}
}

func (this *PriorityQueue) Insert(node *Node) bool {
	if this.capacity == this.count {
		return false
	}
	this.data[count] = node
	this.count = this.count + 1
	// adjust Heap use by Pop()
	return true
}

func (this *PriorityQueue) Pop() *Node {
	if this.count == 0 {
		return &Node{-1, -1}
	}

	adjust(this.data, 1, this.count)
	node := this.data[1]
	this.data[1] = this.data[count]
	this.count = this.count - 1
	return node

}

func adjust(data []Node, start, end int) {
   if start > end {
     return false
   }
   // just keep most high priority value on queue top 
   for i := end/2; start < i; i-- {
     high := i
     if data[high].priortiy < data[high*2].prioprity {
       high = 2i
     }  
     if data[high].priority < data[high*2 + 1] {
       high = 2i + 1 
     } 
     if high == i {
       continue
     }
     data[high], data[i] = data[i], data[high]

   }
}
