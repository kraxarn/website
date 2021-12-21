package main

import (
	"fmt"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/kraxarn/website/chat"
	"github.com/kraxarn/website/common"
	"github.com/kraxarn/website/config"
	"github.com/kraxarn/website/format"
	"github.com/kraxarn/website/sponsor"
	"github.com/kraxarn/website/user"
	"github.com/kraxarn/website/watch"
	"github.com/kraxarn/website/yt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type FileInfo struct {
	Size               int64
	DateModified, Name string
	IsDirectory        bool
}

func main() {
	// Create router and add some middleware
	// (using .Default directly generates a warning)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	token := config.NewToken()

	// For the number generator to generate different numbers
	rand.Seed(time.Now().Unix())

	// Setup some html functions
	router.SetFuncMap(template.FuncMap{
		"dateNow": func() string {
			return time.Now().String()
		},
		"formatFileSize": format.FileSize,
		"currentVersion": func() string {
			return config.CurrentVersion
		},
		"currentWatchVersion": func() string {
			return watch.CurrentVersion
		},
	})

	// Add all files in html folder as templates
	router.LoadHTMLFiles(append([]string{
		"html/index.gohtml", "html/ls.gohtml",
		"html/servers.gohtml", "html/ytdl.gohtml",
	}, watch.HtmlFiles()...)...)

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
		context.HTML(http.StatusOK, "index.gohtml", nil)
	})

	// /servers
	router.GET("servers/*app", func(context *gin.Context) {
		context.HTML(http.StatusOK, "servers.gohtml", gin.H{
			"infos": GetServerInfo(),
		})
	})

	router.GET("ytdl", func(context *gin.Context) {
		context.HTML(http.StatusOK, "ytdl.gohtml", nil)
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
	chat.Route(router, &token)
	watch.Route(router, &token)

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
		context.HTML(http.StatusNotFound, "index.gohtml", nil)
	})

	// Run secure in release, or insecure when debugging
	if config.IsSecure() {
		err = autotls.Run(router, "kraxarn.com")
	} else {
		err = router.Run("127.0.0.1:5000")
	}

	if err != nil {
		log.Fatal(err)
	}
}
