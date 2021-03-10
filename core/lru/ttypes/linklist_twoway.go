package ttypes

import (
	"sync/atomic"
	"unsafe"
)

type LinklistTwoWay struct {
	head *LinkNode
	tail *LinkNode
}

type LinkNode struct {
	key   string
	value interface{}
	next  *LinkNode
	pre   *LinkNode
}

func (this LinklistTwoWay) GetHead() *LinkNode {
	return this.head

}

func (this LinklistTwoWay) GetTail() *LinkNode {
	return this.tail
}

func (this *LinklistTwoWay) SetHead(key string, value interface{}) *LinkNode {
	return this.setHead(key, value)
}

func (this *LinklistTwoWay) SetTail(key string, value interface{}) {
	linkNode := LinkNode{
		key:   key,
		value: value,
		pre:   this.tail,
		next:  nil,
	}

	if this.tail == nil {
		this.head = &linkNode
		this.tail = &linkNode
		return
	}
	this.tail.next = &linkNode
	this.tail = &linkNode
}

func (this *LinklistTwoWay) DelNode(targetNode *LinkNode) {
	if targetNode == this.head {
		this.notLinkHead()
	} else if targetNode == this.tail {
		this.notLinkTail()
	} else {
		this.notLinkNode(targetNode)
	}
	targetNode = nil
}

func (this *LinklistTwoWay) MoveNodeToHead(targetNode *LinkNode) {
	if targetNode == this.head {
		return
	}
	//not link target node
	if targetNode == this.head {
		this.notLinkHead()
	} else if targetNode == this.tail {
		this.notLinkTail()
	} else {
		this.notLinkNode(targetNode)
	}

	//set target node to head
	targetNode.next = this.head
	targetNode.pre = nil
	if this.head != nil {
		this.head.pre = targetNode
	}

	this.head = targetNode
	//when cache only one ele
	if this.tail == nil {
		this.tail = targetNode
	}
}

func (this *LinklistTwoWay) DelTailAndSetHead(key string, value interface{}) *LinkNode {
	this.notLinkTail()
	newNode := this.setHead(key, value)
	return newNode
}

func (this *LinklistTwoWay) setHead(key string, value interface{}) *LinkNode {
	linkNode := LinkNode{
		key:   key,
		value: value,
		pre:   nil,
		next:  this.head,
	}

	if this.head == nil {
		this.head = &linkNode
		this.tail = &linkNode
		return &linkNode
	}
	this.head.pre = &linkNode
	this.head = &linkNode
	return &linkNode
}

func (this *LinklistTwoWay) notLinkNode(targetNode *LinkNode) {
	preNode := targetNode.pre
	nextNode := targetNode.next
	preNode.next = nextNode
	nextNode.pre = preNode
}

func (this *LinklistTwoWay) notLinkHead() {
	if this.tail == this.head {
		this.head = nil
		this.tail = nil
		return
	}
	nextHead := this.head.next
	this.head = nextHead
	nextHead.pre = nil
}

func (this *LinklistTwoWay) notLinkTail() {
	if this.tail == this.head {
		this.head = nil
		this.tail = nil
		return
	}
	nextTail := this.tail.pre
	this.tail = nextTail
	nextTail.next = nil
}

func (this *LinklistTwoWay) LinklistToSlice() []LinkNode {
	node := this.GetHead()
	res := []LinkNode{}
	for node != nil {
		res = append(res, *node)
		node = node.GetNext()
	}
	return res
}

func (this *LinkNode) GetValue() interface{} {
	if this == nil {
		return nil
	}
	return this.value
}

func (this *LinkNode) GetNext() *LinkNode {
	if this == nil {
		return nil
	}
	return this.next
}

func (this *LinkNode) TrySetValue(newValue interface{}) {
	for {
		ptrThisValue := &this.value
		thisValue := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&ptrThisValue)))
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&ptrThisValue)), thisValue, unsafe.Pointer(&newValue)) {
			return
		}
	}
}
func (this *LinkNode) SetValue(newValue interface{}) {
	for {
		ptrThisValue := &this.value
		thisValue := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&ptrThisValue)))
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&ptrThisValue)), thisValue, unsafe.Pointer(&newValue)) {
			return
		}
	}
}

func (this *LinkNode) GetPre() *LinkNode {
	if this == nil {
		return nil
	}
	return this.pre
}

func (this *LinkNode) GetKey() string {
	if this == nil {
		return ""
	}
	return this.key
}
