package main

import (
	"fmt"
	"github.com/houxiangr/go-localcache/core"
	"github.com/houxiangr/go-localcache/core/lru"
)

func main() {
	localcache, _ := core.GetLocalcache(core.LFU, map[string]interface{}{
		lru.SizeKey: 4,
	})

	localcache.Set("1", "1")
	localcache.Set("2", "2")
	//fmt.Println(localcache.Get("1"))
	localcache.Get("1")
	localcache.Set("3", "3")
	localcache.Set("4", "4")
	localcache.Set("5", "5")
	fmt.Println(localcache.CacheToMap())
}
