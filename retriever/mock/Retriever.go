package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	return r.Contents
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

// 类似于实现java的toString()方法
func (r Retriever) String() string {
	return fmt.Sprintf("MockRetriever:{content=%s}", r.Contents)
}
