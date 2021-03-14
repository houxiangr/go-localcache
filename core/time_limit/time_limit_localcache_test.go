package time_limit

import (
	"fmt"
	"github.com/houxiangr/go-localcache/core/time_limit/start_variable"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

var localcache TimeLimitLocalcache

func initEmptyTimeLocalcacheCache() {
	localcache = TimeLimitLocalcache{}
	localcache.Start(map[string]interface{}{
		start_variable.CheckTimeKey: 10,
		start_variable.SizeKey:      10,
	})
}

func initFillTimeLocalcacheCache() {
	localcache = TimeLimitLocalcache{}
	localcache.Start(map[string]interface{}{
		start_variable.CheckTimeKey: 10,
		start_variable.SizeKey:      10,
	})
	localcache.SetWithExpire("1", 1, 100)
	localcache.SetWithExpire("2", 1, 100)
	localcache.SetWithExpire("3", 1, 100)
	localcache.SetWithExpire("4", 1, 100)
	localcache.SetWithExpire("5", 1, 100)
	localcache.SetWithExpire("6", 1, 100)
	localcache.SetWithExpire("7", 1, 100)
	localcache.SetWithExpire("8", 1, 100)
	localcache.SetWithExpire("9", 1, 100)
	localcache.SetWithExpire("10", 1, 100)
}

func initChcekCountTimeLocalcacheCache() {
	localcache = TimeLimitLocalcache{}
	localcache.Start(map[string]interface{}{
		start_variable.CheckTimeKey:       1,
		start_variable.SizeKey:            10,
		start_variable.CheckCount:         1,
		start_variable.CheckCountInterval: 1,
	})
	localcache.SetWithExpire("1", 1, 100)
	localcache.SetWithExpire("2", 1, 100)
	localcache.SetWithExpire("3", 1, 100)
	localcache.SetWithExpire("4", 1, 100)
	localcache.SetWithExpire("5", 1, 100)
	localcache.SetWithExpire("6", 1, 100)
}

func TestTimeLimitLocalcache_SetAndGet(t *testing.T) {
	initEmptyTimeLocalcacheCache()
	tests := []struct {
		name      string
		key       string
		wantvalue interface{}
	}{
		{
			name:      "set cache and get cache",
			key:       "1",
			wantvalue: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := localcache.SetWithExpire(tt.key, tt.wantvalue, 60)
			if err != nil {
				t.Error(err)
			}
			got := localcache.Get(tt.key)
			if !reflect.DeepEqual(got, tt.wantvalue) {
				t.Errorf("localcache.Get() = %v, want %v", got, tt.wantvalue)
			}
		})
	}
}

func TestTimeLimitLocalcache_SetWithExpire(t *testing.T) {
	initEmptyTimeLocalcacheCache()
	tests := []struct {
		name      string
		key       string
		value     int
		wait      int64
		wantvalue interface{}
	}{
		{
			name:      "set cache and get cache with out expire time",
			key:       "1",
			value:     1,
			wantvalue: nil,
			wait:      1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localcache.SetWithExpire(tt.key, tt.wantvalue, tt.wait)
			time.Sleep(time.Second * 2)
			got := localcache.Get(tt.key)
			if !reflect.DeepEqual(got, tt.wantvalue) {
				t.Errorf("localcache.Get() = %v, want %v", got, tt.wantvalue)
			}
		})
	}
}

func TestTimeLimitLocalcache_SetWithExpireErr(t *testing.T) {
	initFillTimeLocalcacheCache()
	tests := []struct {
		name    string
		key     string
		value   int
		wait    int64
		wanterr string
	}{
		{
			name:    "cache is filled",
			key:     "1",
			value:   1,
			wanterr: "local cache is filled",
			wait:    1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := localcache.SetWithExpire(tt.key, tt.wanterr, tt.wait)
			if err == nil || !strings.Contains(err.Error(), tt.wanterr) {
				t.Errorf("localcache.SetWithExpireErr() = %v, want %v", err.Error(), tt.wanterr)
			}
		})
	}
}

func TestTimeLimitLocalcache_SetErr(t *testing.T) {
	initFillTimeLocalcacheCache()
	tests := []struct {
		name    string
		key     string
		value   int
		wanterr string
	}{
		{
			name:    "direct use func err",
			key:     "1",
			value:   1,
			wanterr: "time limit local cache set value type err",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := localcache.Set(tt.key, tt.wanterr)
			if err == nil || !strings.Contains(err.Error(), tt.wanterr) {
				t.Errorf("localcache.SetWithExpireErr() = %v, want %v", err.Error(), tt.wanterr)
			}
		})
	}
}

func TestTimeLimitLocalcache_CheckCount(t *testing.T) {
	initChcekCountTimeLocalcacheCache()
	tests := []struct {
		name       string
		key        string
		value      int
		expireTime int64
		wanterr    string
	}{
		{
			name:       "check count and sleep time",
			key:        "10",
			value:      1,
			expireTime: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localcache.SetWithExpire(tt.key, tt.value, tt.expireTime)
			time.Sleep(1 * time.Second)
			localcache.Get("1")
			time.Sleep(2 * time.Second)
		})
	}
}

func TestMutiGoroutine(t *testing.T) {
	initEmptyTimeLocalcacheCache()
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(index int) {
			defer wg.Done()
			localcache.SetWithExpire(fmt.Sprintf("muti_%d", index), fmt.Sprintf("muti_%d", index), 1)
			got := localcache.Get(fmt.Sprintf("muti_%d", index))
			if got != nil && got != fmt.Sprintf("muti_%d", index) {
				t.Errorf("got = %+v,want = %s", got, fmt.Sprintf("muti_%d", index))
			}
		}(i)
	}
	wg.Wait()
}
