package common

import (
	"testing"
)

var linklist LinklistTwoWay

func init(){
	linklist = LinklistTwoWay{}
}

func TestLinklist(t *testing.T){
	// 默认都为空
	if linklist.GetHead() != nil || linklist.GetTail() != nil {
		t.Fail()
	}
	if linklist.GetHead().GetValue() != nil || linklist.GetTail().GetValue() != nil {
		t.Fail()
	}

	//插入1
	//1
	linklist.SetHead(1)
	if linklist.GetHead().GetValue() != 1 {
		t.Fail()
	}
	if linklist.GetTail().GetValue() != 1 {
		t.Fail()
	}

	//插入2
	//2-1
	linklist.SetHead(2)
	if linklist.GetHead().GetNext().GetValue() != 1{
		t.Fail()
	}
	if linklist.GetTail().GetPre().GetValue() != 2{
		t.Fail()
	}

	//插入3
	//2-1-3
	linklist.SetTail(3)
	if linklist.GetHead().GetNext().GetValue() != 1{
		t.Fail()
	}
	if linklist.GetTail().GetPre().GetValue() != 1{
		t.Fail()
	}

	//删除1
	//2-3
	linklist.DelNode(linklist.GetHead().GetNext())
	if linklist.GetHead().GetNext().GetValue() != 3{
		t.Fail()
	}
	if linklist.GetTail().GetPre().GetValue() != 2{
		t.Fail()
	}

	//删除2
	//3
	linklist.DelNode(linklist.GetHead())
	if linklist.GetHead().GetValue() != 3{
		t.Fail()
	}
	if linklist.GetTail().GetValue() != 3{
		t.Fail()
	}

	//删除3
	//
	linklist.DelNode(linklist.GetTail())
	if linklist.GetHead().GetValue() != nil{
		t.Fail()
	}
	if linklist.GetTail().GetValue() != nil{
		t.Fail()
	}


}
