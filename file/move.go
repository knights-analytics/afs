package file

import (
	"context"
	"github.com/knights-analytics/afs/storage"
	"github.com/pkg/errors"
	"os"
)

// Move moves source to URL
func Move(ctx context.Context, sourceURL, destURL string, options ...storage.Option) error {
	sourcePath := Path(sourceURL)
	destPath := Path(destURL)
	_ = os.RemoveAll(destPath)
	err := EnsureParentPathExists(destPath, DefaultDirOsMode)
	if err != nil {
		return errors.Wrap(err, "unable to create parent for "+destPath)
	}
	return os.Rename(sourcePath, destPath)
}
