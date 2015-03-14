package snapshot

import (
	"time"
)

type SnapshotStore interface {
	Load()
	Save()
	GetLastEventDt() time.Time
	SetLastEventDt(lastEventDt time.Time)
}
