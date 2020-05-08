package mock

type Retriever2 struct {
	contexts string
	name     string
}

func (r *Retriever2) Post(url string, m map[string]string) string {

	r.contexts = url
	r.name = m["name"]
	return "ok"
}

func (r *Retriever2) Get(url string) string {
	return r.contexts + "  haha  " + url
}
