package time_limit

import (
	"fmt"
	"github.com/houxiangr/go-localcache/core/time_limit/start_variable"
	"github.com/houxiangr/go-localcache/core/time_limit/ttypes"
	"reflect"
	"sync"
	"time"
)

const (
	TimeLimitValueType = "ttypes.TimeLimitValue"
)

/*    ___________________________________________________
     | 													|
     v                                                  |
CanCheckCache ----> CheckingCache ------> CheckCacheCycleTime
*/
const (
	CanCheckCache       = 1 //可以检查淘汰内存
	CheckingCache       = 2 //正在检查内存所以不能检查内容
	CheckCacheCycleTime = 3 //由于在检查内存的间隔周期内不能检查内存
)

const (
	DefaultCheckTime          = 10
	DefaultCheckCount         = 100
	DefaultCheckCountInterval = 3
)

type TimeLimitLocalcache struct {
	size      int
	used      int
	cacheMap  map[string]ttypes.TimeLimitValue
	smallHeap ttypes.SmallHeap
	lock      sync.RWMutex
	//增加正在清理缓存状态
	checkSwitch        int8
	checkTime          int
	checkCount         int
	checkCountInterval int
}

func (this *TimeLimitLocalcache) Start(variable map[string]interface{}) error {
	this.cacheMap = make(map[string]ttypes.TimeLimitValue)
	this.smallHeap = ttypes.SmallHeap{}
	var ok bool
	this.size, ok = variable[start_variable.SizeKey].(int)
	if !ok {
		return fmt.Errorf("start variable transfer fail")
	}
	this.checkTime, ok = variable[start_variable.CheckTimeKey].(int)
	if !ok {
		this.checkTime = DefaultCheckTime
	}
	this.checkCount, ok = variable[start_variable.CheckCount].(int)
	if !ok {
		this.checkCount = DefaultCheckCount
	}
	this.checkCountInterval, ok = variable[start_variable.CheckCountInterval].(int)
	if !ok {
		this.checkCount = DefaultCheckCountInterval
	}

	this.checkSwitch = CheckCacheCycleTime
	this.lock = sync.RWMutex{}
	go func() {
		for {
			time.Sleep(time.Second * time.Duration(this.checkTime))
			if this.checkSwitch == CheckingCache {
				time.Sleep(time.Second * time.Duration(this.checkTime))
			}

			if this.checkSwitch == CheckCacheCycleTime {
				this.checkSwitch = CanCheckCache
			}
		}
	}()
	return nil
}
func (this *TimeLimitLocalcache) Get(key string) interface{} {
	this.checkCache()
	this.lock.RLock()
	value := this.cacheMap[key].Value
	this.lock.RUnlock()
	return value
}

func (this *TimeLimitLocalcache) SetWithExpire(key string, value interface{}, expireTime int64) error {
	this.checkCache()
	if this.used >= this.size {
		return fmt.Errorf("local cache is filled")
	}
	err := this.Set(key, ttypes.TimeLimitValue{
		Key:        key,
		Value:      value,
		ExpireTime: time.Now().Add(time.Duration(expireTime) * time.Second).Unix(),
	})
	if err != nil {
		return err
	}
	this.used++
	return nil
}

func (this *TimeLimitLocalcache) Set(key string, value interface{}) error {
	if reflect.TypeOf(value).String() != TimeLimitValueType {
		return fmt.Errorf("time limit local cache set value type err")
	}
	timeLimitValue := value.(ttypes.TimeLimitValue)
	this.lock.Lock()
	this.cacheMap[key] = timeLimitValue
	this.smallHeap.Push(&timeLimitValue)
	this.lock.Unlock()
	return nil
}
func (this TimeLimitLocalcache) DumpFile() {

}
func (this *TimeLimitLocalcache) ImportFile(filename string) {

}

func (this *TimeLimitLocalcache) CacheToMap() map[string]interface{} {
	this.lock.RLock()
	defer this.lock.RUnlock()
	res := make(map[string]interface{})
	for k, v := range this.cacheMap {
		res[k] = map[string]interface{}{
			"key":         v.Key,
			"value":       v.Value,
			"expire_time": v.ExpireTime,
		}
	}
	return res
}

func (this *TimeLimitLocalcache) checkCache() {
	if this.checkSwitch != CanCheckCache {
		return
	}
	this.checkSwitch = CheckingCache
	currentTime := time.Now().Unix()
	this.lock.Lock()
	defer this.lock.Unlock()
	count := 0
	//优化淘汰顺序-堆
	this.smallHeap.Adjust()
	for {
		smallHeapRoot := this.smallHeap.GetRoot()
		if smallHeapRoot == nil {
			break
		}
		if smallHeapRoot.ExpireTime > currentTime {
			break
		}
		delete(this.cacheMap, smallHeapRoot.Key)
		this.smallHeap.DelRoot()
		this.used--
		count++
		//限制一次性淘汰体量优化
		if count >= this.checkCount {
			count = 0
			time.Sleep(time.Duration(this.checkCountInterval) * time.Second)
		}
	}
	this.checkSwitch = CheckCacheCycleTime
}
