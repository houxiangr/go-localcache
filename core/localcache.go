package core

type Localcache interface {
	Start()
	Get(key string)interface{} //get value
	Set(key string,value interface{}) //set value
	DumpFile() //save cache in file
	ImportFile(filename string) //init cache from file
}

const(
	LRU = "LRU"
)

func GetLocalcache(outType string)(Localcache,error){
	var localcache Localcache
	switch outType{
	case LRU:
		localcache = &LRU_localcache{}
	default:
		return nil,nil
	}
	localcache.Start()
	return localcache,nil
}
