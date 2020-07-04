package middleware

import (
	"github.com/gin-gonic/gin" 
	"fmt"
	"io/ioutil"
	"bytes"

	logger "shopping/pkg/logger"
	zap    "go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Admin Logger: ", c.Request.Method, c.Request.URL)
		c.Next()
	}
}

func PreRequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		params,_ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(params))
		logger.NewLogger(c).Info("Get Request Params",  zap.String("params", string(params)))

		c.Next()
	}
}
