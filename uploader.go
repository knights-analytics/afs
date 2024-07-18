package afs

import (
	"context"
	"github.com/knights-analytics/afs/base"
	"github.com/knights-analytics/afs/file"
	"github.com/knights-analytics/afs/storage"
	"github.com/knights-analytics/afs/url"
	"io"
)

// UploadInBatch default implementation for UploadInBatch
func (s *service) Uploader(ctx context.Context, URL string, options ...storage.Option) (storage.Upload, io.Closer, error) {
	URL = url.Normalize(URL, file.Scheme)
	manager, err := s.manager(ctx, URL, options)
	if err != nil {
		return nil, nil, err
	}
	batchUploader, ok := manager.(storage.BatchUploader)
	if !ok {
		batchUploader = base.NewUploader(manager)
	}
	return batchUploader.Uploader(ctx, URL, options...)
}
