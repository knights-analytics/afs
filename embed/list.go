package embed

import (
	"context"
	"fmt"
	"github.com/knights-analytics/afs/file"
	"github.com/knights-analytics/afs/object"
	"github.com/knights-analytics/afs/option"
	"github.com/knights-analytics/afs/storage"
	"github.com/knights-analytics/afs/url"
	"strings"
)

func (s *manager) List(ctx context.Context, URL string, options ...storage.Option) ([]storage.Object, error) {
	if s.err != nil {
		return nil, s.err
	}
	baseURL, filePath := url.Base(URL, Scheme)
	fPath := file.Path(filePath)
	fPath = strings.Trim(fPath, "/")
	fh, err := s.fs.Open(fPath)
	if err != nil {
		return nil, fmt.Errorf("unable to open '%v', %w", fPath, err)
	}
	match, page := option.GetListOptions(options)
	defer func() { _ = fh.Close() }()
	stat, err := fh.Stat()
	if err != nil {
		return nil, err
	}
	if !stat.IsDir() {
		return []storage.Object{
			object.New(URL, stat, nil),
		}, nil
	}
	files, err := s.fs.ReadDir(fPath)
	if err != nil {
		return nil, err
	}
	var result = make([]storage.Object, 0)
	result = append(result, object.New(URL, stat, nil))
	for _, fileInfo := range files {
		info, err := fileInfo.Info()
		if err != nil {
			return nil, fmt.Errorf("failed to get info for: %v, %w", fileInfo.Name(), err)
		}
		if !match(filePath, info) {
			continue
		}
		page.Increment()
		if page.ShallSkip() {
			continue
		}
		fileURL := url.Join(baseURL, filePath, fileInfo.Name())
		result = append(result, object.New(fileURL, info, nil))
		if page.HasReachedLimit() {
			break
		}
	}
	return result, nil
}
