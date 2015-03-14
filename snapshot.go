package snapshot

import (
	"github.com/puffinframework/config"
	"github.com/syndtr/goleveldb/leveldb"
	leveldbErrors "github.com/syndtr/goleveldb/leveldb/errors"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
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
		log.Println(err)
		log.Panic(ErrOpenStore)
	}

	return &leveldbStore{dir: dir, db: db}
}

func (self *leveldbStore) MustLoadSnapshot(key string, snapshot interface{}) {
	value, err := self.db.Get([]byte(key), nil)
	if err != nil {
		if err == leveldbErrors.ErrNotFound {
			return
		} else {
			log.Println(err)
			log.Panic(ErrGetSnapshot)
		}
	}

	if err = bson.Unmarshal(value, snapshot); err != nil {
		log.Println(err)
		log.Panic(ErrUnmarshalSnapshot)
	}
}

func (self *leveldbStore) MustSaveSnapshot(key string, snapshot interface{}) {
	value, err := bson.Marshal(snapshot)
	if err != nil {
		log.Println(err)
		log.Panic(ErrMarshalSnapshot)
	}

	if err = self.db.Put([]byte(key), value, nil); err != nil {
		log.Println(err)
		log.Panic(ErrPutSnapshot)
	}
}

func (self *leveldbStore) MustClose() {
	if err := self.db.Close(); err != nil {
		log.Println(err)
		log.Panic(ErrCloseStore)
	}
}

func (self *leveldbStore) MustDestroy() {
	self.MustClose()
	if err := os.RemoveAll(self.dir); err != nil {
		log.Println(err)
		log.Panic(ErrDestroyStore)
	}
}
