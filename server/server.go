package server

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed www
var web embed.FS

func Init() *http.ServeMux {
	fsys, err := fs.Sub(web, "www")
	if err != nil {
		panic("remove public path")
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(fsys)))
	mux.HandleFunc("/msgs", msgHandle)
	mux.HandleFunc("/get", getHandle)
	// mux.HandleFunc("/search", searchHandle)

	return mux
}
