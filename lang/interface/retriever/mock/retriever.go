package mock

type Retriever struct {
	Content string
}

// 只要实现了 Get方法 就可以认为是 duck type
func (r Retriever) Get(url string) string {
	return r.Content
}
