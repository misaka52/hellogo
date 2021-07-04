package main

import (
	"fmt"
	"hellogo/interf/prod"
)

func getRetriever() retriever {
	return prod.Retriever{}
}

type retriever interface {
	Get(string) string
}

func main() {
	var r retriever = getRetriever()
	fmt.Println(r.Get("http://www.baidu.com"))
}
