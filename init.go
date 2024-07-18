package afs

import (
	"github.com/knights-analytics/afs/file"
	"github.com/knights-analytics/afs/http"
	"github.com/knights-analytics/afs/mem"
	"github.com/knights-analytics/afs/scp"
	"github.com/knights-analytics/afs/ssh"
	"github.com/knights-analytics/afs/tar"
	"github.com/knights-analytics/afs/zip"
)

func init() {
	registry := GetRegistry()
	registry.Register(file.Scheme, file.Provider)
	registry.Register(mem.Scheme, mem.Provider)
	registry.Register(http.Scheme, http.Provider)
	registry.Register(http.SecureScheme, http.Provider)
	registry.Register(scp.Scheme, scp.Provider)
	registry.Register(ssh.Scheme, scp.Provider)
	registry.Register(zip.Scheme, zip.Provider)
	registry.Register(tar.Scheme, tar.Provider)
}
