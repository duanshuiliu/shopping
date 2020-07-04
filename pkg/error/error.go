package error

import (
	// 
)

type MessageError struct {
	// 错误码
	Code int
	
	// 错误信息
	Message string
}

func (this *MessageError) Error() string {
	return this.Message
}

func (this *MessageError) GetCode() int {
	code := this.Code

	if code == 0 {
		code = 500
	}

	return code
}