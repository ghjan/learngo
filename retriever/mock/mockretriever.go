package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	return url + " " + r.Contents
}
