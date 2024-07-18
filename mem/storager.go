package mem

import (
	"github.com/knights-analytics/afs/base"
	"github.com/knights-analytics/afs/file"
	"github.com/knights-analytics/afs/storage"
	"github.com/knights-analytics/afs/url"
	"sync"
)

type storager struct {
	base.Storager
	scheme string
	Root   *Folder
	mux    sync.Mutex
}

func (s *storager) Close() error {
	return nil
}

// NewStorager create a new in memeory storage service
func NewStorager(baseURL string) storage.Storager {
	baseURL, _ = url.Base(baseURL, Scheme)
	result := &storager{
		Root:   NewFolder(baseURL, file.DefaultDirOsMode),
		scheme: url.Scheme(baseURL, Scheme),
	}
	result.Storager.List = result.List
	return result
}
