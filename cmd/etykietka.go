package cmd

import (
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

func Root() {
	opts := badger.DefaultOptions("/tmp/badger")
	opts.Logger = nil
	db, err := badger.Open(opts)

	if err != nil {
		log.Fatal(err)
	}

	setTransaction(db)
	getTransaction(db)
	defer db.Close()
}

func setTransaction(db *badger.DB) error {
	txn := db.NewTransaction(true)
	defer txn.Discard()

	err := txn.Set([]byte("path"), []byte("tag"))

	if err != nil {
		log.Fatal(err)
	}

	if err := txn.Commit(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func getTransaction(db *badger.DB) error {
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("path"))

		if err != nil {
			log.Fatal(err)
		}

		err = item.Value(func(val []byte) error {
			fmt.Printf("%s\n", val)
			return nil
		})

		if err != nil {
			log.Fatal(err)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
