package main

import (
	"iproute.org/learngo/lang/errorhandling/filestringserver3/filelisting"
	"log"
	"net/http"
	// 要加下划线
	_ "net/http/pprof"
	"os"
)

/**
第一步 type 一个 function 包装原来的function

第二步 原来的 function 错误直接返回
*/
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 第三步，统一的错误处理
func errorWrapping(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		// panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		// 业务处理
		err := handler(writer, request)
		log.Println("Error occurred ")
		// 统一的错误处理
		if err != nil {
			code := http.StatusOK

			// user error
			if userError, ok := err.(userError); ok {
				http.Error(writer,
					userError.Message(),
					http.StatusBadGateway)
				return
			}

			// system error
			switch {
			case os.IsNotExist(err):
				// http.Error(writer, http.StatusText(http.StatusNotFound),
				// 	http.StatusNotFound)
				code = http.StatusNotFound
			case os.IsPermission(err):
				// 权限错误
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code),
				code)
		}

	}
}

/**
go 语言的接口 各自管各自的定义

*/
type userError interface {
	error
	Message() string
}

func main() {

	// http.HandleFunc("/list/", filelisting.FileList)
	http.HandleFunc("/list/", errorWrapping(filelisting.FileList))

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}
}
