package snapshot_test

import (
	"github.com/puffinframework/config"
	"github.com/puffinframework/snapshot"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type MySnapshot struct {
	SeqNum int64
	Data   string
}

func TestSnapshotStore(t *testing.T) {
	os.Setenv(config.ENV_VAR_NAME, config.MODE_TEST)
	store := snapshot.NewLeveldbStore()
	assert.NotNil(t, store)
	defer store.MustDestroy()

	mySnapshot1 := &MySnapshot{}
	store.MustLoadSnapshot("MySnapshot", mySnapshot1)
	assert.Equal(t, 0, mySnapshot1.SeqNum)
	assert.Equal(t, "", mySnapshot1.Data)

	mySnapshot1.SeqNum = 1
	mySnapshot1.Data = "data 1"
	store.MustSaveSnapshot("MySnapshot", mySnapshot1)

	mySnapshot2 := &MySnapshot{}
	store.MustLoadSnapshot("MySnapshot", mySnapshot2)
	assert.Equal(t, mySnapshot1, mySnapshot2)
}
