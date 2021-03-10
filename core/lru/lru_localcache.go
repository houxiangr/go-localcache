package lru

import (
	"github.com/houxiangr/go-localcache/core/lru/ttypes"
	"sync"
)

type LRULocalcache struct {
	size int64
	used int64

	linklistTwoWay ttypes.LinklistTwoWay
	cacheMap       map[string]interface{}
	lock           sync.Mutex
}

func (this *LRULocalcache) Start(size int64) {
	this.cacheMap = make(map[string]interface{})
	this.lock = sync.Mutex{}
	this.size = size
}
func (this *LRULocalcache) Get(key string) interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	res, ok := this.cacheMap[key]
	if !ok || res == nil {
		return nil
	}
	resValue := res.(*ttypes.LinkNode)
	this.linklistTwoWay.MoveNodeToHead(resValue)
	return resValue.GetValue()
}
func (this *LRULocalcache) Set(key string, value interface{}) error {
	this.lock.Lock()
	res, ok := this.cacheMap[key]
	this.lock.Unlock()
	if ok {
		resValue := res.(*ttypes.LinkNode)
		resValue.TrySetValue(value)
		this.linklistTwoWay.MoveNodeToHead(resValue)
		return nil
	}
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.used >= this.size {
		this.slowSet(key, value)
		return nil
	}
	this.used++
	this.cacheMap[key] = this.linklistTwoWay.SetHead(key, value)
	return nil
}
func (this LRULocalcache) DumpFile() {

}
func (this *LRULocalcache) ImportFile(filename string) {

}

func (this *LRULocalcache) slowSet(key string, value interface{}) {
	delete(this.cacheMap, this.linklistTwoWay.GetTail().GetKey())
	// del node and set new node
	newNode := this.linklistTwoWay.DelTailAndSetHead(key, value)

	this.cacheMap[key] = newNode
}

func (this *LRULocalcache) CacheToMap() map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range this.cacheMap {
		res[k] = v
	}
	return res
}
