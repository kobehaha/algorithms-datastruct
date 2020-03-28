package main

import (
	"fmt"
)


type LruCache struct {

	head *DataNode
	count int
	size int
}



type DataNode struct {
	k string
	v string
	next *DataNode
}

func NewLruCacheNode(size int) *LruCache {

	return &LruCache{
		head : &DataNode{k:"", v:"", next: nil },
		count : 0,
		size: size,
	}
}

func (cache *LruCache) Get(k string) (bool, string){

	tmpNode := cache.head
	for tmpNode.next != nil{
		if tmpNode.k == k{
			return true, tmpNode.v
		}
		tmpNode = tmpNode.next
	}
	return false, ""
}

func (cache *LruCache) Push(k, v string) bool {

	if cache.count == cache.size {
		//fill
		tmpNode := cache.head
		for tmpNode.next != nil && tmpNode.next.next != nil {
			tmpNode = tmpNode.next
			if tmpNode.k == k {
				newNode := &DataNode{k:k , v:v, next: cache.head}
				cache.head = newNode
				tmpNode.next = tmpNode.next.next
				return true
			}
		}
		tmpNode.next = nil
		newNode := &DataNode{k:k , v:v, next: cache.head}
		cache.head = newNode
		return true
	}else{
		newNode := &DataNode{k:k , v:v, next: cache.head}
		cache.head = newNode
		return false
	}
	return false
}

func main(){

	newLru := NewLruCacheNode(100000)
	newLru.Push("k", "k-v")
	newLru.Push("k2", "k-v2")
	newLru.Push("k3", "k-v3")
	newLru.Push("k4", "k-v4")
	newLru.Push("k5", "k-v5")
	newLru.Push("k6", "k-v6")
	newLru.Push("k7", "k-v7")
	newLru.Push("k8", "k-v8")
	_, value := newLru.Get("k")
	_, value2 := newLru.Get("k8")
	fmt.Println("value %+s", value)
	fmt.Println("value2 %+s", value2)

}