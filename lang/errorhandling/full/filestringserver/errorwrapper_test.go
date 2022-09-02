package main

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errorPanic(writer http.ResponseWriter,
	request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errorUserError(writer http.ResponseWriter,
	request *http.Request) error {
	return testingUserError("user error")
}

func errorNotFound(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrNotExist
}

func errorNoPermission(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrPermission
}

func errorUnknown(writer http.ResponseWriter,
	request *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter,
	request *http.Request) error {
	return nil
}

/**
提取公共的测试参数
*/
var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	// {errorPanic, 500, ""},
	{errorUserError, 502, "user error"},
	{errorNotFound, 404, "Not Found"},
	{errorNoPermission, 403, "Forbidden"},
	{errorUnknown, 500, "Internal Server Error"},
	{noError, 200, ""},
}

/**
方法的测试
*/
func TestErrorWrapper(t *testing.T) {

	for _, tt := range tests {
		f := errorWrapping(tt.h)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet, "https://ipcrystal.com", nil,
		)
		f(response, request)

		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")
		if response.Code != tt.code || body != tt.message {
			t.Errorf("excpect (%d %s) "+
				"got (%d, %s)", tt.code, tt.message,
				response.Code, body)
		}

	}
}

/**
真正启动一个服务器来测试
*/
func TestErrorWrapperInServer(t *testing.T) {

	for _, tt := range tests {
		f := errorWrapping(tt.h)

		server := httptest.NewServer(http.HandlerFunc(f))

		resp, _ := http.Get(server.URL)

		b, _ := ioutil.ReadAll(resp.Body)
		body := strings.Trim(string(b), "\n")
		if resp.StatusCode != tt.code || body != tt.message {
			t.Errorf("excpect (%d %s) "+
				"got (%d, %s)", tt.code, tt.message,
				resp.StatusCode, body)
		}

	}

}
