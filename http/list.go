package http

import (
	"context"
	"fmt"
	"github.com/knights-analytics/afs/file"
	"github.com/knights-analytics/afs/object"
	"github.com/knights-analytics/afs/storage"
	"github.com/knights-analytics/afs/url"
	"net/http"
	"path"
	"time"
)

const lastModifiedHeader = "Last-Modified"

var assetMode, _ = file.NewMode("-rw-r--r--")

func (s *manager) List(ctx context.Context, URL string, options ...storage.Option) ([]storage.Object, error) {
	request, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.run(ctx, URL, request, options...)
	if err != nil {
		return nil, err
	}
	defer s.closeResponse(response)
	if !IsStatusOK(response) {
		return nil, fmt.Errorf("resource not found, statusCode: %v, url: %v", response.StatusCode, URL)
	}
	_, URLPath := url.Base(URL, Scheme)
	_, name := path.Split(URLPath)
	modified := HeaderTime(response.Header, lastModifiedHeader, time.Now())
	info := file.NewInfo(name, response.ContentLength, assetMode, modified, false)
	asset := object.New(URL, info, response)
	return []storage.Object{
		asset,
	}, nil
}
