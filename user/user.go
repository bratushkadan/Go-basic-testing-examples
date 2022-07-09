package user

type APIClient interface {
	GetData(query string) (*Response, error)
}

type Response struct {
	Text       string
	StatusCode uint16
}
