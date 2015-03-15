package snapshot

import (
	"time"
)

type Data interface {
	LoadFrom(ss Store)
	SaveTo(ss Store)
	GetLastEventDt() time.Time
	SetLastEventDt(lastEventDt time.Time)
}
