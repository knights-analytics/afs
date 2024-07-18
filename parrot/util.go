package parrot

import (
	"github.com/knights-analytics/afs/file"
	"github.com/knights-analytics/afs/url"
)

// Pkg returns package name for location
func Pkg(location string) string {
	parent, _ := url.Split(location, file.Scheme)
	_, pkg := url.Split(parent, file.Scheme)
	return pkg
}
