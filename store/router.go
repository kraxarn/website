package store

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kraxarn/website/common"
	"github.com/kraxarn/website/config"
	"io/ioutil"
	"net/http"
)

type KeyValue struct {
	key   string
	value string
}

func Route(router *gin.Engine) {
	router.POST("/store/set", func(context *gin.Context) {
		var keyValue KeyValue
		if err := context.BindJSON(&keyValue); err != nil {
			context.JSON(http.StatusOK, common.NewError(err))
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
		jsonData[keyValue.key] = keyValue.value

		if out, err := json.Marshal(&jsonData); err != nil {
			if err = ioutil.WriteFile(path, out, 0644); err != nil {
				context.JSON(http.StatusOK, common.NewError(err))
				return
			}
		}

		context.JSON(http.StatusOK, map[string]interface{}{
			"error": "",
		})
	})
}
