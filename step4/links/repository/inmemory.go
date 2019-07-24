package repository

import (
	"fmt"
	"math/rand"
	nurl "net/url"
)

type inmem struct {
	links map[string]*Link
}

func randomString() string { return fmt.Sprintf("%X", rand.Int63()) }

// NewInMemory returns a new in-memory Link
func NewInMemory() LinkRepository {
	return &inmem{make(map[string]*Link)}
}

func (r *inmem) New(url string) (*Link, error) {
	if _, err := nurl.ParseRequestURI(url); err != nil {
		return nil, err
	}
	l := &Link{
		ID:  randomString(),
		URL: url,
	}
	r.links[l.ID] = l
	return l, nil
}

func (r *inmem) Get(id string) (*Link, error) {
	l, ok := r.links[id]
	if !ok {
		return nil, ErrNoSuchLink
	}
	return l, nil
}

func (r *inmem) CountVisit(id string) error {
	l, ok := r.links[id]
	if !ok {
		return ErrNoSuchLink
	}
	l.Count++
	return nil
}
