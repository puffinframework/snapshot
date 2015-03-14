package snapshot

import (
	"log"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"labix.org/v2/mgo/bson"
)

type storeLeveldb struct {
	dir string
	db  *leveldb.DB
}

func NewLeveldbStore(dir string) Store {
	db, err := leveldb.OpenFile(dir, nil)
	if err != nil {
		log.Println(err)
		log.Panic(ErrOpenStore)
	}

	return &storeLeveldb{dir: dir, db: db}
}

func (self *storeLeveldb) MustLoadSnapshot(key string, snapshot interface{}) {
	value, err := self.db.Get([]byte(key), nil)
	if err != nil {
		if err == errors.ErrNotFound {
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

func (self *storeLeveldb) MustSaveSnapshot(key string, snapshot interface{}) {
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

func (self *storeLeveldb) MustClose() {
	if err := self.db.Close(); err != nil {
		log.Println(err)
		log.Panic(ErrCloseStore)
	}
}

func (self *storeLeveldb) MustDestroy() {
	self.MustClose()
	if err := os.RemoveAll(self.dir); err != nil {
		log.Println(err)
		log.Panic(ErrDestroyStore)
	}
}
