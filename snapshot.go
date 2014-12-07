package snapshot

import (
	"errors"
	"github.com/puffinframework/config"
)

var (
	ErrOpenDB error = errors.New("snapshot: couldnt open database")
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

func (self *leveldbStore) LoadSnapshot(key string, snapshot interface{}) {
}

func (self *leveldbStore) SaveSnapshot(key string, snapshot interface{}) {
}

func (self *leveldbStore) Close() {
	self.db.Close()
}
