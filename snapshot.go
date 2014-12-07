package snapshot

import (
	"errors"
	"github.com/puffinframework/config"
	"labix.org/v2/mgo/bson"
)

var (
	ErrOpenDB            error = errors.New("snapshot: couldn't open database")
	ErrGetSnapshot       error = errors.New("snapshot: couldn't get the snapshot from the db")
	ErrPutSnapshot       error = errors.New("snapshot: couldn't put the snapshot from the db")
	ErrUnmarshalSnapshot error = errors.New("snapshot: couldn't unmarshal the snapshot")
	ErrMarshalSnapshot   error = errors.New("snapshot: couldn't marshal the snapshot")
)

type Store interface {
	LoadSnapshot(key string, snapshot interface{})
	SaveSnapshot(key string, snapshot interface{})
	Close()
}

type leveldbStoreConfig struct {
	LeveldbDir string
}

type leveldbStore struct {
	db *leveldb.DB
}

func NewLeveldbStore() *Store {
	cfg := &leveldbStoreConfig{}
	config.MustReadConfig(cfg)

	db, err := leveldb.OpenFile(cfg.LeveldbDir, nil)
	if err != nil {
		log.Panic(ErrOpenDB)
	}

	return &leveldbStore{db: db}
}

func (self *leveldbStore) LoadSnapshot(key string, snapshot interface{}) error {
	value, err := self.db.Get([]byte(key), nil)
	if err != nil {
		return ErrGetSnapshot
	}

	if err = bson.Unmarshal(value, snapshot); err != nil {
		return ErrUnmarshalSnapshot
	}

	return nil
}

func (self *leveldbStore) SaveSnapshot(key string, snapshot interface{}) error {
	value, err := bson.Marshal(snapshot)
	if err != nil {
		return ErrMarshalSnapshot
	}

	if err = db.Put([]byte(string), value, nil); err != nil {
		return ErrPutSnapshot
	}

	return nil
}

func (self *leveldbStore) Close() {
	self.db.Close()
}
