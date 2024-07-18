package file

import (
	"github.com/knights-analytics/afs/storage"
)

// Provider provider function
func Provider(options ...storage.Option) (storage.Manager, error) {
	return New(), nil
}
