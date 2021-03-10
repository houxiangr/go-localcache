package ttypes

import (
	"reflect"
	"testing"
)

var linklist LinklistTwoWay

func initLinklist() {
	linklist = LinklistTwoWay{}
	linklist.SetHead("5", 5)
	linklist.SetHead("4", 4)
	linklist.SetHead("3", 3)
	linklist.SetHead("2", 2)
	linklist.SetHead("1", 1)
}

func TestGetHeadTail(t *testing.T) {
	initLinklist()
	tests := []struct {
		name     string
		wantHead interface{}
		wantTail interface{}
	}{
		{
			name:     "get head and tail value",
			wantHead: 1,
			wantTail: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := linklist.GetHead()
			if !reflect.DeepEqual(head.GetValue(), tt.wantHead) {
				t.Errorf("GetHead() = %v, want %v", head.GetValue(), tt.wantHead)
			}
			tail := linklist.GetTail()
			if !reflect.DeepEqual(tail.GetValue(), tt.wantTail) {
				t.Errorf("GetTail() = %v, want %v", tail.GetValue(), tt.wantTail)
			}
		})
	}
}

func TestGetHeadTail2(t *testing.T) {
	emptyLinklist := LinklistTwoWay{}
	tests := []struct {
		name     string
		wantHead *LinkNode
		wantTail *LinkNode
	}{
		{
			name:     "get empty linklist",
			wantHead: nil,
			wantTail: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := emptyLinklist.GetHead()
			if !reflect.DeepEqual(head, tt.wantHead) {
				t.Errorf("GetHead() = %v, want %v", head, tt.wantHead)
			}
			tail := emptyLinklist.GetTail()
			if !reflect.DeepEqual(tail, tt.wantTail) {
				t.Errorf("GetTail() = %v, want %v", tail, tt.wantTail)
			}
		})
	}
}

func TestSetHeadTail(t *testing.T) {
	initLinklist()
	tests := []struct {
		name     string
		wantHead interface{}
		wantTail interface{}
	}{
		{
			name:     "set head and tail value",
			wantHead: 0,
			wantTail: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linklist.SetHead("0", 0)
			linklist.SetTail("6", 6)
			head := linklist.GetHead()
			if !reflect.DeepEqual(head.GetValue(), tt.wantHead) {
				t.Errorf("GetHead() = %v, want %v", head.GetValue(), tt.wantHead)
			}
			tail := linklist.GetTail()
			if !reflect.DeepEqual(tail.GetValue(), tt.wantTail) {
				t.Errorf("GetTail() = %v, want %v", tail.GetValue(), tt.wantTail)
			}
		})
	}
}

func TestDelNode(t *testing.T) {
	initLinklist()
	tests := []struct {
		name     string
		delNode  *LinkNode
		wantNext *LinkNode
		wantHead *LinkNode
		wantTail *LinkNode
	}{
		//1-3-4-5
		{
			name:     "del middle node",
			delNode:  linklist.GetHead().GetNext(),
			wantNext: linklist.GetHead().GetNext().GetNext(),
			wantHead: linklist.GetHead(),
			wantTail: linklist.GetTail(),
		},
		//3-4-5
		{
			name:     "del head node",
			delNode:  linklist.GetHead(),
			wantNext: linklist.GetHead().GetNext().GetNext().GetNext(),
			wantHead: linklist.GetHead().GetNext().GetNext(),
			wantTail: linklist.GetTail(),
		},
		//3-4
		{
			name:     "del tail node",
			delNode:  linklist.GetTail(),
			wantNext: linklist.GetHead().GetNext().GetNext().GetNext(),
			wantHead: linklist.GetHead().GetNext().GetNext(),
			wantTail: linklist.GetTail().GetPre(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linklist.DelNode(tt.delNode)
			node := linklist.GetHead().GetNext()
			if !reflect.DeepEqual(node, tt.wantNext) {
				t.Errorf("GetNext() = %v, want %v", node.GetValue(), tt.wantNext.GetValue())
			}
			head := linklist.GetHead()
			if !reflect.DeepEqual(head, tt.wantHead) {
				t.Errorf("GetHead() = %v, want %v", head.GetValue(), tt.wantHead.GetValue())
			}
			tail := linklist.GetTail()
			if !reflect.DeepEqual(tail, tt.wantTail) {
				t.Errorf("GetTail() = %v, want %v", tail.GetValue(), tt.wantTail.GetValue())
			}
		})
	}
}

func TestMoveNodeToHead(t *testing.T) {
	initLinklist()
	tests := []struct {
		name       string
		wantHead   *LinkNode
		wantNext   *LinkNode
		moveTarget *LinkNode
	}{
		{
			name:       "get head and tail value",
			wantHead:   linklist.GetHead().GetNext(),
			wantNext:   linklist.GetHead(),
			moveTarget: linklist.GetHead().GetNext(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linklist.MoveNodeToHead(tt.moveTarget)
			head := linklist.GetHead()
			if !reflect.DeepEqual(head, tt.wantHead) {
				t.Errorf("GetHead() = %v, want %v", head.GetValue(), tt.wantHead.GetValue())
			}
			next := linklist.GetHead().GetNext()
			if !reflect.DeepEqual(next, tt.wantNext) {
				t.Errorf("GetNext() = %v, want %v", next.GetValue(), tt.wantNext.GetValue())
			}
		})
	}
}

func TestLinklistTwoWay_DelTailAndSetHead(t *testing.T) {
	initLinklist()
	tests := []struct {
		name     string
		wantHead interface{}
		wantTail interface{}
	}{
		{
			//0-1-2-3-4
			name:     "del tail and set head",
			wantHead: 0,
			wantTail: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linklist.DelTailAndSetHead("0", 0)
			head := linklist.GetHead()
			if !reflect.DeepEqual(head.GetValue(), tt.wantHead) {
				t.Errorf("GetHead() = %v, want %v", head.GetValue(), tt.wantHead)
			}
			tail := linklist.GetTail()
			if !reflect.DeepEqual(tail.GetValue(), tt.wantTail) {
				t.Errorf("GetTail() = %v, want %v", tail.GetValue(), tt.wantTail)
			}
		})
	}
}
