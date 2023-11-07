package okele

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

const (
	defaultDBName = "default"
)

type Collection struct {
	Bucket *bbolt.Bucket
}

type Okele struct {
	db *bbolt.DB
}

type M map[string]string

func New() (*Okele, error) {
	dbname := fmt.Sprintf("%s.okele", defaultDBName)
	db, err := bbolt.Open(dbname, 0666, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &Okele{
		db: db,
	}, nil
}

func (o *Okele) CreateCollection(name string) (*Collection, error) {
	tx, err := o.db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	bucket, err := tx.CreateBucketIfNotExists([]byte(name))
	if err != nil {
		return nil, err
	}
	return &Collection{Bucket: bucket}, nil
}

func (o *Okele) Insert(collName string, data M) (uuid.UUID, error) {
	id := uuid.New()
	tx, err := o.db.Begin(true)
	if err != nil {
		return id, err
	}
	defer tx.Rollback()
	bucket, err := tx.CreateBucketIfNotExists([]byte(collName))
	if err != nil {
		return id, err
	}
	for k, v := range data {
		if err := bucket.Put([]byte(k), []byte(v)); err != nil {
			return id, err
		}
	}
	if err := bucket.Put([]byte("id"), []byte(id.String())); err != nil {
		return id, err
	}
	return id, tx.Commit()
}

func (o *Okele) Select(coll string, query M) (M, error) {
	tx, err := o.db.Begin(false)
	if err != nil {
		return nil, err
	}

	bucket := tx.Bucket([]byte(coll))
	if bucket == nil {
		return nil, fmt.Errorf("collection (%s) not found", coll)
	}

	bucket.NextSequence()
}
