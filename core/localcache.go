package core

import (
	"github.com/houxiangr/go-localcache/core/lru"
	"github.com/houxiangr/go-localcache/core/time_limit"
)

type Localcache interface {
	Start(size int64)
	Get(key string) interface{}              //get value
	Set(key string, value interface{}) error //set value
	DumpFile()                               //save cache in file
	ImportFile(filename string)              //init cache from file
	CacheToMap() map[string]interface{}      // cache to map
}

const (
	LRU       = "LRU"
	TimeLimit = "time_limit"
)

func GetLocalcache(outType string, size int64) (Localcache, error) {
	var localcache Localcache
	switch outType {
	case LRU:
		localcache = &lru.LRULocalcache{}
	case TimeLimit:
		localcache = &time_limit.TimeLimitLocalcache{}
	default:
		return nil, nil
	}

	localcache.Start(size)
	return localcache, nil
}
