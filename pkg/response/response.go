package response

import (
	"github.com/gin-gonic/gin"
	// "encoding/json"

	// logger "shopping/pkg/logger"
	// zap    "go.uber.org/zap"
)

type ResponseData struct {
	// 1=正确输出 0=错误输出
	Ret int `json:"ret"`

	// 错误码
	Code int `json:"code"`

	// 输出信息
	Info string `json:"info"`

	// 输出数据
	Data interface{} `json:"data"`
}

func ResponseSuccess(c *gin.Context, data interface{}, code int, info string) {
	response := &ResponseData{
		Ret : 1,
		Code: code,
		Info: info,
		Data: data,
	}

	// jsonStr,_ := json.Marshal(response)
	// logger.NewLogger(c).Info("Response", zap.String("data", string(jsonStr)))
	ResponseFormat(c, response)
}

func ResponseError(c *gin.Context, info string, code int, data interface{}) {
	response := &ResponseData{
		Ret : 0,
		Code: code,
		Info: info,
		Data: data,
	}

	// jsonStr,_ := json.Marshal(response)
	// logger.NewLogger(c).Error("Response", zap.String("data", string(jsonStr)))
	ResponseFormat(c, response)
}

func ResponseFormat(c *gin.Context, res *ResponseData) {
	c.JSON(200, res)
}