package main

import (
	"fmt"
	"hellogo/retriever/mock"
	"hellogo/retriever/real"
	"time"
)

type retrieverIntf interface {
	Get(string) string
}

type Poster interface {
	Post(url string,
		from map[string]string) string
}

func download(r retrieverIntf) string {
	return r.Get("http://www.baidu.com")
}

const url = "www.baidu.com"

func post(p Poster) {
	p.Post(url,
		map[string]string{
			"contents": "ccmouse",
			"course":   "golang",
		})
}

type RetrieverPoster interface {
	retrieverIntf
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another fake imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r retrieverIntf
	mockRetriever := mock.Retriever{Contents: "this is mock retriever"}
	r = &mockRetriever
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println("This is mock retriever. content=", mockRetriever.Contents)
	} else {
		fmt.Println("This is not mock retriever")
	}

	fmt.Println("before ", mockRetriever.Contents)
	session(&mockRetriever)
	fmt.Println("after ", mockRetriever.Contents)

	//fmt.Println(download(r))
}

func inspect(r retrieverIntf) {
	fmt.Println("inspecting:", r)
	// %T：打印类型，%v：打印值
	fmt.Printf("type:%T value:%v\n", r, r)
	fmt.Print(" > Type switch: ")
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("mock retriever, ", v.Contents)
	case *real.Retriever:
		fmt.Println("real retriever, ", v.UserAgent)
	default:
		fmt.Println("other type retriever")
	}
	fmt.Println()
}
