package infra

import (
	"io/ioutil"
	"net/http"
)

type Retriever struct{}

func (Retriever) Get(url string) string {
	resp, err := http.Get("https://ipcrystal.com")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}
