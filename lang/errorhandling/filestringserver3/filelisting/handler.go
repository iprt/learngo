package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

const prefix = "/list/"

func FileList(writer http.ResponseWriter, request *http.Request) error {

	if !strings.Contains(request.URL.Path, prefix) {
		// return errors.New("Path must start with /list/")
		return userError(
			"Path must start with " + prefix,
		)
	}

	// /list/fib.txt -> fib.txt
	path := request.URL.Path[len(prefix):]

	file, err := os.Open(path)
	if err != nil {
		// panic(err)
		// http.Error(writer,
		// 	err.Error(),
		// 	http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
