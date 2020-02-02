package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
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

func main() {
	// Create router and add some middleware
	// (using .Default directly generates a warning)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Setup some html functions
	router.SetFuncMap(template.FuncMap{
		"dateNow": func() string {
			return time.Now().String()
		},
		"formatFileSize": FormatFileSize,
	})

	// Add all files in html folder as templates
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

	// Show index when loading root
	router.GET("", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	// Add all folders in files
	fileFiles, err := ioutil.ReadDir("files")
	if err != nil {
		fmt.Println("warning: failed to read 'files' directory:", err)
	} else {
		for _, file := range fileFiles {
			router.GET(fmt.Sprintf("%v/*file", file.Name()), func(context *gin.Context) {
				HandleList(context)
			})
		}
	}

	// When page is not found, redirect page to home
	router.NoRoute(func(context *gin.Context) {
		context.Redirect(http.StatusFound, "/")
	})

	// Start listening on port 8080
	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
