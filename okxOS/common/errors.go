package common

import (
	"errors"
	"reflect"
)

// APIError OKX API 错误
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// CheckAPIError 检查 API 响应中的错误
func CheckAPIError(response interface{}) error {
	v := reflect.ValueOf(response)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}

	codeField := v.FieldByName("Code")
	msgField := v.FieldByName("Msg")
	if codeField.IsValid() && msgField.IsValid() && codeField.Int() != 0 {
		return errors.New(msgField.String())
	}
	return nil
}
