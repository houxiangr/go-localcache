package ttypes

import (
	"reflect"
	"testing"
)

var smallHeap SmallHeap

func initSmallHeap() {
	smallHeap = SmallHeap{}
	smallHeap.Push(&TimeLimitValue{
		Value:      6,
		ExpireTime: 6,
	})
	smallHeap.Push(&TimeLimitValue{
		Value:      1,
		ExpireTime: 1,
	})
	smallHeap.Push(&TimeLimitValue{
		Value:      3,
		ExpireTime: 3,
	})
	smallHeap.Push(&TimeLimitValue{
		Value:      5,
		ExpireTime: 5,
	})
	smallHeap.Push(&TimeLimitValue{
		Value:      2,
		ExpireTime: 2,
	})
	smallHeap.Push(&TimeLimitValue{
		Value:      4,
		ExpireTime: 4,
	})

}

func TestSmallHeap_Push(t *testing.T) {
	initSmallHeap()
	tests := []struct {
		name      string
		pushvalue *TimeLimitValue
		want      int
	}{
		{
			name: "push first value",
			pushvalue: &TimeLimitValue{
				Value:      1,
				ExpireTime: 1,
			},
			want: 7,
		},
		{
			name: "push second value",
			pushvalue: &TimeLimitValue{
				Value:      2,
				ExpireTime: 2,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smallHeap.Push(tt.pushvalue)
			if !reflect.DeepEqual(len(smallHeap.heapSlice), tt.want) {
				t.Errorf("smallHeap.Push() = %v, want %v", len(smallHeap.heapSlice), tt.want)
			}
		})
	}
}

func TestSmallHeap_Adjust(t *testing.T) {
	initSmallHeap()
	tests := []struct {
		name string
		want []*TimeLimitValue
	}{
		{
			name: "adjust small heap",
			want: []*TimeLimitValue{
				&TimeLimitValue{
					Value:      1,
					ExpireTime: 1,
				},
				&TimeLimitValue{
					Value:      2,
					ExpireTime: 2,
				},
				&TimeLimitValue{
					Value:      3,
					ExpireTime: 3,
				},
				&TimeLimitValue{
					Value:      5,
					ExpireTime: 5,
				},
				&TimeLimitValue{
					Value:      6,
					ExpireTime: 6,
				},
				&TimeLimitValue{
					Value:      4,
					ExpireTime: 4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smallHeap.Adjust()
			if !reflect.DeepEqual(smallHeap.heapSlice, tt.want) {
				t.Errorf("smallHeap.Adjust() = %v, want %v", smallHeap.heapSlice, tt.want)
			}
		})
	}
}

func TestSmallHeap_Rise(t *testing.T) {
	initSmallHeap()
	tests := []struct {
		name string
		rise int
		want []*TimeLimitValue
	}{
		{
			name: "rise elem",
			rise: 1,
			want: []*TimeLimitValue{
				&TimeLimitValue{
					Value:      1,
					ExpireTime: 1,
				},
				&TimeLimitValue{
					Value:      6,
					ExpireTime: 6,
				},
				&TimeLimitValue{
					Value:      3,
					ExpireTime: 3,
				},
				&TimeLimitValue{
					Value:      5,
					ExpireTime: 5,
				},
				&TimeLimitValue{
					Value:      2,
					ExpireTime: 2,
				},
				&TimeLimitValue{
					Value:      4,
					ExpireTime: 4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smallHeap.Rise(tt.rise)
			if !reflect.DeepEqual(smallHeap.heapSlice, tt.want) {
				t.Errorf("smallHeap.Rise() = %v, want %v", smallHeap.heapSlice, tt.want)
			}
		})
	}
}

func TestSmallHeap_Sink(t *testing.T) {
	initSmallHeap()
	tests := []struct {
		name string
		sink int
		want []*TimeLimitValue
	}{
		{
			name: "sink elem",
			sink: 0,
			want: []*TimeLimitValue{
				&TimeLimitValue{
					Value:      1,
					ExpireTime: 1,
				},
				&TimeLimitValue{
					Value:      2,
					ExpireTime: 2,
				},
				&TimeLimitValue{
					Value:      3,
					ExpireTime: 3,
				},
				&TimeLimitValue{
					Value:      5,
					ExpireTime: 5,
				},
				&TimeLimitValue{
					Value:      6,
					ExpireTime: 6,
				},
				&TimeLimitValue{
					Value:      4,
					ExpireTime: 4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smallHeap.Sink(tt.sink)
			if !reflect.DeepEqual(smallHeap.heapSlice, tt.want) {
				t.Errorf("smallHeap.Sink() = %v, want %v", smallHeap.heapSlice, tt.want)
			}
		})
	}
}

func TestSmallHeap_DelRoot(t *testing.T) {
	initSmallHeap()
	tests := []struct {
		name string
		want []*TimeLimitValue
	}{
		{
			name: "del root",
			want: []*TimeLimitValue{
				&TimeLimitValue{
					Value:      2,
					ExpireTime: 2,
				},
				&TimeLimitValue{
					Value:      4,
					ExpireTime: 4,
				},
				&TimeLimitValue{
					Value:      3,
					ExpireTime: 3,
				},
				&TimeLimitValue{
					Value:      5,
					ExpireTime: 5,
				},
				&TimeLimitValue{
					Value:      6,
					ExpireTime: 6,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smallHeap.Adjust()
			smallHeap.DelRoot()
			if !reflect.DeepEqual(smallHeap.heapSlice, tt.want) {
				t.Errorf("smallHeap.DelRoot() = %v, want %v", smallHeap.heapSlice, tt.want)
			}
		})
	}
}
