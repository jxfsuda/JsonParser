package JsonParser

import (
	"testing"
)

type TestStruct struct{
	Key string	`json:"key"`
	Key1 string	`json:"key1"`
	Key2 []string	`json:"key2"`
}


//go test --cover Parser_test.go Parser.go
// go test Parser_test.go Parser.go
func TestParser(t *testing.T) {
    json:=`{
	//example description //http://www.example.com/
	"key": "http://www.example.com/aaa.action?url=http%3a%2f%2fbbb.aaa.com%2fTools%2furlencode.asp",
// wowowowow
	"key1": 19999,
/*
12344

*/
	"key2": ["1","2"]
}`

	json, err := Discard(json)
	if err!=nil {
		t.Logf("%v",err)
	}
	t.Logf("%v",json)
	ret:= &TestStruct{}
	t.Logf("json.key1 : %v ",ret )
	err = UnMarshal(json, ret)
	t.Logf("json.key2 : %v ",ret )
	if err!=nil{
		t.Logf("---> %v , \n %v \n",err,json)
	}
	t.Logf("json.key3 : %v \n  %v",ret,ret.Key)
	t.Logf("json对象: %v ", json)

	t.Logf("json对象_end: %v ", json)


}
