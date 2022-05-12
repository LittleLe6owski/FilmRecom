package queryBuilder

type RequestParams struct {
	Version  string
	Path     string
	RawQuery map[string]string
}

func (r *RequestParams) CompilePath() string {
	return r.Version + r.Path
}

type QueryBuilder interface {
	Get(RequestParams) ([]byte, error)
}
