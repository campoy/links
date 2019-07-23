package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
)

type link struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Count int    `json:"count"`
}

var errNoSuchLink = errors.New("no such link")

var links = make(map[string]link)

func randomString() string {
	return fmt.Sprintf("%X", rand.Int63())
}

func newLink(u string) (*link, error) {
	if _, err := url.ParseRequestURI(u); err != nil {
		return nil, err
	}
	l := link{
		ID:  randomString(),
		URL: u,
	}
	links[l.ID] = l
	return &l, nil
}

func getLink(id string) (*link, error) {
	l, ok := links[id]
	if !ok {
		return nil, errNoSuchLink
	}
	return &l, nil
}

func recordVisit(id string) error {
	l, ok := links[id]
	if !ok {
		return errNoSuchLink
	}
	l.Count++
	links[id] = l
	return nil
}
