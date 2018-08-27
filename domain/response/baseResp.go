package response

import (
	"beego_framework/domain"
	)

type BaseResp struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type BasePaginationResp struct {
	BaseResp
	domain.PaginatorResp
}

func (baseResp *BaseResp) AppendErrorMessage(msg string) () {
	if msg != "" {
		baseResp.ErrorMessage += "-[" + msg + "]"
	}
}
