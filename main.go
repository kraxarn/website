package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type FileInfo struct {
	Size int64
	DateModified, Name string
	IsDirectory bool
}

func FormatFileSize(size int64) string {
	// gb
	if size > 1000000000 {
		return fmt.Sprintf("%dG", size / 1000000000)
	}
	// mb
	if size > 1000000 {
		return fmt.Sprintf("%dM", size / 1000000)
	}
	// kb
	if size > 1000 {
		return fmt.Sprintf("%dk", size / 1000)
	}
	// b
	return fmt.Sprintf("%d", size)
}

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
	file := context.Param("file")
	if strings.HasSuffix(file, "/") {
		file = file[:len(file)-1]
	}
	path := context.FullPath()
	path = path[:len(path)-6] + file
	if IsFile(path) {
		context.File(filepath.Join("files", path))
	} else {
		context.HTML(http.StatusOK, "ls.html", gin.H{
			"path": path,
			"dir": List(path),
		})
	}
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.SetFuncMap(template.FuncMap{
		"dateNow": func() string {
			return time.Now().String()
		},
		"formatFileSize": FormatFileSize,
	})

	router.LoadHTMLGlob("html/*.html")

	// Add all folders and files in static folder
	staticFiles, _ := ioutil.ReadDir("static")
	for _, file := range staticFiles {
		path := fmt.Sprintf("static/%v", file.Name())
		if file.IsDir() {
			router.Static(file.Name(), path)
		} else {
			router.StaticFile(file.Name(), path)
		}
	}

	router.GET("", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("minecraft/*file", func(context *gin.Context) {
		HandleList(context)
	})

	// When page is not found, redirect page to home
	router.NoRoute(func(context *gin.Context) {
		context.Redirect(http.StatusFound, "/")
	})

	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
