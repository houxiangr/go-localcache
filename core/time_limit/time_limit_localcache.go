package time_limit

import (
	"fmt"
	"reflect"
	"time"
)

const (
	TimeLimitValueType = "time_limit.timeLimitValue"
)

type timeLimitValue struct {
	value      interface{}
	expireTime int64
}

type TimeLimitLocalcache struct {
	size        int64
	used        int64
	cacheMap    map[string]timeLimitValue
	checkSwitch bool
}

func (this *TimeLimitLocalcache) Start(size int64) {
	this.cacheMap = make(map[string]timeLimitValue)
	this.size = size
	this.checkSwitch = false
	go func() {
		for {
			this.checkSwitch = !this.checkSwitch
			time.Sleep(time.Second * 10)
		}
	}()
}
func (this *TimeLimitLocalcache) Get(key string) interface{} {
	this.checkCache()
	return this.cacheMap[key].value
}

func (this *TimeLimitLocalcache) SetWithExpire(key string, value interface{}, expireTime int64) error {
	this.checkCache()
	if this.used >= this.size {
		return fmt.Errorf("local cache is filled")
	}
	err := this.Set(key, timeLimitValue{
		value:      value,
		expireTime: time.Now().Add(time.Duration(expireTime) * time.Second).Unix(),
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
	this.cacheMap[key] = value.(timeLimitValue)
	return nil
}
func (this TimeLimitLocalcache) DumpFile() {

}
func (this *TimeLimitLocalcache) ImportFile(filename string) {

}

func (this *TimeLimitLocalcache) CacheToMap() map[string]interface{} {
	return nil
}

func (this *TimeLimitLocalcache) checkCache() {
	if !this.checkSwitch {
		return
	}
	for k, v := range this.cacheMap {
		if v.expireTime < time.Now().Unix() {
			delete(this.cacheMap, k)
			this.used--
		}
	}
}
