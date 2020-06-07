package main

import (
	"encoding/json"
	"fmt"
	_ "gf01/boot"
	_ "gf01/router"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	// 常规注册
	group := s.Group("/")

	// 模板文件
	group.GET("/", func(r *ghttp.Request) {
		dataMap:=g.Map{
			"title": "列表页面",
			"show": true,
			"desc":  "desc1111111",
			"listData": g.List{
				g.Map{
					"date":    "2020-04-01",
					"name":    "朱元璋",
					"address": "江苏110号",
				},
				g.Map{
					"date":    "2020-04-02",
					"name":    "徐达",
					"address": "江苏111号",
				},
				g.Map{
					"date":    "2020-04-03",
					"name":    "李善长",
					"address": "江苏112号",
				},
			}}
		dataJson ,_:=json.Marshal(dataMap)
		fmt.Println(111111111)
		fmt.Println(string(dataJson))
		dataMap2:=g.Map{"data":string(dataJson)}

		r.Response.WriteTpl("index.html",dataMap2)
	})

	// 模板文件
	group.GET("/index2", func(r *ghttp.Request) {
		//var item  map[string]interface{}
		//var list  []map[string]interface{}
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

		dataMap:=g.Map{
			"title": "列表页面",
			"show": true,
			"desc":  "desc1111111",
			"listData": g.List{
				g.Map{
					"date":    "2020-04-01",
					"name":    "朱元璋",
					"address": "江苏110号",
				},
				g.Map{
					"date":    "2020-04-02",
					"name":    "徐达",
					"address": "江苏111号",
				},
				g.Map{
					"date":    "2020-04-03",
					"name":    "李善长",
					"address": "江苏112号",
				},
			}}
		r.Response.WriteTpl("index2.html",dataMap)
	})

	// 字符串传入
	group.GET("/template", func(r *ghttp.Request) {
		tplContent := `id:${.id}, name:${.name}`
		r.Response.WriteTplContent(tplContent, g.Map{
			"id"   : 123,
			"name" : "john",
		})
	})

	s.Run()
}