package main

import (
	"encoding/json"
	"fmt"
)

type OrderIterm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID         string       `json:"id"`
	Iterm      []OrderIterm `json:"iterm"`
	Quantity   int          `json:"quantity"`
	TotalPrice float64      `json:"totalPrice"`
}

func main() {

	marshal()

	unmarshal()

	parseNLP()

	parseNLP2()
}

func marshal() {
	o := Order{
		ID: "1234",
		Iterm: []OrderIterm{
			{
				ID:   "iterm_1",
				Name: "go go go",
			},
			{
				ID:   "iterm_2",
				Name: "java java java",
			},
		},
		Quantity:   3,
		TotalPrice: 30,
	}

	fmt.Printf("%v\n", o)
	fmt.Printf("%+v\n", o)

	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)

}

func unmarshal() {
	s := `{"id":"1234","iterm":[{"id":"iterm_1","name":"go go go"},{"id":"iterm_2","name":"java java java"}],"quantity":3,"totalPrice":30}`
	var o Order
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
}

func parseNLP() {
	res := `{
		"data": [
			{
				"synonym":"",
				"weight":"0.6",
				"word": "真丝",
				"tag":"材质"
			},
			{
				"synonym":"",
				"weight":"0.8",
				"word": "韩都衣舍",
				"tag":"品牌"
			},
			{
				"synonym":"连身裙;联衣裙",
				"weight":"1.0",
				"word": "连衣裙",
				"tag":"品类"
			}
		]
}
`
	// interface{} 代表不知道是什么类型
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", m["data"].([]interface{})[2].(map[string]interface{})["synonym"])

}

func parseNLP2() {
	res := `{
		"data": [
			{
				"synonym":"",
				"weight":"0.6",
				"word": "真丝",
				"tag":"材质"
			},
			{
				"synonym":"",
				"weight":"0.8",
				"word": "韩都衣舍",
				"tag":"品牌"
			},
			{
				"synonym":"连身裙;联衣裙",
				"weight":"1.0",
				"word": "连衣裙",
				"tag":"品类"
			}
		]
}
`

	m := struct {
		Data []struct {
			Synonym string `json:"synonym"`
		} `json:"data"`
	}{}

	err := json.Unmarshal([]byte(res), &m)

	if err != nil {
		panic(err)
	}

	fmt.Println(m.Data[2].Synonym)
}
