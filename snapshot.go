package snapshot

import (
	"errors"
	"github.com/puffinframework/config"
	"github.com/syndtr/goleveldb/leveldb"
	leveldbErrors "github.com/syndtr/goleveldb/leveldb/errors"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
)

var (
	ErrOpenDB            error = errors.New("snapshot: couldn't open database")
	ErrCloseDB           error = errors.New("snapshot: couldn't close database")
	ErrGetSnapshot       error = errors.New("snapshot: couldn't get the snapshot from the db")
	ErrPutSnapshot       error = errors.New("snapshot: couldn't put the snapshot from the db")
	ErrUnmarshalSnapshot error = errors.New("snapshot: couldn't unmarshal the snapshot")
	ErrMarshalSnapshot   error = errors.New("snapshot: couldn't marshal the snapshot")
)

type Store interface {
	MustLoadSnapshot(key string, snapshot interface{})
	MustSaveSnapshot(key string, snapshot interface{})
	MustClose()
	MustDestroy()
}

type leveldbStoreConfig struct {
	SnapshotStore struct {
		LeveldbDir string
	}
}

type leveldbStore struct {
	dir string
	db  *leveldb.DB
}

func NewLeveldbStore() Store {
	cfg := &leveldbStoreConfig{}
	config.MustReadConfig(cfg)

	dir := cfg.SnapshotStore.LeveldbDir

	db, err := leveldb.OpenFile(dir, nil)
	if err != nil {
		log.Panic(ErrOpenDB)
	}

	return &leveldbStore{dir: dir, db: db}
}

func (self *leveldbStore) MustLoadSnapshot(key string, snapshot interface{}) {
	value, err := self.db.Get([]byte(key), nil)
	if err != nil && err != leveldbErrors.ErrNotFound {
		log.Panic(ErrGetSnapshot)
	}

	if err = bson.Unmarshal(value, snapshot); err != nil {
		log.Panic(ErrUnmarshalSnapshot)
	}
}

func (self *leveldbStore) MustSaveSnapshot(key string, snapshot interface{}) {
	value, err := bson.Marshal(snapshot)
	if err != nil {
		log.Panic(ErrMarshalSnapshot)
	}

	if err = self.db.Put([]byte(key), value, nil); err != nil {
		log.Panic(ErrPutSnapshot)
	}
}

func (self *leveldbStore) MustClose() {
	self.db.Close()
}

func (self *leveldbStore) MustDestroy() {
	self.db.Close()
	os.RemoveAll(self.dir)

}
