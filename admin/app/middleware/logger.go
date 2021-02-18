package middleware

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"bytes"

	"shopping/admin/pkg/logger"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		params,_ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(params))

		logger.NewLogger().Info("Request start", zap.String("url", c.Request.URL.String()), zap.String("method", c.Request.Method), zap.String("params", string(params)))
		//fmt.Println("Admin Logger: ", c.Request.Method, c.Request.URL)
		c.Next()
	}
}
