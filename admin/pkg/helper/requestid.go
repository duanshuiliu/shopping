package helper

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"runtime"
)

func GetGID() string {
	b := make([]byte, 64)
    b = b[:runtime.Stack(b, false)]
    b = bytes.TrimPrefix(b, []byte("goroutine "))
    b = b[:bytes.IndexByte(b, ' ')]
    return string(b)
}

// 关于生成RequestID的一点思考：
// 之前使用的是GroutineID的生成RequestID, 这样有一个弊端：在http请求中又开一了一个groutine怎么办？
// 独立性：不依赖任何可变的外部环境
// 唯一性：全局唯一
func GetRequestID() string {
	s := GetGID()
	m := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", m)
}