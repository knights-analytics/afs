package cache

import (
	"context"
	"github.com/knights-analytics/afs"
	"github.com/knights-analytics/afs/file"
	"github.com/knights-analytics/afs/option"
	"github.com/knights-analytics/afs/storage"
	"github.com/knights-analytics/afs/url"
	"strings"
)

// Package creates cache file for source URL with rewrite
func Package(ctx context.Context, sourceURL string, rewriteBaseURL string, options ...storage.Option) error {
	var cacheOption = &option.Cache{}
	option.Assign(options, &cacheOption)
	if cacheOption.Name == "" {
		cacheOption.Name = CacheFile
	}
	cacheOption.Init()
	cacheURL := url.Join(sourceURL, cacheOption.Name)
	fs := afs.New()
	cache, err := build(ctx, sourceURL, cacheOption.Name, fs, options...)
	if err != nil || len(cache.Items) == 0 {
		return err
	}
	sourceURL = url.Normalize(sourceURL, file.Scheme)
	for _, entry := range cache.Items {
		location := strings.Replace(entry.URL, sourceURL, "", 1)
		entry.URL = url.Join(rewriteBaseURL, location)
	}
	return uploadCacheFile(ctx, cache, cacheURL, fs)
}
