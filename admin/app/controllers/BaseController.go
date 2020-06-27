package controllers

import (
	"github.com/gin-gonic/gin"
	"encoding/json"

	logger   "shopping/pkg/logger"
	zap      "go.uber.org/zap"
)

type BaseController struct {}

type ResponseData struct {
	Ret  int  `json:"ret"`
	Code int  `json:"code"`
	Info string `json:"info"`
	Data interface{} `json:"data"`
}

func (this *BaseController) ResponseSuccess(c *gin.Context, data interface{}, code int, info string) {
	response := &ResponseData{
		Ret : 1,
		Code: code,
		Info: info,
		Data: data,
	}

	jsonStr,_ := json.Marshal(response)

	logger.NewLogger(c).Info("Response", zap.String("data", string(jsonStr)))
	this.ResponseFormat(c, response)
}

func (this *BaseController) ResponseError(c *gin.Context, info string, code int, data interface{}) {
	response := &ResponseData{
		Ret : 0,
		Code: code,
		Info: info,
		Data: data,
	}

	jsonStr,_ := json.Marshal(response)

	logger.NewLogger(c).Error("Response", zap.String("data", string(jsonStr)))
	this.ResponseFormat(c, response) 
}

func (this *BaseController) ResponseFormat(c *gin.Context, res *ResponseData) {
	c.JSON(200, res)
}

