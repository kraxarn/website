package main

import (
	"fmt"
	"github.com/kraxarn/website/common"
	"github.com/kraxarn/website/config"
	"github.com/kraxarn/website/sponsor"
	"github.com/kraxarn/website/user"
	"github.com/kraxarn/website/yt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FileInfo struct {
	Size               int64
	DateModified, Name string
	IsDirectory        bool
}

func FormatFileSize(size int64) string {
	// gb
	if size > 1_000_000_000 {
		return fmt.Sprintf("%dG", size/1_000_000_000)
	}
	// mb
	if size > 1_000_000 {
		return fmt.Sprintf("%dM", size/1_000_000)
	}
	// kb
	if size > 1_000 {
		return fmt.Sprintf("%dk", size/1_000)
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
	token := config.NewToken()

	// Setup some html functions
	router.SetFuncMap(template.FuncMap{
		"dateNow": func() string {
			return time.Now().String()
		},
		"formatFileSize": FormatFileSize,
		"currentVersion": func() string {
			return config.CurrentVersion
		},
	})

	// Add all files in html folder as templates
	router.LoadHTMLGlob("html/*.html")

	// Add all folders and files in static folder
	staticFiles, _ := ioutil.ReadDir("static")
	for _, file := range staticFiles {
		filePath := fmt.Sprintf("static/%v", file.Name())
		if file.IsDir() {
			router.Static(file.Name(), filePath)
		} else {
			router.StaticFile(file.Name(), filePath)
		}
	}

	// Show index when loading root
	router.GET("", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	// /servers
	router.GET("servers/*app", func(context *gin.Context) {
		context.HTML(http.StatusOK, "servers.html", gin.H{
			"infos": GetServerInfo(),
		})
	})

	router.GET("ytdl", func(context *gin.Context) {
		context.HTML(http.StatusOK, "ytdl.html", nil)
	})

	router.GET("changes", func(context *gin.Context) {
		releases, err := common.GetVersions()
		if !common.NewError(err).SendIfError(context) {
			return
		}
		context.JSON(http.StatusOK, releases)
	})

	user.Route(router, &token)
	yt.Route(router)
	sponsor.Route(router)

	// Add all folders in files
	fileFiles, err := ioutil.ReadDir("files")
	if err == nil {
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
	if err := router.Run("127.0.0.1:5000"); err != nil {
		fmt.Println(err)
	}
}
