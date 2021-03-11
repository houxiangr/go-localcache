package time_limit

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

var localcache TimeLimitLocalcache

func initEmptyTimeLocalcacheCache() {
	localcache = TimeLimitLocalcache{}
	localcache.Start(10)
}

func initFillTimeLocalcacheCache() {
	localcache = TimeLimitLocalcache{}
	localcache.Start(10)
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
	localcache.SetWithExpire("11", 1, 100)
}

func TestTimeLimitLocalcache_SetAndGet(t *testing.T) {
	initEmptyTimeLocalcacheCache()
	tests := []struct {
		name      string
		key       string
		value     int
		wantvalue interface{}
	}{
		{
			name:      "set cache and get cache",
			key:       "1",
			value:     1,
			wantvalue: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localcache.SetWithExpire(tt.key, tt.wantvalue, 60)
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
			if err==nil || !strings.Contains(err.Error(), tt.wanterr) {
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
			if err==nil || !strings.Contains(err.Error(), tt.wanterr) {
				t.Errorf("localcache.SetWithExpireErr() = %v, want %v", err.Error(), tt.wanterr)
			}
		})
	}
}
