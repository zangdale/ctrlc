package server

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed www
var web embed.FS

func Init(webDir string) *http.ServeMux {
	fsys, err := fs.Sub(web, "www")
	if err != nil {
		panic("remove public path")
	}
	fs := http.FS(fsys)
	if webDir != "" {
		fs = http.Dir(webDir)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(fs))
	mux.HandleFunc("/msgs", msgHandle)
	mux.HandleFunc("/get", getHandle)
	// mux.HandleFunc("/search", searchHandle)

	return mux
}
