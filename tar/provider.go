package tar

import "github.com/knights-analytics/afs/storage"

// Provider returns a http manager
func Provider(options ...storage.Option) (storage.Manager, error) {
	return New(options...), nil
}
