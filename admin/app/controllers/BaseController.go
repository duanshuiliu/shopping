package controllers

import (
	"github.com/gin-gonic/gin"
)

type BaseController struct {}

func (this *BaseController) ResponseSuccess(c *gin.Context, data interface{}, code int, info string) {
	this.ResponseFormat(c, 1, code, info, data)
}

func (this *BaseController) ResponseError(c *gin.Context, info string, code int, data interface{}) {
	this.ResponseFormat(c, 0, code, info, data) 
}

func (this *BaseController) ResponseFormat(c *gin.Context, ret, code int, info string, data interface{}) {
	c.JSON(200, gin.H{
		"ret" : ret,
		"code": code,
		"info": info,
		"data": data,
	})
}

