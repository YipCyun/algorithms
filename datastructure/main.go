package main

import (
	"fmt"
	"github.com/YipCyun/algorithms/datastructure/reverselinkedlist"
)

func main() {
	/*
		lru
	*/
	//var cache lru.LRUCache
	//cache.InitLru(2)
	//cache.Put(2, 2)
	//fmt.Println(cache.Get(2))
	//fmt.Println(cache.Get(1))
	//
	//cache.Put(1, 1)
	//cache.Put(1, 5)
	//fmt.Println(cache.Get(1))
	//fmt.Println(cache.Get(2))
	//cache.Put(8, 8)
	//fmt.Println(cache.Get(1))
	//fmt.Println(cache.Get(8))

	/*
		reverse linked list
	*/
	initialValues := []int{0, 1, 2, 3}
	headNode := reverselinkedlist.InitLinkedList(initialValues)

	fmt.Println(reverselinkedlist.List2Integers(headNode))
	headNode.Reverse()
}
