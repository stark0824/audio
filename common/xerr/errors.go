package xerr

import (
	"fmt"
)

/**
常用通用固定错误
*/

type CodeError struct {
	ErrCode uint32
	ErrMsg  string
}

// 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.ErrCode
}

// 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.ErrMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.ErrCode, e.ErrMsg)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{ErrCode: errCode, ErrMsg: errMsg}
}
func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{ErrCode: errCode, ErrMsg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{ErrCode: UNKNOWN, ErrMsg: errMsg}
}
