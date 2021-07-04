package main

import (
	"fmt"
	"hellogo/errorhandling/fileserver/filelisting"
	"log"
	"net/http"

	_ "net/http/pprof"
	"os"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter,
	*http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				fmt.Printf("panic:%v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if userErr, ok := err.(userError); ok {
			log.Println("userError occurred:", userErr)
			http.Error(writer, userErr.Message(), http.StatusBadRequest)
			return
		}

		code := http.StatusOK
		if err != nil {
			log.Printf("handle error。err: %s", err)
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
		}
		http.Error(writer, http.StatusText(code), code)
	}
}

type userError interface {
	error
	Message() string
}

/**
开启了一个文件服务器，指定tcp端口
内含异常处理，结合：error+panic+recover
*/
func main() {
	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileListing))

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}
}
