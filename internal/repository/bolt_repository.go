package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kalunik/tgBot-urShorty/internal/entity"
	bolt "go.etcd.io/bbolt"
)

type EmptyDatabaseError struct {
	err string
}

func (e *EmptyDatabaseError) Error() string {
	return e.err
}

type BoltRepository interface {
	Set(collection entity.Collection) error
	Get(chatID string) (entity.Collection, error)
}

type boltRepo struct {
	db *bolt.DB
}

func NewBoltRepository(boltDB *bolt.DB) BoltRepository {
	return &boltRepo{db: boltDB}
}

func (b boltRepo) Set(collection entity.Collection) error {
	if err := b.db.Update(func(tx *bolt.Tx) error {
		bytes, err := json.Marshal(collection.MetaPath)
		if err != nil {
			return errors.New("failed to marshal collection")
		}
		if err = tx.Bucket([]byte("collection")).Put([]byte(collection.ChatID), bytes); err != nil {
			return fmt.Errorf("bucket put failure: %w", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("set transaction failure: %w", err)
	}
	return nil
}

func (b boltRepo) Get(chatID string) (entity.Collection, error) {
	paths := entity.Collection{ChatID: chatID}
	if err := b.db.View(func(tx *bolt.Tx) error {
		bytes := tx.Bucket([]byte("collection")).Get([]byte(chatID))
		if bytes == nil {
			return &EmptyDatabaseError{err: "we can't find any data for you"}
		}
		err := json.Unmarshal(bytes, &paths.MetaPath)
		if err != nil {
			return fmt.Errorf("failed to unmarshal bytes: %w", err)
		}
		return nil
	}); err != nil {
		return entity.Collection{}, err
	}
	return paths, nil
}
