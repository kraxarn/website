package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func IsFile(path string) bool {
	info, err := os.Stat(filepath.Join("files", path))
	return err != nil || !info.IsDir()
}

func List(path string) []FileInfo {
	fileInfo := make([]FileInfo, 0)
	files, _ := ioutil.ReadDir(filepath.Join("files", path))
	for _, file := range files {
		fileInfo = append(fileInfo, FileInfo{
			Size:         file.Size(),
			DateModified: file.ModTime().Format("2006-01-02 15:04"),
			Name:         file.Name(),
			IsDirectory:  file.IsDir(),
		})
	}
	return fileInfo
}

func HandleList(context *gin.Context) {
	path := filepath.Join(
		context.FullPath()[:len(context.FullPath())-5],
		context.Param("file"))
	if IsFile(path) {
		context.File(filepath.Join("files", path))
	} else {
		context.HTML(http.StatusOK, "ls.html", gin.H{
			"path": path,
			"dir": List(path),
		})
	}
}