package core

import (
	"github.com/houxiangr/go-localcache/common"
	"sync"
)

type LRU_localcache struct {
	size int64
	used int64

	linklistTwoWay common.LinklistTwoWay
	cacheMap       map[string]*common.LinkNode

	lock sync.Mutex
}

func (this *LRU_localcache) Start(size int64) {
	this.cacheMap = make(map[string]*common.LinkNode)
	this.size = size
}
func (this *LRU_localcache) Get(key string) interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	res := this.cacheMap[key]
	if res == nil {
		return nil
	}
	this.linklistTwoWay.MoveNodeToHead(res)
	return res.GetValue()
}
func (this *LRU_localcache) Set(key string, value interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.used >= this.size {
		this.slowSet(key, value)
		return
	}
	this.used++

	this.linklistTwoWay.SetHead(key, value)
	this.cacheMap[key] = this.linklistTwoWay.GetHead()
}
func (this LRU_localcache) DumpFile() {

}
func (this *LRU_localcache) ImportFile(filename string) {

}

func (this *LRU_localcache) slowSet(key string, value interface{}) {
	// del node
	tailNode := this.linklistTwoWay.GetTail()
	delete(this.cacheMap, tailNode.GetKey())
	this.linklistTwoWay.DelNode(tailNode)

	// set node
	this.linklistTwoWay.SetHead(key, value)
	this.cacheMap[key] = this.linklistTwoWay.GetHead()
}

func (this *LRU_localcache) CacheToMap() map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range this.cacheMap {
		res[k] = v.GetValue()
	}
	return res
}
