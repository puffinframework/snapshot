package snapshot_test

import (
	"github.com/puffinframework/snapshot"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type MySnapshot struct {
	seqNum int64
	data   string
}

func TestSnapshotStore(t *testing.T) {
	os.Setenv("PF_MODE", "test")
	store := snapshot.NewLeveldbStore()
	assert.NotNil(t, store)
	defer store.MustDestroy()

	mySnapshot1 := &MySnapshot{}
	store.MustLoadSnapshot("MySnapshot", mySnapshot1)
	assert.Equal(t, 0, mySnapshot1.seqNum)
	assert.Equal(t, "", mySnapshot1.data)

	mySnapshot1.seqNum = 1
	mySnapshot1.data = "data 1"
	store.MustSaveSnapshot("MySnapshot", mySnapshot1)
}
