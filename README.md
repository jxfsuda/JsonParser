# JsonParser Plus
 Golang json parser 扩展 ,过滤部分代码来源: https://github.com/sipt/GoJsoner       
 去除 空格过滤   
 加入url:// 转义为 :\\/\\/   
 增加 Marshal UnMarshal 的快捷方式,直接使用string和返回string   
 
 
 
 
 
 
 使用   
 <code>  import "github.com/jxfsuda/JsonParser"   
   
    ret:= &TestStruct{}   
    err := JsonParser.UnMarshal(jsonStr, ret)   


    j,_:= JsonParser.Marshal(ret)

    UnmarshalByJsonFile("json.json", ret)
 </code> 
