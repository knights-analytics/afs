package embed

import (
	"context"
	"embed"
	"github.com/knights-analytics/afs/file"
	"github.com/knights-analytics/afs/option"
	"github.com/knights-analytics/afs/storage"
	"io"
	"os"
	"strings"
)

// Open downloads asset for supplied object
func (s *manager) OpenURL(ctx context.Context, URL string, options ...storage.Option) (io.ReadCloser, error) {
	filePath := file.Path(URL)
	filePath = strings.Trim(filePath, "/")
	var efs *embed.FS
	if _, ok := option.Assign(options, &efs); ok {
		return efs.Open(filePath)
	}
	return os.Open(filePath)
}
