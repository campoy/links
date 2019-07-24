package repository

import (
	"encoding/json"
	nurl "net/url"

	"github.com/dgraph-io/badger"
	"github.com/pkg/errors"
)

type badgerRepo struct {
	*badger.DB
}

// NewDiskRepository returns a new LinkRepository that stores data in
// disk using Badger DB.
func NewDiskRepository(path string) (LinkRepository, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}
	return &badgerRepo{db}, nil
}

func (db badgerRepo) New(url string) (*Link, error) {
	if _, err := nurl.ParseRequestURI(url); err != nil {
		return nil, err
	}

	link := &Link{ID: randomString(), URL: url}
	b, err := json.Marshal(link)
	if err != nil {
		return nil, errors.Wrapf(err, "could not marshal link")
	}

	txn := db.NewTransaction(true)
	if err := txn.Set([]byte(link.ID), b); err != nil {
		return nil, err
	}

	return link, txn.Commit()
}

func (db badgerRepo) get(txn *badger.Txn, id string) (*Link, error) {
	item, err := txn.Get([]byte(id))
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return nil, ErrNoSuchLink
		}
		return nil, err
	}
	var link Link
	err = item.Value(func(data []byte) error {
		return json.Unmarshal(data, &link)
	})
	if err != nil {
		return nil, errors.Wrapf(err, "could not unmarshal link")
	}
	return &link, nil
}

func (db badgerRepo) Get(id string) (*Link, error) {
	txn := db.NewTransaction(false)
	return db.get(txn, id)
}

func (db badgerRepo) CountVisit(id string) error {
	txn := db.NewTransaction(true)

	link, err := db.get(txn, id)
	if err != nil {
		return err
	}
	link.Count++

	b, err := json.Marshal(link)
	if err != nil {
		return errors.Wrapf(err, "could not marshal link")
	}
	if err := txn.Set([]byte(link.ID), b); err != nil {
		return err
	}
	return txn.Commit()
}
