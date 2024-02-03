package response

type BaseResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func APIResponse(code int, status string, payload interface{}) BaseResponse {
	return BaseResponse{
		Code:   code,
		Status: status,
		Data:   payload,
	}
}
