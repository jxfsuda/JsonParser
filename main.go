package JsonParser

func main() {
    json:=`{
	//example description //http://www.example.com/
	"key": "http://www.example.com/aaa.action?url=http%3a%2f%2fbbb.aaa.com%2fTools%2furlencode.asp",
	"key1": 19999,
	"key2": ["1","2"],
}`
	json, err := Discard(json)
	if err!=nil {
		print(err)
	}
	ret:= make(map[string]interface{})
	json.Unmarshal(json,ret)
	print(ret,json)

}
