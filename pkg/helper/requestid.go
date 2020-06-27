package helper

import (
	"bytes"
    "runtime"
	"strconv"
	"crypto/md5"
	"fmt"
	"time"
)

func GetGID() string {
	b := make([]byte, 64)
    b = b[:runtime.Stack(b, false)]
    b = bytes.TrimPrefix(b, []byte("goroutine "))
    b = b[:bytes.IndexByte(b, ' ')]
    return string(b)
}

func GetUniqueID() string {
	// 这里只使用是时间(纳秒), 并没有加随机数(rand)：一般rand.Seed也是以时间来算，加上了没有意义
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

// 关于生成RequestID的一点思考：
// 之前使用的是GroutineID的生成RequestID, 这样有一个弊端：在http请求中又开一了一个groutine怎么办？
// 独立性：不依赖任何可变的外部环境
// 唯一性：全局唯一
func GetRequestID() string {
	// s := GetGID()
	s := GetUniqueID()
	m := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", m)
}