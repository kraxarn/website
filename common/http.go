package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get: %v", err)
	}

	data, err := ioutil.ReadAll(response.Body)
	bodyErr := response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %v", err)
	}
	if bodyErr != nil {
		return nil, fmt.Errorf("response.Body.Close: %v", bodyErr)
	}

	return data, nil
}
