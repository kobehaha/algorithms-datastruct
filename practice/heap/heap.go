package  main


import "fmt"

type Heap struct {

  data []int
  count int
  capacity int
}

func NewHeap(capacity int) *Heap {
  return &Heap{
    data : make([]int, capacity+1),
    count : 0,
    capacity: capacity,
  }
}

// from down ---> top, !!  save data from index 1 not 0 !!!!!!!!!!!!!!  
func (heap *Heap)Insert(node int) bool{
  if heap.count == heap.capacity {
    return false 
  }

  heap.count = heap.count + 1
  heap.data[heap.count] = node

  // swith values
  i := heap.count
  parent := heap.count / 2

  fmt.Println("i value= %d",i)
  fmt.Println("parent value = %d", parent)
  for parent > 0 && heap.data[parent] < heap.data[i]{
    swap(heap, parent, i)
    i = parent
    parent =  i / 2
  }

  return true

}

// from top ---> down
func (heap *Heap)RemoveMax() bool{
  if heap.count == 0 {
    return false
  }

  swap(heap, 1, heap.count)
  i := 1 
  for i < heap.count/2 {
    maxIndex := i 
    if heap.data[i] < heap.data[2 * i]  {
      maxIndex = 2* i
    }
    if (heap.data[i] <= heap.data[2 * i + 1]) {
      maxIndex = 2*i + 1 
    }
    if (maxIndex == i ){
      break
    }
    swap(heap, i, maxIndex)
    i = maxIndex

  }
  heap.count = heap.count - 1
  return true
}

func swap(heap *Heap, i int, j int){
  temp := heap.data[i] 
  heap.data[i] = heap.data[j]
  heap.data[j] = temp
}


func main(){

  fmt.Println("heap test")
  heap := NewHeap(30)
  heap.Insert(10)
  heap.Insert(70)
  heap.Insert(100)
  heap.Insert(110)
  heap.Insert(130)
  heap.Insert(112)
  heap.Insert(120)
  heap.Insert(140)
  heap.Insert(131)
  heap.RemoveMax()
  fmt.Println(heap.data)

}
