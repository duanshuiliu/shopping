package middleware

import (
	"github.com/gin-gonic/gin"
	"fmt"
	// "errors"

	pError    "shopping/pkg/error"
	pResponse "shopping/pkg/response"
	logger    "shopping/pkg/logger"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
	        if err := recover(); err != nil {
				// fmt.Printf("Admin Recovery: %#v", err)

				switch serr := err.(type) {
					case *pError.MessageError:
						pResponse.ResponseError(c, serr.Error(), serr.GetCode(), nil)
					default:
						logger.NewLogger(c).Error(fmt.Sprintf("%s", serr))
						pResponse.ResponseError(c, "系统繁忙，请稍后重试", 500, nil)
				}
				return
	        }
		}()
		
		c.Next()
	}
}
