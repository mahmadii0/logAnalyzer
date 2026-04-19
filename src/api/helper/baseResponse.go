package helper

import (
	"logAnalyzer/api/validations"
)

type BaseHttpResponse struct {
	Result          any                            `json:result`
	Success         bool                           `json:success`
	ResultCode      int                            `json:resultCode`
	ValidationError *[]validations.ValidationError `json:validationError`
	Error           any                            `json:error`
}

func GenerateBaseR(result any, success bool, resultCode int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
	}
}

func GenerateBaseRwithError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err.Error(),
	}
}

func GenerateBaseRwithValidationError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:          result,
		Success:         success,
		ResultCode:      resultCode,
		ValidationError: validations.GetValidationErrors(err),
	}
}
