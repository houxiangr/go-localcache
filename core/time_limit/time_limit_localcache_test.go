package time_limit

import (
	"testing"
)

var localcache TimeLimitLocalcache

func initEmptyTimeLocalcacheCache() {
	localcache = TimeLimitLocalcache{}
	localcache.Start(10)
}

func TestTimeLimitLocalcache_Set(t *testing.T) {

}
