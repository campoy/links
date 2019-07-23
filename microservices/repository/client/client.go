package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/campoy/links/microservices/repository"
)

// New creates a LinkRepository that fetches client information.
func New(addr string) repository.LinkRepository {
	return client{addr}
}

type client struct{ addr string }

func (c client) do(method string, path string, body []byte, code int) (*http.Response, error) {
	req, err := http.NewRequest(method, c.addr+path, bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrapf(err, "could not create request")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "could not send request")
	}
	if res.StatusCode != code {
		return nil, errors.Errorf("got status %s", res.Status)
	}
	return res, nil
}

func (c client) New(url string) (*repository.Link, error) {
	body, err := json.Marshal(struct {
		URL string `json:"url"`
	}{url})
	if err != nil {
		return nil, errors.Wrapf(err, "could not marshal request")
	}

	res, err := c.do(http.MethodPost, "/link/", body, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var link repository.Link
	if err := json.NewDecoder(res.Body).Decode(&link); err != nil {
		return nil, errors.Wrapf(err, "could not decode link")
	}
	return &link, nil
}

func (c client) Get(id string) (*repository.Link, error) {
	res, err := c.do(http.MethodGet, "/link/"+id, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var link repository.Link
	if err := json.NewDecoder(res.Body).Decode(&link); err != nil {
		return nil, errors.Wrapf(err, "could not decode link")
	}
	return &link, nil
}

func (c client) CountVisit(id string) error {
	_, err := c.do(http.MethodPost, "/link/"+id, nil, http.StatusCreated)
	return err
}
