package teamspeak

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kraxarn/website/config"
	"io"
	"net/http"
)

type Api struct {
	httpClient *http.Client
	baseUrl    string
}

func NewApi() (Api, error) {
	url, err := config.TeamSpeakUrl()
	if err != nil {
		return Api{}, err
	}

	return Api{
		httpClient: &http.Client{},
		baseUrl:    url,
	}, nil
}

func (a Api) get(path string, value any) error {
	response, err := a.httpClient.Get(fmt.Sprintf("%s%s", a.baseUrl, path))
	if err != nil {
		return err
	}

	var data []byte
	data, err = io.ReadAll(response.Body)
	if err != nil {
		return errors.Join(err, response.Body.Close())
	}

	err = response.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, value)
	if err != nil {
		return err
	}

	return nil
}
