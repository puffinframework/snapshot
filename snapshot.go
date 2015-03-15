package snapshot

import (
	"time"
)

type Snapshot interface {
	LoadFrom(ss Store)
	SaveTo(ss Store)
	GetLastEventDt() time.Time
	SetLastEventDt(lastEventDt time.Time)
}
