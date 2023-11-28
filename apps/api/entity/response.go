package entity

type OkResponse[T any] struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   T      `json:"data"`
}

type ErrResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
