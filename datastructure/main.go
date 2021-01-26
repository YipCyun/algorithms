package main

import (
	"fmt"
	"github.com/YipCyun/algorithms/datastructure/lru"
)

func main() {
	var cache lru.LRUCache
	cache.InitLru(2)
	cache.Put(2, 2)
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(1))

	cache.Put(1, 1)
	cache.Put(1, 5)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
	cache.Put(8, 8)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(8))
}
