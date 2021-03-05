package main

import (
	"fmt"
	"github.com/houxiangr/go-localcache/core"
)

func main(){
	localcache,_ := core.GetLocalcache(core.LRU)

	localcache.Set("key","value")
	fmt.Println(localcache.Get("key"))
}
