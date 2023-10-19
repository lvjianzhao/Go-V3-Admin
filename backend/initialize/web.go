package initialize

import (
	"embed"
	"errors"
	"io/fs"
	"path"
	"path/filepath"
	"server/resources"
	"strings"
)

type Resource struct {
	fs   embed.FS
	path string
}

func NewResource() *Resource {
	return &Resource{
		fs: resources.Static,
		//这里的路径要注意，别写错了
		path: "dist",
	}
}

func (r *Resource) Open(name string) (fs.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}
	fullName := filepath.ToSlash(filepath.Join(r.path, path.Clean(name)))
	file, err := r.fs.Open(fullName)
	if err != nil {
		name = "static/" + name
		fullName = filepath.ToSlash(filepath.Join(r.path, path.Clean(name)))
		file, err = r.fs.Open(fullName)
	}
	return file, err
}
