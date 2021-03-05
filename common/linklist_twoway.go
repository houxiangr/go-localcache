package common

type LinklistTwoWay struct {
	head *LinkNode
	tail *LinkNode
}

type LinkNode struct {
	value interface{}
	next  *LinkNode
	pre   *LinkNode
}

func (this LinklistTwoWay) GetHead()*LinkNode{
	return this.head

}

func (this LinklistTwoWay) GetTail()*LinkNode{
	return this.tail
}

func (this *LinklistTwoWay) SetHead(value interface{}){
	linkNode := LinkNode{
		value:value,
		pre:nil,
		next:this.head,
	}
	if this.head == nil {
		this.head = &linkNode
		this.tail = &linkNode
		return
	}
	this.head.pre = &linkNode
	this.head = &linkNode
}

func (this *LinklistTwoWay) SetTail(value interface{}){
	linkNode := LinkNode{
		value:value,
		pre:this.tail,
		next:nil,
	}
	if this.tail == nil {
		this.head = &linkNode
		this.tail = &linkNode
		return
	}
	this.tail.next = &linkNode
	this.tail = &linkNode
}

func (this *LinklistTwoWay) DelNode(targeNode *LinkNode) {
	preNode := targeNode.pre
	nextNode := targeNode.next

	if targeNode == this.head && targeNode == this.tail {
		this.head = nil
		this.tail = nil
	}else if targeNode == this.head {
		this.head = nextNode
		nextNode.pre = nil
	}else if targeNode == this.tail {
		this.tail = preNode
		preNode.next = nil
	}else{
		preNode.next = nextNode
		nextNode.pre = preNode
	}
	targeNode = nil
}

func (this *LinkNode)GetValue()interface{}{
	if this == nil {
		return nil
	}
	return this.value
}

func (this *LinkNode)GetNext()*LinkNode{
	if this == nil {
		return nil
	}
	return this.next
}

func (this *LinkNode)GetPre()*LinkNode{
	if this == nil {
		return nil
	}
	return this.pre
}
