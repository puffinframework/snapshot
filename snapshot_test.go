package snapshot_test

import (
	"github.com/puffinframework/snapshot"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSnapshotStore(t *testing.T) {
	os.Setenv("PF_MODE", "test")
	store := snapshot.NewLeveldbStore()
	assert.NotNil(t, store)
	defer store.MustDestroy()
}
