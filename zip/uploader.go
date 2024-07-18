package zip

import (
	"archive/zip"
	"bytes"
	"context"
	"github.com/knights-analytics/afs/file"
	"github.com/knights-analytics/afs/option"
	"github.com/knights-analytics/afs/storage"
	"io"
	"os"
	"path"
	"strings"
)

type uploader struct {
	uploader storage.Uploader
	buffer   *bytes.Buffer
}

func (u *uploader) Uploader(ctx context.Context, URL string, options ...storage.Option) (storage.Upload, io.Closer, error) {
	var uploader storage.Uploader
	option.Assign(options, &u.buffer, &uploader)
	if uploader == nil {
		uploader = u.uploader
	}
	if u.buffer == nil {
		u.buffer = new(bytes.Buffer)
	}
	writer := newWriter(ctx, u.buffer, URL, uploader)
	return func(ctx context.Context, parent string, info os.FileInfo, reader io.Reader) error {
		filename := path.Join(parent, info.Name())
		mode := info.Mode().Perm()
		if info.IsDir() {
			mode |= os.ModeDir
		}
		info = file.NewInfo(filename, info.Size(), mode, info.ModTime(), info.IsDir())
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Method = zip.Deflate
		header.Name = filename
		if info.IsDir() && !strings.HasSuffix(filename, "/") {
			header.Name += "/"
		}
		writer, err := writer.CreateHeader(header)
		if reader != nil {
			_, err = io.Copy(writer, reader)
			if err != nil {
				return err
			}
		}
		return err
	}, writer, nil
}

// newBatchUploader returns a batch uploader
func newBatchUploader(dest storage.Uploader) *uploader {
	return &uploader{uploader: dest}
}

// NewBatchUploader returns a batch uploader
func NewBatchUploader(dest storage.Uploader) storage.BatchUploader {
	return newBatchUploader(dest)
}
