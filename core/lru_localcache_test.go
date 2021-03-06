package core

import (
	"reflect"
	"testing"
)

var localcache LRU_localcache

func initEmptyLRUCache() {
	localcache = LRU_localcache{}
	localcache.Start(10)
}

func TestLRU_localcache_Get_And_Set(t *testing.T) {
	initEmptyLRUCache()
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
			localcache.Set(tt.key, tt.wantvalue)
			got := localcache.Get(tt.key)
			if !reflect.DeepEqual(got, tt.wantvalue) {
				t.Errorf("localcache.Get() = %v, want %v", got, tt.wantvalue)
			}
		})
	}
}

func initFillLruCache() {
	localcache = LRU_localcache{}
	localcache.Start(10)
	localcache.Set("1",1)
	localcache.Set("2",2)
	localcache.Set("3",3)
	localcache.Set("4",4)
	localcache.Set("5",5)
	localcache.Set("6",6)
	localcache.Set("7",7)
	localcache.Set("8",8)
	localcache.Set("9",9)
	localcache.Set("10",10)
}

func TestLRU_localcache_Set_Fill(t *testing.T) {
	initFillLruCache()
	tests := []struct {
		name      string
		key       string
		wantvalue interface{}
		wantNotExistKey string
	}{
		{
			name:      "set cache and get cache",
			key:       "11",
			wantvalue: 11,
			wantNotExistKey: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localcache.Set(tt.key, tt.wantvalue)
			got := localcache.Get(tt.key)
			if !reflect.DeepEqual(got, tt.wantvalue) {
				t.Errorf("localcache.Get() = %v, want %v", got, tt.wantvalue)
			}
			notExist := localcache.Get(tt.wantNotExistKey)
			if notExist != nil {
				t.Errorf("localcache.Get() not exist = %v, want %v", notExist, nil)
			}
		})
	}
}

func TestLRU_localcache_GetMoveHead(t *testing.T){
	initFillLruCache()
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
			localcache.Set(tt.key, tt.wantvalue)
			localcache.Get(tt.key)
			if !reflect.DeepEqual(localcache.linklistTwoWay.GetHead().GetValue(), tt.wantvalue) {
				t.Errorf("localcache,head = %v, want %v", localcache.linklistTwoWay.GetHead().GetValue(), tt.wantvalue)
			}
		})
	}
}
