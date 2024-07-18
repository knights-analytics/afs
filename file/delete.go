package file

import (
	"context"
	"github.com/knights-analytics/afs/storage"
	"os"
)

// Delete removes file or directory
func Delete(ctx context.Context, URL string, options ...storage.Option) error {
	filePath := Path(URL)
	return os.RemoveAll(filePath)
}
