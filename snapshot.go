package snapshot

import (
	"time"
)

type Snapshot interface {
	Load()
	Save()
	GetLastEventDt() time.Time
	SetLastEventDt(lastEventDt time.Time)
}
