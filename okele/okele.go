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
	bucket *bbolt.Bucket
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
	coll := Collection{}
	err := o.db.Update(func(tx *bbolt.Tx) error {
		var (
			err    error
			bucket *bbolt.Bucket
		)
		bucket = tx.Bucket([]byte(name))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(name))
			if err != nil {
				return err
			}
		}

		coll.bucket = bucket
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &coll, nil
}

func (o *Okele) Insert(collName string, data M) (uuid.UUID, error) {
	id := uuid.New()
	coll, err := o.CreateCollection(collName)
	if err != nil {
		return id, err
	}

	for k, v := range data {
		if err := coll.bucket.Put([]byte(k), []byte(v)); err != nil {
			return id, err
		}
	}

	if err := coll.bucket.Put([]byte("id"), []byte(id.String())); err != nil {
		return id, err
	}

	return id, nil
	//coll.bucket.Put([]byte(id), []byte(coll))
}

// func (o *Okele) Get(coll *string, k string, query any) {

// }

// db.Update(func(tx *bbolt.Tx) error {
// 	bucket, err := tx.CreateBucket([]byte("users"))
// 	if err != nil {
// 		return err
// 	}
// 	id := uuid.New()

// 	for k, v := range user {
// 		if err := bucket.Put([]byte(k), []byte(v)); err != nil {
// 			return err
// 		}
// 	}

// 	if err := bucket.Put([]byte("id"), []byte(id.String())); err != nil {
// 		return err
// 	}

// 	return nil
// })
