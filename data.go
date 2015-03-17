package snapshot

import (
	"time"
)

type Data interface {
	LoadFrom(ss Store) error
	SaveTo(ss Store) error
	GetLastEventDt() (time.Time, error)
	SetLastEventDt(lastEventDt time.Time) error
}
