package response

type response struct {
	Msg      string      `json:"msg,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Page     int         `json:"page,omitempty"`
	PageSize int         `json:"page_size,omitempty"`
	Total    int         `json:"total,omitempty"`
}

func Success(data interface{}) response {
	return response{
		Msg:  "success",
		Data: data,
	}
}

func SuccessPage(data interface{}, page, pageSize int, total int) response {
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
