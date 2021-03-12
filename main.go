package main

import (
	"fmt"
	"github.com/houxiangr/go-localcache/core"
	"github.com/houxiangr/go-localcache/core/start_variable"
)

func main() {
	localcache, _ := core.GetLocalcache(core.LRU, map[string]interface{}{
		start_variable.SizeKey:10,
	})

	localcache.Set("1", "1")
	localcache.Set("2", "2")
	fmt.Println(localcache.Get("1"))
	localcache.Set("3", "3")
	localcache.Set("4", "4")
	localcache.Set("5", "5")
	fmt.Println(localcache.CacheToMap())
}

//package main
//
//import (
//	"fmt"
//	"reflect"
//	"sync/atomic"
//	"unsafe"
//)
//
//type T struct {
//	value int
//}
//
//func Swap(dest **T, old, new *T) {
//	udest := (*unsafe.Pointer)(unsafe.Pointer(dest))
//	res := atomic.CompareAndSwapPointer(udest,
//		unsafe.Pointer(old),
//		unsafe.Pointer(new),
//	)
//	fmt.Println(res)
//}
//
//func main() {
//	x := &T{42}
//	n := &T{50}
//	fmt.Println(*x, *n)
//
//	var temp interface{}
//	temp = 1
//	ptrTemp := &temp
//	fmt.Println(reflect.TypeOf(&ptrTemp))
//
//	p := x
//	Swap(&x, p, n)
//	fmt.Println(*x, *n)
//}
