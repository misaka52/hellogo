package prod

import (
	"io/ioutil"
	"net/http"
)

func retriever(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	resBytes, _ := ioutil.ReadAll(resp.Body)

	return string(resBytes)
}

type Retriever struct {
}

func (Retriever) Get(url string) string {
	return retriever(url)
}
