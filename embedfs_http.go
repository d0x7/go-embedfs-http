package embedfshttp

import (
	"embed"
	iofs "io/fs"
	"net/http"
	"path/filepath"
)

type NonIndexingFileSystem struct {
	fs http.FileSystem
}

func (fs NonIndexingFileSystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	defer func(f http.File) {
		_ = f.Close()
		// We never write to the file, so an error should never happen.
	}(f)

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(name, "index.html")
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err // return the os.ErrNotExist error we just got
		}
	}

	return f, nil
}

func New(fs embed.FS, path string) http.Handler {
	dir, err := iofs.Sub(fs, path)
	if err != nil {
		panic(err)
	}

	newFs := NonIndexingFileSystem{http.FS(dir)}

	return http.FileServer(newFs)
}
