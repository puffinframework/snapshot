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

	mySnapshot := &MySnapshot{}
	store.MustLoadSnapshot("MySnapshot", mySnapshot)
	assert.Equal(t, 0, mySnapshot.seqNum)
	assert.Equal(t, "", mySnapshot.data)
}
