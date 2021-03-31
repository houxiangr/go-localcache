package lfu

import (
	"fmt"
	"github.com/houxiangr/go-localcache/core/lfu/ttypes"
	"sync"
)

type LFULocalcache struct {
	size      int
	cacheMap  map[string]*ttypes.LFUValue
	freqFloor map[int]*ttypes.LinkedMap
	minFloor  int

	sync.Mutex
}

func (this *LFULocalcache) Start(variable map[string]interface{}) error {
	this.cacheMap = make(map[string]*ttypes.LFUValue)
	this.freqFloor = make(map[int]*ttypes.LinkedMap)
	var ok bool
	this.size, ok = variable[SizeKey].(int)
	if !ok {
		return fmt.Errorf("start variable transfer fail")
	}
	return nil
}

func (this *LFULocalcache) Get(key string) interface{} {
	this.Lock()
	defer this.Unlock()
	lfuValue, ok := this.cacheMap[key]
	if !ok {
		return nil
	}

	this.freqFloor[lfuValue.Freq].DelKey(key)
	if lfuValue.Freq == this.minFloor && this.freqFloor[lfuValue.Freq].Len() == 0 {
		this.freqFloor[lfuValue.Freq] = nil
		this.minFloor++
	}

	lfuValue.Freq++
	if this.freqFloor[lfuValue.Freq] == nil {
		this.freqFloor[lfuValue.Freq] = ttypes.GetLinkMap()
	}
	this.freqFloor[lfuValue.Freq].SetTail(key)

	return lfuValue.Value
}

func (this *LFULocalcache) Set(key string, value interface{}) error {
	this.Lock()
	defer this.Unlock()
	if len(this.cacheMap) >= this.size {
		targetKey := this.freqFloor[this.minFloor].GetHead().GetKey()
		this.freqFloor[this.minFloor].DelHead()
		delete(this.cacheMap, targetKey)
	}

	lfuValue := ttypes.LFUValue{
		Value: value,
		Freq:  1,
	}

	this.cacheMap[key] = &lfuValue

	if this.freqFloor[lfuValue.Freq] == nil {
		this.freqFloor[lfuValue.Freq] = ttypes.GetLinkMap()
	}

	this.freqFloor[lfuValue.Freq].SetTail(key)

	this.minFloor = 1
	return nil
}

func (this *LFULocalcache) DumpFile() {

}

func (this *LFULocalcache) ImportFile(filename string) {

}

func (this *LFULocalcache) CacheToMap() map[string]interface{} {
	res := make(map[string]interface{})
	for k, _ := range this.cacheMap {
		res[k] = map[string]interface{}{
			"value": this.cacheMap[k].Value,
			"freq":  this.cacheMap[k].Freq,
		}
	}
	return res
}
