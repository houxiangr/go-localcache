package ttypes

import "fmt"

type LinkedMap struct {
	head    *LinkedNode
	tail    *LinkedNode
	nodeMap map[string]*LinkedNode
}

type LinkedNode struct {
	key      string
	nextNode *LinkedNode
	preNode  *LinkedNode
}

func GetLinkMap() *LinkedMap {
	linkedMap := LinkedMap{}
	linkedMap.nodeMap = make(map[string]*LinkedNode)
	return &linkedMap
}

func (this *LinkedMap) SetTail(key string) {
	linkedNode := LinkedNode{
		key: key,
	}
	this.nodeMap[key] = &linkedNode

	if this.head == nil || this.tail == nil {
		this.head = &linkedNode
		this.tail = &linkedNode
	} else {
		this.tail.nextNode = &linkedNode
		linkedNode.preNode = this.tail
		this.tail = &linkedNode
	}
}

func (this *LinkedMap) DelHead() {
	if this.head == nil {
		return
	}

	delete(this.nodeMap, this.head.key)

	headnext := this.head.nextNode
	if headnext == nil {
		this.head = nil
		this.tail = nil
		return
	}
	this.head = headnext
}

func (this *LinkedMap) DelKey(key string) {

	node := this.nodeMap[key]

	if node == nil {
		return
	}

	if node == this.head && node == this.tail {
		this.head = nil
		this.tail = nil
	} else if node == this.head {
		this.head = this.head.nextNode
	} else if node == this.tail {
		this.tail = this.tail.preNode
	} else {
		node.preNode.nextNode = node.nextNode
		node.nextNode.preNode = node.preNode
	}

	node = nil
	delete(this.nodeMap, key)
}

func (this *LinkedMap) Len() int {
	return len(this.nodeMap)
}

func (this *LinkedMap) GetHead() *LinkedNode {
	return this.head
}

func (this LinkedNode) GetKey() string {
	return this.key
}

func (this LinkedMap) Range() {
	start := this.head
	for start != nil {
		fmt.Println(start.key)
		start = start.nextNode
	}
	fmt.Println("---------")
}
