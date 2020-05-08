package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type RetrieverPointer struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *RetrieverPointer) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)
	err = resp.Body.Close()
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	return string(result)
}
