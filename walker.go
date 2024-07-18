package afs

import (
	"context"
	"errors"
	"github.com/knights-analytics/afs/file"
	"github.com/knights-analytics/afs/storage"
	"github.com/knights-analytics/afs/url"
	"github.com/knights-analytics/afs/walker"
)

// Walk visits all location recursively within provided sourceURL
func (s *service) Walk(ctx context.Context, URL string, handler storage.OnVisit, options ...storage.Option) error {
	if URL == "" {
		return errors.New("URL was empty")
	}
	URL = url.Normalize(URL, file.Scheme)
	manager, err := s.manager(ctx, URL, options)
	if err != nil {
		return err
	}
	URL = url.Normalize(URL, file.Scheme)
	managerWalker, ok := manager.(storage.Walker)
	if ok {
		return managerWalker.Walk(ctx, URL, handler, options...)
	}
	managerWalker = walker.New(manager)
	return managerWalker.Walk(ctx, URL, handler, options...)
}
