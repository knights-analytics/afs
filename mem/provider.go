package mem

import (
	"github.com/knights-analytics/afs/storage"
)

// Provider manager provider function
func Provider(options ...storage.Option) (storage.Manager, error) {
	return Singleton(options...), nil
}
