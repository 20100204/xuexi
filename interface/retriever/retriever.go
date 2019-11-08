package retriever

type Retriever interface {
	Get(url string) string
}
type Poster interface {
	Post(url string,form map[string]string) string
}
