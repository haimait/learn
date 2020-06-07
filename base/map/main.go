package main

import (
	"encoding/json"
	"fmt"
	g "haimait/learn/base/util"
)



func main() {
	test1()
}

func test1()  {
	data := map[string]interface{}{
		"title" : "title",
		"name" : "lili",
		"list" : []map[string]interface{}{
			map[string]interface{}{
				"date":    "2020-04-01",
				"name":    "朱元璋",
				"address": "江苏110号",
			},
			map[string]interface{}{
				"date":    "2020-04-01",
				"name":    "朱元璋",
				"address": "江苏110号",
			},
		},


	}
	fmt.Println(111111111)
	fmt.Println(data)
	datajson ,_:= json.Marshal(data)
	fmt.Println(string(datajson))
}

func test2()  {
	data := g.Map{
		"name":"lisi",
		"age":"19",
		"list":g.List{
			g.Map{
				"date":    "2020-04-01",
				"name":    "朱元璋11111111",
				"address": "江苏110号",
			},
			g.Map{
				"date":    "2020-04-01",
				"name":    "朱元璋2222",
				"address": "江苏110号",
			},
		},
	}
	fmt.Println(111111111)
	fmt.Println(data)
	dataJson ,_:= json.Marshal(data)
	fmt.Println(string(dataJson))

}