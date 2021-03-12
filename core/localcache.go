package core

import (
	"fmt"
	"github.com/houxiangr/go-localcache/core/lru"
	"github.com/houxiangr/go-localcache/core/time_limit"
)

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
)

func GetLocalcache(outType string, variable map[string]interface{}) (Localcache, error) {
	var localcache Localcache
	switch outType {
	case LRU:
		localcache = &lru.LRULocalcache{}
	case TimeLimit:
		localcache = &time_limit.TimeLimitLocalcache{}
	default:
		return nil, fmt.Errorf("not match cache type")
	}

	err := localcache.Start(variable)
	if err != nil {
		return nil, err
	}
	return localcache, nil
}
