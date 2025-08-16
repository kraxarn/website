package teamspeak

import "errors"

type ApiStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ApiResponse[T interface{}] struct {
	Body   T         `json:"body"`
	Status ApiStatus `json:"status"`
}

func StatusError[T interface{}](resp ApiResponse[T]) error {
	if resp.Status.Code == 0 {
		return nil
	}

	return errors.New(resp.Status.Message)
}
