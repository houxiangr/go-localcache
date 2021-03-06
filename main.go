package main

import (
	"fmt"
	"github.com/houxiangr/go-localcache/core"
)

func main() {
	localcache, _ := core.GetLocalcache(core.LRU, 2)

	localcache.Set("1", "1")
	localcache.Set("2", "2")
	fmt.Println(localcache.Get("1"))
	localcache.Set("3", "3")
	localcache.Set("4", "4")
	localcache.Set("5", "5")
	fmt.Println(localcache.CacheToMap())
}
