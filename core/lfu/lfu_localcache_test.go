package lfu

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

var lfuLocalcache LFULocalcache

func initEmptyLFUCache() {
	lfuLocalcache = LFULocalcache{}
	lfuLocalcache.Start((map[string]interface{}{
		SizeKey: 10,
	}))
}

func TestLRU_localcache_Get_And_Set(t *testing.T) {
	initEmptyLFUCache()
	tests := []struct {
		name      string
		key       string
		wantvalue interface{}
	}{
		{
			name:      "set cache and get cache",
			key:       "1",
			wantvalue: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lfuLocalcache.Set(tt.key, tt.wantvalue)
			got := lfuLocalcache.Get(tt.key)
			if !reflect.DeepEqual(got, tt.wantvalue) {
				t.Errorf("localcache.Get() = %v, want %v", got, tt.wantvalue)
			}
		})
	}
}

func initFillLruCache() {
	lfuLocalcache = LFULocalcache{}
	lfuLocalcache.Start((map[string]interface{}{
		SizeKey: 10,
	}))
	lfuLocalcache.Set("1", 1)
	lfuLocalcache.Set("2", 2)
	lfuLocalcache.Set("3", 3)
	lfuLocalcache.Set("4", 4)
	lfuLocalcache.Set("5", 5)
	lfuLocalcache.Set("6", 6)
	lfuLocalcache.Set("7", 7)
	lfuLocalcache.Set("8", 8)
	lfuLocalcache.Set("9", 9)
	lfuLocalcache.Set("10", 10)

}

func TestLRU_localcache_Set_Fill(t *testing.T) {
	initFillLruCache()
	lfuLocalcache.Get("10")
	lfuLocalcache.Get("9")
	lfuLocalcache.Get("8")
	lfuLocalcache.Get("7")
	lfuLocalcache.Get("6")
	lfuLocalcache.Get("1")
	tests := []struct {
		name            string
		key             string
		wantvalue       interface{}
		wantNotExistKey string
	}{
		{
			name:            "set cache and get cache",
			key:             "11",
			wantvalue:       11,
			wantNotExistKey: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lfuLocalcache.Set(tt.key, tt.wantvalue)
			got := lfuLocalcache.Get(tt.key)
			if !reflect.DeepEqual(got, tt.wantvalue) {
				t.Errorf("localcache.Get() = %v, want %v", got, tt.wantvalue)
			}
			notExist := lfuLocalcache.Get(tt.wantNotExistKey)
			if notExist != nil {
				t.Errorf("localcache.Get() not exist = %v, want %v", notExist, nil)
			}
		})
	}
}

func TestMutiGoroutine(t *testing.T) {
	initEmptyLFUCache()
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(index int) {
			defer wg.Done()
			lfuLocalcache.Set(fmt.Sprintf("muti_%d", index), fmt.Sprintf("muti_%d", index))
			got := lfuLocalcache.Get(fmt.Sprintf("muti_%d", index))
			if got != nil && got != fmt.Sprintf("muti_%d", index) {
				t.Errorf("got = %+v,want = %s", got, fmt.Sprintf("muti_%d", index))
			}
		}(i)
	}
	wg.Wait()
}
