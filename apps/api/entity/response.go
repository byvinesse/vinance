package entity

type OkResponse[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}

type ErrResponse struct {
	Status    string `json:"status"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}
