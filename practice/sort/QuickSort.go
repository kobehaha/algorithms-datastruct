package main


import "fmt"

// quickSort(p,r) = quickSort(p..q-1) + quick(q+1 , r) 
// if p >= r
func QuickSort(arr []int, n int){
   if len(arr) > n {
     return
   }
   quickSort(arr, 0, n-1)
}

func partition(arr []int, start, end int) int{
  q := arr[end] 

  i := start 

  for j:= start; j < end; j++{
    if arr[j] < q {
      if !(i==j){
        arr[i], arr[j] = arr[j], arr[i]
      }
      i++
    }
  }

  arr[i], arr[end] = arr[end], arr[i]

  return i 
}

func quickSort(arr []int, start, end int){

  if start >= end {
    return
  }
  q := partition(arr, start, end) 
  quickSort(arr, start, q-1)
  quickSort(arr, q+1, end)
}

func main(){
  arr := []int{1,2,9,3,2}
  QuickSort(arr, 5)
  fmt.Println(arr)
}
