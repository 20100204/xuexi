package real

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r Retriever) String() string {
	 return fmt.Sprintf("Retriever:{Content=%s}",r.UserAgent)
}

func (r Retriever) Post(url string, form map[string]string) string {
	//panic("implement me")
	return url;
}

func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	res, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(res)
}





