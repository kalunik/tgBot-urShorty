package db

import (
	"fmt"
	"github.com/kalunik/tgBot-urShorty/config"
	bolt "go.etcd.io/bbolt"
)

func NewBoltStorage(config config.AppConfig) (*bolt.DB, error) {
	database, err := bolt.Open(config.Bolt.DbPath, config.Bolt.Mode, nil)
	if err != nil {
		return nil, fmt.Errorf("couldn't open db: %w", err)
	}

	if err = initBuckets(database); err != nil {
		return nil, err
	}

	return database, nil
}

func initBuckets(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("collection"))
		if err != nil {
			return fmt.Errorf("couldn't create a collection bucket: %w", err)
		}
		return nil
	})
}
