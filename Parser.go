package JsonParser

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

//Map alias map[string]string
type Map map[string]string

//Maches discard from "start" to "end"
var Maches = []Map{
	Map{"start": "//", "end": "\n"},
	Map{"start": "/*", "end": "*/"},
}

/**
去除注释
转义url
 */
func Discard(content string) (string, error) {
	var (
		buffer    bytes.Buffer
		flag      int
		v         rune
		protected bool
	)
	content=strings.ReplaceAll(content,"://",":\\/\\/")
	runes := []rune(content)
	flag = -1
	for i := 0; i < len(runes); {
		v = runes[i]
		if flag == -1 {
			//match start
			for f, v := range Maches {
				l := match(&runes, i, v["start"])
				if l != 0 {
					flag = f
					i += l
					break
				}
			}
			if flag == -1 {
				if protected {
					buffer.WriteRune(v)
					if v == '"' {
						protected = true
					}
				} else {
					r := filter(v)
					if r != 0 {
						buffer.WriteRune(v)
					}
				}
			} else {
				continue
			}
		} else {
			//match end
			l := match(&runes, i, Maches[flag]["end"])
			if l != 0 {
				flag = -1
				i += l
				continue
			}
		}
		i++
	}
	return buffer.String(), nil
}

func filter(v rune) rune {
	switch v {
	//	case ' ':
	case '\n':
	case '\t':
	default:
		return v
	}
	return 0
}

func match(runes *[]rune, i int, dst string) int {
	dstLen := len([]rune(dst))
	// fmt.Println("dstLen:", dstLen, ", index:", i, ",runesLen:", len(*runes))
	// fmt.Println(string((*runes)[i : i+dstLen]))
	if len(*runes)-i >= dstLen && string((*runes)[i:i+dstLen]) == dst {
		return dstLen
	}
	return 0
}


/*
反序列化
*/
func UnMarshal(str string,v interface{}) error{
	str,err:=Discard(str)
	if err!=nil {
		return  err
	}
	err = json.Unmarshal([]byte(str), v)
	return err
}

/**
序列化
 */
func Marshal(v interface{}) (string,error){
	j,err:=json.Marshal(v)
	if err!=nil {
		return "",err
	}

	return string(j),nil
}

/**
从文件读取json
 */
func UnmarshalByJsonFile(filePath string,v interface{}) error {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	jsonStr := string(byteValue[:])
	err =  UnMarshal(jsonStr,v)
	return err
}









