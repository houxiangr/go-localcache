package core

type Localcache interface {
	Start(size int64)
	Get(key string) interface{}         //get value
	Set(key string, value interface{})  //set value
	DumpFile()                          //save cache in file
	ImportFile(filename string)         //init cache from file
	CacheToMap() map[string]interface{} // cache to map
}

const (
	LRU = "LRU"
)

func GetLocalcache(outType string, size int64) (Localcache, error) {
	var localcache Localcache
	switch outType {
	case LRU:
		localcache = &LRU_localcache{}
	default:
		return nil, nil
	}

	localcache.Start(size)
	return localcache, nil
}
