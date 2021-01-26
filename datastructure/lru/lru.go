package lru

type QNode struct {
	key, value int
	prev, next *QNode
}

type Queue struct {
	front, rear *QNode
}

type LRUCache struct {
	capacity, size int
	pageList       Queue
	pageMap        map[int]*QNode
}

func addQNode(key int, value int) *QNode {
	return &QNode{
		key:   key,
		value: value,
		prev:  nil,
		next:  nil,
	}
}

func (q *Queue) isEmpty() bool {
	return q.rear == nil
}

func (q *Queue) addFrontPage(key int, value int) *QNode {
	page := addQNode(key, value)
	if q.front == nil && q.rear == nil {
		q.front, q.rear = page, page
	} else {
		page.next = q.front
		q.front.prev = page
		q.front = page
	}
	return page
}

func (q *Queue) moveToFront(page *QNode) {
	if page == q.front {
		return
	} else if page == q.rear {
		q.rear = q.rear.prev
		q.rear.next = nil
	} else {
		page.prev.next = page.next
		page.next.prev = page.prev
	}

	page.next = q.front
	q.front.prev = page
	q.front = page
}

func (q *Queue) removeRear() {
	if q.isEmpty() {
		return
	} else if q.front == q.rear {
		q.front, q.rear = nil, nil
	} else {

	}
}

func (q *Queue) getRear() *QNode {
	return q.rear
}

func (lru *LRUCache) InitLru(capacity int) {
	lru.capacity = capacity
	lru.pageMap = make(map[int]*QNode)
}

func (lru *LRUCache) Get(key int) int {
	if _, found := lru.pageMap[key]; !found {
		return -1
	}

	val := lru.pageMap[key].value
	lru.pageList.moveToFront(lru.pageMap[key])
	return val
}

func (lru *LRUCache) Put(key int, value int) {
	if _, found := lru.pageMap[key]; found {
		// cover the value
		lru.pageMap[key].value = value
		lru.pageList.moveToFront(lru.pageMap[key])
		return
	}

	// not found
	if lru.size == lru.capacity {
		key := lru.pageList.getRear().key
		lru.pageList.removeRear()
		lru.size--
		delete(lru.pageMap, key)
	}

	page := lru.pageList.addFrontPage(key, value)
	lru.size++
	lru.pageMap[key] = page
}
