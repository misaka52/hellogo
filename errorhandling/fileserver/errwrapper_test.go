package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(_ http.ResponseWriter, _ *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(_ http.ResponseWriter,
	_ *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(_ http.ResponseWriter,
	_ *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(_ http.ResponseWriter,
	_ *http.Request) error {
	return os.ErrPermission
}

func errUnknown(_ http.ResponseWriter,
	_ *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter,
	_ *http.Request) error {
	//fmt.Fprint(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "OK"},
}

// 单元测试，直接掉接口
func TestErrWrapper(t *testing.T) {

	for _, tt := range tests {
		wrapper := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet,
			"http://www.baidu.com",
			nil)

		wrapper(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}

}

// 服务器测试，启动一个服务器
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		wrapper := errWrapper(tt.h)
		// 模拟服务器
		server := httptest.NewServer(http.HandlerFunc(wrapper))
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectCode int, expectMsg string, t *testing.T) {
	bodyByte, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(bodyByte), "\n")
	if resp.StatusCode != expectCode || body != expectMsg {
		t.Errorf("get %d \"%s\", but expect %d \"%s\"",
			resp.StatusCode, body, expectCode, expectMsg)
	}
}
