package boltdb

import (
	"fmt"

	"github.com/boltdb/bolt"
)

func GetPathsUrls(PathsUrls *map[string]string) error {

	db, err := setUpDB()
	if err != nil {
		return err
	}

	err = putUrlsInDB(db)
	if err != nil {
		return err
	}

	tx, err := db.Begin(false)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	pu := make(map[string]string)

	b := tx.Bucket([]byte("DB")).Bucket([]byte("PATHS_URLS"))
	err = b.ForEach(func(k, v []byte) error {
		//fmt.Println(string(k), string(v))
		pu[string(k)] = string(v)
		return nil
	})
	if err != nil {
		return err
	}

	*PathsUrls = pu
	return nil
}

func setUpDB() (*bolt.DB, error) {
	db, err := bolt.Open("database.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	db.Update(func(t *bolt.Tx) error {
		root, err := t.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}

		_, err = root.CreateBucketIfNotExists([]byte("PATHS_URLS"))
		if err != nil {
			return fmt.Errorf("could not create weight bucket: %v", err)
		}

		return nil
	})

	return db, nil
}

func putUrlsInDB(db *bolt.DB) error {
	err := db.Update(func(t *bolt.Tx) error {

		err := t.Bucket([]byte("DB")).Bucket([]byte("PATHS_URLS")).Put([]byte("/urlshort-godoc"), []byte("https://godoc.org/github.com/gophercises/urlshort"))
		if err != nil {
			return err
		}

		err = t.Bucket([]byte("DB")).Bucket([]byte("PATHS_URLS")).Put([]byte("/yaml-godoc"), []byte("https://godoc.org/gopkg.in/yaml.v2"))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
