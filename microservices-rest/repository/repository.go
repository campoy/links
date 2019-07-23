package repository

import "errors"

// A LinkRepository provides all the operations to create and store links.
type LinkRepository interface {
	New(url string) (*Link, error)
	Get(id string) (*Link, error)
	CountVisit(id string) error
}

// A Link contains the information related to a shorten link.
type Link struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Count int    `json:"count"`
}

// ErrNoSuchLink is returned when a given Link does not exist.
var ErrNoSuchLink = errors.New("no such link")
