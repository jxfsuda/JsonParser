# JsonParser Plus
 Golang json parser 
 扩展   
 去除 空格过滤   
 加入url:// 转义为 :\/\/  
 
 使用   
 <code>  import "https://github.com/jxfsuda/JsonParser"   
   
    ret:= &TestStruct{}   
    err := JsonParser.UnMarshal(jsonStr, ret)   


    j,_:= JsonParser.Marshal(ret)


 </code> 
