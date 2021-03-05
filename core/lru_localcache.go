package core

type LRU_localcache struct{
	size int
	used int
}

func (this LRU_localcache)Start(){

}
func (this LRU_localcache)Get(key string)interface{}{
	return nil
}
func (this *LRU_localcache)Set(key string,value interface{}){

}
func (this LRU_localcache)DumpFile(){

}
func (this *LRU_localcache)ImportFile(filename string){

}