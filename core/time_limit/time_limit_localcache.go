package time_limit

import (
	"fmt"
	"github.com/houxiangr/go-localcache/core/start_variable"
	"github.com/houxiangr/go-localcache/core/time_limit/ttypes"
	"reflect"
	"sync"
	"time"
)

const (
	TimeLimitValueType = "time_limit.timeLimitValue"
)

type TimeLimitLocalcache struct {
	size        int
	used        int
	cacheMap    map[string]ttypes.TimeLimitValue
	lock        sync.RWMutex
	//todo 增加正在清理缓存状态
	checkSwitch bool
}

func (this *TimeLimitLocalcache) Start(variable map[string]interface{}) error {
	this.cacheMap = make(map[string]ttypes.TimeLimitValue)
	var ok bool
	this.size, ok = variable[start_variable.SizeKey].(int)
	if !ok {
		return fmt.Errorf("start variable transfer fail")
	}
	checkTime, ok := variable[start_variable.CheckTimeKey].(int)
	if !ok {
		return fmt.Errorf("start variable transfer fail")
	}
	this.checkSwitch = false
	this.lock = sync.RWMutex{}
	go func() {
		for {
			this.checkSwitch = !this.checkSwitch
			time.Sleep(time.Second * time.Duration(checkTime))
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
	this.lock.Lock()
	this.cacheMap[key] = value.(ttypes.TimeLimitValue)
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
			"value":       v.Value,
			"expire_time": v.ExpireTime,
		}
	}
	return res
}

//todo 限制一次性淘汰体量优化
func (this *TimeLimitLocalcache) checkCache() {
	if !this.checkSwitch {
		return
	}
	this.lock.Lock()
	defer this.lock.Unlock()
	//todo 优化淘汰顺序-堆
	for k, v := range this.cacheMap {
		if v.ExpireTime < time.Now().Unix() {
			delete(this.cacheMap, k)
			this.used--
		}
	}
}
