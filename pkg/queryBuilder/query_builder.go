package queryBuilder

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

/*
	Небольшая обёртка для удобного использования http запросов.
*/

type QueryConstructor struct {
	headers map[string]string
	schema  string
	host    string
}

func (q *QueryConstructor) AddHeader(key, value string) {
	q.headers[key] = value
}

func NewQueryBuilder(schema, host string, headers map[string]string) QueryBuilder {
	return &QueryConstructor{
		headers: headers,
		schema:  schema,
		host:    host,
	}
}

func (q *QueryConstructor) Get(pReq RequestParams) (body []byte, err error) {
	client := http.Client{}
	header := http.Header{}
	for k, v := range q.headers {
		header.Add(k, v)
	}

	queryParams := url.Values{}
	for query := range pReq.RawQuery {
		queryParams.Add(query, pReq.RawQuery[query])
	}

	reqUrl := url.URL{
		Scheme:   q.schema,
		Host:     q.host,
		Path:     pReq.CompilePath(),
		RawQuery: queryParams.Encode()}
	req := http.Request{Method: http.MethodGet, URL: &reqUrl, Header: header}

	log.Printf("Do sends an HTTP request by url %v", reqUrl)
	res, err := client.Do(&req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}

	if res.StatusCode > 299 {
		log.Printf("Request completed unsuccessfully, error code = %d", res.StatusCode)
		return
	}
	return
}
