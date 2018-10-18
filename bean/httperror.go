package bean

import (
	"beego_framework/domain/logicerror"
	"beego_framework/domain/response"
	"fmt"
)

const (
	UNKNOWN_ERROR = "ew_000"
	INPUT_ERROR   = "ew_001"
)

func InitErrorMap() {
	ErrorMap = make(map[string]response.BaseResp)
	ErrorMap[UNKNOWN_ERROR] = response.BaseResp{
		ErrorCode: UNKNOWN_ERROR, ErrorMessage: "Unknown error",
	}
	ErrorMap[INPUT_ERROR] = response.BaseResp{
		ErrorCode: INPUT_ERROR, ErrorMessage: "Error input",
	}
}

func HandleError(err interface{}) (*response.BaseResp) {
	var baseResp response.BaseResp
	switch err.(type) {
	case logicerror.InputInvalidError:
		baseResp = ErrorMap[INPUT_ERROR]
	default:
		baseResp = ErrorMap[UNKNOWN_ERROR]
	}
	if baseResp.ErrorCode != UNKNOWN_ERROR {
		(&baseResp).AppendErrorMessage((err.(error)).Error())
	}
	fmt.Println(err)
	return &baseResp
}
