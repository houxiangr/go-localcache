package ttypes

type LinkedMap struct {
	head *LinkedNode
	tail *LinkedNode
	nodeMap map[string]*LinkedNode
}

type LinkedNode struct {
	key string
	nextNode *LinkedNode
	preNode *LinkedNode
}

func GetLinkMap() LinkedMap{
	linkedMap := LinkedMap{}
	linkedMap.nodeMap = make(map[string]*LinkedNode)
	return linkedMap
}

func (this *LinkedMap)SetTail(key string){
	linkedNode := LinkedNode{
		key:key,
	}
	this.nodeMap[key] = &linkedNode

	if this.head == nil || this.tail == nil {
		this.head = &linkedNode
		this.tail = &linkedNode
	}else{
		this.tail.nextNode = &linkedNode
		this.tail = linkedNode.nextNode
	}
}

func(this *LinkedMap)DelHead(){
	if this.head == nil {
		return
	}

	delete(this.nodeMap,this.head.key)

	headnext := this.head.nextNode
	if headnext == nil {
		this.head = nil
		this.tail = nil
		return
	}
	this.head = headnext
}
