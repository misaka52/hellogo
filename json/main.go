package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    int32   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	ID int32 `json:"id"`
	// 忽略空值
	Name       string       `json:"name,omitempty"`
	TotalPrice float64      `json:"total_price"`
	OrderItems []*OrderItem `json:"order_items"`
}

func main() {
	//json1()
	parseNLP()
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
	// 结构体打印
	vs := struct {
		Data []struct {
			Synonym string `json:"synonym"`
			Tag     string `json:"tag"`
		} `json:"data"`
	}{}
	_ = json.Unmarshal([]byte(res), &vs)
	fmt.Println(vs.Data[2].Synonym, vs.Data[2].Tag)

	// map接收打印
	var mapVs map[string]interface{}
	_ = json.Unmarshal([]byte(res), &mapVs)
	fmt.Println(mapVs["data"].([]interface{})[2].(map[string]interface{})["synonym"])

}

func json1() {
	o := Order{
		ID:         1,
		Name:       "o1",
		TotalPrice: 30,
		OrderItems: []*OrderItem{
			{
				ID:    101,
				Name:  "it1",
				Price: 20,
			},
			{
				ID:    102,
				Name:  "it2",
				Price: 15,
			},
		},
	}
	marshal, err := json.Marshal(o)
	if err != err {
		panic(err)
	}
	fmt.Printf("%s\n", marshal)

	var newOrder Order
	_ = json.Unmarshal(marshal, &newOrder)

	fmt.Printf("%+v\n", newOrder)
}
