package ttypes

import (
	"reflect"
	"testing"
)

var linkedMap *LinkedMap

func initLinkedMap() {
	linkedMap = GetLinkMap()
}

func TestLinkedMap_SetTail(t *testing.T) {
	initLinkedMap()
	tests := []struct {
		name     string
		key      string
		wantHead interface{}
		wantTail interface{}
	}{
		{
			name:     "set tail 1",
			key:      "1",
			wantHead: "1",
			wantTail: "1",
		},
		{
			name:     "set tail 2",
			key:      "2",
			wantHead: "1",
			wantTail: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linkedMap.SetTail(tt.key)
			head := linkedMap.GetHead()
			if !reflect.DeepEqual(head.GetKey(), tt.wantHead) {
				t.Errorf("head.GetKey() = %v, key %v", head.GetKey(), tt.wantHead)
			}
			tail := linkedMap.GetTail()
			if !reflect.DeepEqual(tail.GetKey(), tt.wantTail) {
				t.Errorf("tail.GetKey() = %v, want %v", tail.GetKey(), tt.wantTail)
			}
		})
	}
}

func TestLinkedMap_DelHead(t *testing.T) {
	initLinkedMap()
	linkedMap.SetTail("1")
	linkedMap.SetTail("2")
	tests := []struct {
		name     string
		wantHead interface{}
	}{
		{
			name:     "del head 1",
			wantHead: "2",
		},
		{
			name:     "del head 2",
			wantHead: nil,
		},
		{
			name:     "del head nil",
			wantHead: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linkedMap.DelHead()
			head := linkedMap.GetHead()
			if head != nil && !reflect.DeepEqual(head.GetKey(), tt.wantHead) {
				t.Errorf("head.GetKey() = %v, key %v", head.GetKey(), tt.wantHead)
			}
		})
	}
}

func Test_DelKey(t *testing.T) {
	initLinkedMap()
	linkedMap.SetTail("1")
	linkedMap.SetTail("2")
	linkedMap.SetTail("3")
	linkedMap.SetTail("4")
	linkedMap.SetTail("5")
	tests := []struct {
		name     string
		delKey   string
		wantHead interface{}
		wantTail interface{}
	}{
		{
			name:     "del other",
			delKey:   "3",
			wantHead: "1",
			wantTail: "5",
		},
		{
			name:     "del head",
			delKey:   "1",
			wantHead: "2",
			wantTail: "5",
		},
		{
			name:     "del tail",
			delKey:   "5",
			wantHead: "2",
			wantTail: "4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linkedMap.DelKey(tt.delKey)
			head := linkedMap.GetHead()
			if !reflect.DeepEqual(head.GetKey(), tt.wantHead) {
				t.Errorf("head.GetKey() = %v, key %v", head.GetKey(), tt.wantHead)
			}
			tail := linkedMap.GetTail()
			if !reflect.DeepEqual(tail.GetKey(), tt.wantTail) {
				t.Errorf("tail.GetKey() = %v, want %v", tail.GetKey(), tt.wantTail)
			}
		})
	}
}
