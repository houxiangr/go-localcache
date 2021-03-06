package common

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

func (this *LinklistTwoWay) SetHead(key string, value interface{}) {
	linkNode := LinkNode{
		key:   key,
		value: value,
		pre:   nil,
		next:  this.head,
	}
	if this.head == nil {
		this.head = &linkNode
		this.tail = &linkNode
		return
	}
	this.head.pre = &linkNode
	this.head = &linkNode
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
	this.NotLinkNode(targetNode)
	targetNode = nil
}

func (this *LinklistTwoWay)NotLinkNode(targetNode *LinkNode){
	preNode := targetNode.pre
	nextNode := targetNode.next

	if targetNode == this.head && targetNode == this.tail {
		this.head = nil
		this.tail = nil
	} else if targetNode == this.head {
		this.head = nextNode
		nextNode.pre = nil
	} else if targetNode == this.tail {
		this.tail = preNode
		preNode.next = nil
	} else {
		preNode.next = nextNode
		nextNode.pre = preNode
	}
}

func (this *LinklistTwoWay)MoveNodeToHead(targetNode *LinkNode){
	if targetNode == this.head {
		return
	}
	//not link target node
	this.NotLinkNode(targetNode)

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

func (this *LinklistTwoWay)LinklistToSlice()[]LinkNode{
	node := this.GetHead()
	res := []LinkNode{}
	for node != nil {
		res = append(res, *node)
		node = node.GetNext()
	}
	return res
}

func NewLinkNode(key string, value interface{}) *LinkNode {
	return &LinkNode{
		key:   key,
		value: value,
	}
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
