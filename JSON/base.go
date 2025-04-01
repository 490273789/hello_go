package json_test

import (
	"encoding/json"
	"fmt"
)

// json javascript object notation 是一种比XML更轻量的数据传输方式，易于读和编写，也易于程序的解析和生成。

func CallJson() {
	sToJ()
	mToJ()
	jToS()
	mToS()
}

type Company struct {
	Name    string   `json:"-"`       // - 隐藏这个字段
	Subject []string `json:"subject"` // 定义json的名字
	Address string   `json:"address"`
	IsOpen  bool     `json:"isOpen"`
	Price   float32  `json:"price,string"` // 转义字段的类型
}

// 将结构体转换为JSON
func sToJ() {
	company := Company{"六度云", []string{"go", "python"}, "龙跃苑", true, 2000}
	// buf, err := json.Marshal(company)
	buf, err := json.MarshalIndent(company, "", "	") // 格式化json
	if err != nil {
		fmt.Printf("JSON格式转化错误")
	}

	fmt.Println("json的转化结果为：", string(buf))
}

func mToJ() {
	company := make(map[string]interface{}, 4)
	company["company"] = "华控"
	company["address"] = "3区"
	company["subject"] = []string{"go", "java"}
	company["isOpen"] = true
	company["price"] = 300

	buf, err := json.MarshalIndent(company, "", "	")
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Println("Map转json的转化结果：", string(buf))
}

func jToS() {
	jsonBuf := `{
			"subject": [
				"go",
				"python"
			],
			"address": "龙跃苑",
			"isOpen": true,
			"price": "2000",
			"name": "wsn"
		}`

	var temp Company
	err := json.Unmarshal([]byte(jsonBuf), &temp)
	if err != nil {
		fmt.Println("err = ", err)
	}

	fmt.Printf("json转struct：%+v \n", temp)
}

func mToS() {
	jsonBuf := `{
			"subject": [
				"go",
				"python"
			],
			"address": "龙跃苑",
			"isOpen": true,
			"price": "2000",
			"name": "wsn"
		}`

	company := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(jsonBuf), &company)
	if err != nil {
		fmt.Println("err = ", err)
	}
	// 读取需要借助类型断言
	for key, value := range company {
		switch data := value.(type) {
		case string:
			fmt.Printf("map[%s]的值类型为string，value-%s \n", key, data)
		case bool:
			fmt.Printf("map[%s]的值类型为string，value-%t \n", key, data)
		case []string:
			fmt.Printf("map[%s]的值类型为[]string，value-%v \n", key, data)
		case float64:
			fmt.Printf("map[%s]的值类型为float64，value-%f \n", key, data)
		case []interface{}:
			fmt.Printf("map[%s]的值类型为interface，value-%v \n", key, data)
		}
	}

	fmt.Printf("json转struct：%+v \n", company)
}
