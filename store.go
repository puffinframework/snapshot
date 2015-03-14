package snapshot

type Store interface {
	MustLoadSnapshot(key string, snapshot interface{})
	MustSaveSnapshot(key string, snapshot interface{})
	MustClose()
	MustDestroy()
}
