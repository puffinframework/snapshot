package snapshot

import(
	"errors"
)

var (
	ErrOpenStore         error = errors.New("snapshot: couldn't open the store")
	ErrCloseStore        error = errors.New("snapshot: couldn't close the store")
	ErrDestroyStore      error = errors.New("snapshot: couldn't destroy the store")
	ErrGetSnapshot       error = errors.New("snapshot: couldn't get the snapshot from the store")
	ErrPutSnapshot       error = errors.New("snapshot: couldn't put the snapshot from the store")
	ErrMarshalSnapshot   error = errors.New("snapshot: couldn't marshal the snapshot")
	ErrUnmarshalSnapshot error = errors.New("snapshot: couldn't unmarshal the snapshot")
)
