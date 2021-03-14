package ttypes

type SmallHeap struct {
	heapSlice []*TimeLimitValue
}

func (this *SmallHeap) Push(value *TimeLimitValue) {
	this.heapSlice = append(this.heapSlice, value)
}

func (this *SmallHeap) Adjust() {
	size := len(this.heapSlice)
	//从根节点开始整理堆
	for i := 0; i < size/2; i++ {
		left := getLeft(i)
		right := getRight(i)
		minNode := this.getMinNode(left, right)

		if this.heapSlice[minNode].ExpireTime < this.heapSlice[i].ExpireTime {
			this.heapSlice[minNode], this.heapSlice[i] = this.heapSlice[i], this.heapSlice[minNode]
		}
	}
}

func (this *SmallHeap) Rise(n int) {
	for {
		parent := getParent(n)
		if this.heapSlice[parent].ExpireTime > this.heapSlice[n].ExpireTime {
			this.heapSlice[parent], this.heapSlice[n] = this.heapSlice[n], this.heapSlice[parent]
			n = parent
		} else {
			break
		}

		if n == 0 {
			break
		}
	}
}

func (this *SmallHeap) Sink(n int) {
	size := len(this.heapSlice)
	if size == 0{
		return
	}
	for {
		left := getLeft(n)
		right := getRight(n)
		minNode := this.getMinNode(left, right)
		if this.heapSlice[n].ExpireTime > this.heapSlice[minNode].ExpireTime {
			this.heapSlice[minNode], this.heapSlice[n] = this.heapSlice[n], this.heapSlice[minNode]
			n = minNode
		} else {
			break
		}

		if n >= size/2 {
			break
		}
	}
}

func (this *SmallHeap) DelRoot() {
	size := len(this.heapSlice)
	this.heapSlice[0], this.heapSlice[size-1] = this.heapSlice[size-1], this.heapSlice[0]
	this.heapSlice = this.heapSlice[:size-1]
	//本来只有一个元素，删除后就空了
	if size == 1 {
		return
	}
	this.Sink(0)
}

func (this *SmallHeap) GetRoot() *TimeLimitValue {
	if len(this.heapSlice) == 0 {
		return nil
	}
	return this.heapSlice[0]
}

func getLeft(n int) int {
	return n*2 + 1
}

func getRight(n int) int {
	return n*2 + 2
}

func getParent(n int) int {
	return n / 2
}

func (this *SmallHeap) getMinNode(left, right int) int {
	size := len(this.heapSlice)
	var leftNode *TimeLimitValue
	var rightNode *TimeLimitValue
	if left < size {
		leftNode = this.heapSlice[left]
	}
	if right < size {
		rightNode = this.heapSlice[right]
	}
	if rightNode == nil {
		return left
	}
	if leftNode.ExpireTime <= rightNode.ExpireTime {
		return left
	}
	return right
}
