package http

import (
	"context"
	"fmt"
	"github.com/knights-analytics/afs/option"
	"github.com/knights-analytics/afs/storage"
	"io"
	"net/http"
	"os"
)

// Create send post request
func (s *manager) Create(ctx context.Context, URL string, mode os.FileMode, isDir bool, options ...storage.Option) error {
	var reader io.Reader
	option.Assign(options, &reader)
	request, err := http.NewRequest(http.MethodPost, URL, reader)
	if err != nil {
		return err
	}
	response, err := s.run(ctx, URL, request, options...)
	if err != nil {
		return err
	}
	s.closeResponse(response)
	if IsStatusOK(response) {
		return nil
	}
	return fmt.Errorf("invalid status code: %v", response.StatusCode)

}
