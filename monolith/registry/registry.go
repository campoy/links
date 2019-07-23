package registry

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
)

// A Link contains the information related to a shorten link.
type Link struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Count int    `json:"count"`
}

// ErrNoSuchLink is returned when a given Link does not exist.
var ErrNoSuchLink = errors.New("no such link")

var links = make(map[string]Link)

func randomString() string {
	return fmt.Sprintf("%X", rand.Int63())
}

func NewLink(u string) (*Link, error) {
	if _, err := url.ParseRequestURI(u); err != nil {
		return nil, err
	}
	l := Link{
		ID:  randomString(),
		URL: u,
	}
	links[l.ID] = l
	return &l, nil
}

func GetLink(id string) (*Link, error) {
	l, ok := links[id]
	if !ok {
		return nil, ErrNoSuchLink
	}
	return &l, nil
}

func RecordVisit(id string) error {
	l, ok := links[id]
	if !ok {
		return ErrNoSuchLink
	}
	l.Count++
	links[id] = l
	return nil
}
