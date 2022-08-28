package really

import (
	"io/ioutil"
	"net/http"
)

type HttpRetriever struct {
}

func (r HttpRetriever) Get(url string) string {

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	return string(bytes)
}
