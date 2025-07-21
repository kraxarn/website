package helper

import (
	"encoding/json"
	"io"
	"net/http"
)

func Get(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var data []byte
	data, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = response.Body.Close()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetJson[T *any](url string) (T, error) {
	body, err := Get(url)
	if err != nil {
		return nil, err
	}

	var value T
	err = json.Unmarshal(body, value)
	if err != nil {
		return nil, err
	}

	return value, nil
}
