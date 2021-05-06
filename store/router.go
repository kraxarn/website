package store

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kraxarn/website/common"
	"github.com/kraxarn/website/config"
	"io/ioutil"
	"net/http"
)

func Route(router *gin.Engine) {
	router.GET("/store/set", func(context *gin.Context) {
		key := context.Query("key")
		if key == "" {
			context.JSON(http.StatusOK, common.NewError(fmt.Errorf("no key")))
			return
		}

		value := context.Query("value")
		if value == "" {
			context.JSON(http.StatusOK, common.NewError(fmt.Errorf("no value")))
			return
		}

		path := config.GetPath("store.json")
		var jsonData map[string]string
		data, err := ioutil.ReadFile(path)
		if err == nil {
			if err = json.Unmarshal(data, &jsonData); err != nil {
				context.JSON(http.StatusOK, common.NewError(err))
				return
			}
		}

		if jsonData == nil {
			jsonData = make(map[string]string)
		}
		jsonData[key] = value

		var out []byte
		if out, err = json.Marshal(jsonData); err != nil {
			context.JSON(http.StatusOK, common.NewError(err))
			return
		}
		if err = ioutil.WriteFile(path, out, 0644); err != nil {
			context.JSON(http.StatusOK, common.NewError(err))
			return
		}

		context.JSON(http.StatusOK, map[string]interface{}{
			"error": "",
		})
	})
}
