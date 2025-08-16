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
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", a.baseUrl, path), nil)
	if err != nil {
		return err
	}

	var apiKey string
	apiKey, err = config.TeamSpeakApiKey()
	if err != nil {
		return err
	}
	req.Header.Add("x-api-key", apiKey)

	var resp *http.Response
	resp, err = a.httpClient.Do(req)
	if err != nil {
		return err
	}

	var data []byte
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return errors.Join(err, resp.Body.Close())
	}

	err = resp.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, value)
	if err != nil {
		return err
	}

	return nil
}
