package JsonParser

import (
	"testing"
)



type TestStruct struct{
	Key string	`json:"key"`
	Key1 int	`json:"key1"`
	Key2 []string	`json:"key2"`
}


func TestParser(t *testing.T) {
	jsonStr:=` { //注释
	//example description //http://www.example.com/
	"key": "http://www.example.com/aaa.action?url=http%3a%2f%2fbbb.aaa.com%2fTools%2furlencode.asp",  /*注释 https://342432 */ "key1": 19999,
/*
12344

*/
	"key2": ["1","2"]
} `


	ret:= &TestStruct{}
	err := UnMarshal(jsonStr, ret)
	t.Logf("json.key2 : %v ",ret )
	if err!=nil{
		t.Logf("---> %v  ",err)
	}
	j,_:= Marshal(ret)
	t.Logf("json.key3 : %v \n ---> %v", j ,ret.Key)


	//从文件读取
	ret =  &TestStruct{}
	UnmarshalByJsonFile("json.json", ret)
	j,_= Marshal(ret)
	t.Logf("json.key3 : %v \n ---> %v", j,ret.Key)
}
