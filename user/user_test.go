package user

import (
	"errors"
	"testing"
)

type MockAPIClient interface {
	APIClient
	SetResponse(resp Response, err error)
}

type Mock struct {
	Resp Response
	Err  error
}

func (m *Mock) GetData(query string) (Response, error) {
	return m.Resp, m.Err
}

func (m *Mock) SetResponse(resp Response, err error) {
	m.Resp = resp
	m.Err = err
}

var EmptyQueryError = errors.New("empty query")

func GetResponse(client MockAPIClient, query string) (*Response, error) {
	if query == "" {
		return nil, EmptyQueryError
	}

	resp := &Response{
		Text:       "Woo!",
		StatusCode: 200,
	}
	return resp, nil
}

func TestUserApiClient(t *testing.T) {
	table := []struct {
		query string
		resp  Response
		err   error
	}{
		{"success", Response{"Success", 200}, nil},
		{"error", Response{"", 500}, errors.New("Internal Server Error.")},
		{"empty", Response{"", 204}, nil},
	}

	for _, item := range table {
		m := &Mock{}
		m.SetResponse(item.resp, item.err)

		resp, err := m.GetData(item.query)
		if resp.Text != item.resp.Text {
			t.Errorf(`%s: want %v got %v`, item.query, item.resp.Text, resp.Text)
		}
		if !errors.Is(err, item.err) {
			t.Errorf(`%s: want %v got %v`, item.query, item.err, err)
		}
	}
}
