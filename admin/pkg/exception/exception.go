package exception

import (
//
)

type MsgException struct {
	// 错误码
	Code int
	// 错误信息
	Message string
}

func (this *MsgException) Error() string {
	return this.Message
}

func (this *MsgException) GetCode() int {
	code := this.Code

	if code == 0 {
		code = 500
	}

	return code
}

type ValidateException struct {
	MsgException
}