package admin

import (
	"github.com/gin-gonic/gin" 
	"fmt"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Admin Logger: ", c.Request.Method, c.Request.URL)
		c.Next()
	}
}
