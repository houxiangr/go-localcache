package lfu

import (
	"fmt"
	"github.com/houxiangr/go-localcache/core/lfu/ttypes"
	"sync"
)

type LFULocalCache struct {
	size      int
	cacheMap  map[string]ttypes.LFUValue
	freqFloor map[string]ttypes.FreqFloor

	lock sync.Mutex
}

func (this *LFULocalCache) Start(variable map[string]interface{}) error {
	this.cacheMap = make(map[string]ttypes.LFUValue)
	this.freqFloor = make(map[string]ttypes.FreqFloor)
	this.lock = sync.Mutex{}
	var ok bool
	this.size, ok = variable[SizeKey].(int)
	if !ok {
		return fmt.Errorf("start variable transfer fail")
	}
	return nil
}

func (this *LFULocalCache) Get(key string) interface{} {

	return nil
}

func (this *LFULocalCache) Set(key string, value interface{}) error {
	return nil
}

func (this *LFULocalCache) DumpFile() {

}

func (this *LFULocalCache) ImportFile(filename string) {

}

func (this *LFULocalCache) CacheToMap() map[string]interface{} {
	return nil
}
