package core

import (
	"fmt"
	"github.com/houxiangr/go-localcache/core/lfu"
	"github.com/houxiangr/go-localcache/core/lru"
	"github.com/houxiangr/go-localcache/core/time_limit"
)

//todo 针对缓存穿透做优化
//todo 进行map分区，减少锁粒度
type Localcache interface {
	Start(variable map[string]interface{}) error
	Get(key string) interface{}              //get value
	Set(key string, value interface{}) error //set value
	DumpFile()                               //save cache in file
	ImportFile(filename string)              //init cache from file
	CacheToMap() map[string]interface{}      // cache to map
}

const (
	LRU       = "LRU"
	TimeLimit = "TL"
	LFU       = "LFU"
)

func GetLocalcache(outType string, variable map[string]interface{}) (Localcache, error) {
	var localcache Localcache
	switch outType {
	case LRU:
		localcache = &lru.LRULocalcache{}
	case TimeLimit:
		localcache = &time_limit.TimeLimitLocalcache{}
	case LFU:
		localcache = &lfu.LFULocalCache{}
	default:
		return nil, fmt.Errorf("not match cache type")
	}

	err := localcache.Start(variable)
	if err != nil {
		return nil, err
	}
	return localcache, nil
}
