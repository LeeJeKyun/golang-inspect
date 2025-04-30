package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// 예시 json 생성
	js := `{
    		"person": {
        		"name": "LYJ",
				"age": 27,
				"position": "Developer"
			},
			"items": [
				{
					"bob": "tom",
					"john": "mayer",
					"halo": "rose"
				},{
					"bob": "ttt",
					"john": "ma",
					"halo": "roro"
				}
			],
			"index": 123,
			"ace": 3353
		}`
	// string:interface{} 형태의 맵 생성
	jsMap := make(map[string]interface{})

	// json -> map
	json.Unmarshal([]byte(js), &jsMap)

	fmt.Println("jsonMap:", jsMap) //map[ace:3353 index:123 items:[map[bob:tom halo:rose john:mayer] map[bob:ttt halo:roro john:ma]] person:map[age:27 name:LYJ position:Developer]]

	//형 변환
	m := jsMap["person"].(map[string]interface{})

	//내부 요소 출력
	name := m["name"]
	age := m["age"]
	position := m["position"]
	fmt.Println("Inner Object:", name, age, position) //JYJ 27 Developer

	//형 변환(interface{} -> []interface{})
	s := jsMap["items"].([]interface{})
	fmt.Println("Inner Array:", s) //[map[bob:tom halo:rose john:mayer] map[bob:ttt halo:roro john:ma]]
}
