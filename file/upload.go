package file

import (
	"context"
	"github.com/knights-analytics/afs/object"
	"github.com/knights-analytics/afs/option"
	"github.com/knights-analytics/afs/storage"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// EnsureParentPathExists create parent path if needed
func EnsureParentPathExists(filename string, fileMode os.FileMode) error {
	filename = normalize(filename)
	stat, err := os.Stat(filename)
	if err == nil {
		if stat.Mode() != fileMode {
			_ = os.Chmod(filename, fileMode)
		}
		return err
	}
	parent, _ := filepath.Split(filename)
	if runtime.GOOS == "windows" && strings.Count(parent, `\`) == 1 {
		return nil
	}
	return os.MkdirAll(parent, fileMode)
}

// Upload writes reader content to supplied URL path.
func Upload(ctx context.Context, URL string, mode os.FileMode, reader io.Reader, options ...storage.Option) error {
	filePath := Path(URL)
	err := EnsureParentPathExists(filePath, DefaultDirOsMode)
	if err != nil {
		return errors.Wrap(err, "unable to create parent for "+filePath)
	}
	link := &object.Link{}
	option.Assign(options, &link)
	if link.Linkname != "" {
		return os.Symlink(filePath, link.Linkname)
	}

	stat, _ := os.Stat(filePath)
	if stat != nil {
		_ = os.Remove(filePath)
	}
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, mode)
	if err != nil {
		return errors.Wrapf(err, "unable to open file: %v ", filePath)
	}
	_, err = io.Copy(file, reader)
	if closeErr := file.Close(); err == nil {
		err = closeErr
	}
	return err
}
