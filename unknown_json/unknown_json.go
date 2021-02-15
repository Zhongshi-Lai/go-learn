package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// TODO(laizhongshi): 验证一下
	// 注意,json数据里面数组会被转换为[] interface{}
	// json里面的对象会被转换为map[string] interface{}

	// 准备一段json
	b := []byte(`{
		"Title": "Go语言编程",
		"Authors": ["XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan",
		"XuDaoli"],
		"Publisher": "ituring.com.cn",
		"IsPublished": true,
		"Price": 9.99,
		"Sales": 1000000
		}`)
	var r interface{}
	err := json.Unmarshal(b, &r)

	fmt.Println(err)
	fmt.Println(r)

	// 一个一个字段的判断是什么类型

	gobook, ok := r.(map[string]interface{})
	if ok {
		for k, v := range gobook {
			switch v2 := v.(type) {
			case string:
				fmt.Println(k, "is string", v2)
			case int:
				fmt.Println(k, "is int", v2)
			case bool:
				fmt.Println(k, "is bool", v2)
			case []interface{}:
				fmt.Println(k, "is array:")
				for i, iv := range v2 {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println(k, "is unknown_type")
			}
		}
	}

}
