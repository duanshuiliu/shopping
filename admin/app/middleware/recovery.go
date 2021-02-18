package middleware

import (
	"github.com/gin-gonic/gin"
	"fmt"

	"shopping/admin/pkg/exception"
	"shopping/admin/pkg/helper"
	"shopping/admin/pkg/logger"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
	        if err := recover(); err != nil {
				// fmt.Printf("Admin Recovery: %#v", err)

				switch rerr := err.(type) {
					case *exception.MsgException:
						helper.ResponseError(c, rerr.Error(), rerr.GetCode(), nil)
					default:
						logger.NewLogger().Error(fmt.Sprintf("%s", rerr))
						helper.ResponseError(c, "系统繁忙，请稍后重试", 500, nil)
				}
				return
	        }
		}()
		
		c.Next()
	}
}
