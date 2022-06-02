package response

type response struct {
	Msg      string      `json:"msg,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Page     int32       `json:"page,omitempty"`
	PageSize int32       `json:"page_size,omitempty"`
	Total    int64       `json:"total,omitempty"`
}

func Success(data interface{}) response {
	return response{
		Msg:  "success",
		Data: data,
	}
}

func SuccessPage(data interface{}, page, pageSize int32, total int64) response {
	return response{
		Msg:      "success",
		Data:     data,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}
}

func Error(err error) response {
	return response{
		Msg: err.Error(),
	}
}
