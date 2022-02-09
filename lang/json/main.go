package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
type Order struct { //tag
	ID         string       `json:"id"`
	Items      []*OrderItem `json:"items"`
	TotalPrice float64      `json:"total_price"`
}

func main() {
	o := Order{
		ID: "1234",
		Items: []*OrderItem{
			{ID: "item_1",
				Name:  "learn go",
				Price: 15,
			},
			{
				ID:    "item_2",
				Name:  "interview",
				Price: 10,
			},
		},
		TotalPrice: 30,
	}
	fmt.Printf("%+v\n", o)
	fmt.Printf("%#v\n", o)
	if oContent, err := json.Marshal(o); err == nil && oContent != nil {
		fmt.Printf("%s\n", oContent)
		unmarshalOrder(string(oContent))
	}

	parseNLP()
}

func unmarshalOrder(s string) {
	//s := `{"id":"1234","Items":[{"id":"item_1","name":"learn go","price":15},{"id":"item_2","name":"interview","price":10}],"total_price":30}`
	var o Order
	if err := json.Unmarshal([]byte(s), &o); err == nil {
		fmt.Printf("%#v\n", o)
	}
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
}`
	//匿名struct类型
	m := struct {
		Data []struct {
			Synonym string `json:"synonym"`
			Weight  string `json:"weight"`
			Word    string `json:"word"`
			Tag     string `json:"tag"`
		} `json:"data"`
	}{}
	err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", m)

	if oContent, err := json.Marshal(m); err == nil && oContent != nil {
		fmt.Printf("%s\n", oContent)
	}

	fmt.Printf("%+v, %+v\n", m.Data[2].Synonym, m.Data[2].Tag)

}
