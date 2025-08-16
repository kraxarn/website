package teamspeak

type ApiStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ApiResponse[T interface{}] struct {
	Body   T         `json:"body"`
	Status ApiStatus `json:"status"`
}
