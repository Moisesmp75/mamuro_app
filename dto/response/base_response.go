package dto

import (
	"encoding/json"
	"mamuro_app/common"
)

type BaseResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
	Meta 		*common.Meta `json:"meta"`
}

func NewResponse[T any](data T) BaseResponse[T] {
	return BaseResponse[T]{
		Success: true,
		Message: "",
		Data:    data,
	}
}

func NewResponsePagination[T any](data T, meta common.Meta) BaseResponse[T] {
	return BaseResponse[T] {
		Success: true,
		Message: "",
		Data: data,
		Meta: &meta,
	}
}

func ErrorResponse(errorMessage string) BaseResponse[*string] {
	return BaseResponse[*string]{
		Success: false,
		Message: errorMessage,
		Data:    nil,
	}
}

func JsonResponse[T any](response any) string {
	jsonData, err := json.Marshal(response)
	if err != nil {
		return "Error to convert json error"
	}
	return string(jsonData)
}