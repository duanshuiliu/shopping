package middleware

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
	        if err := recover(); err != nil {
	        	fmt.Println("Admin Recovery: ", err)
	        }
		}()
		c.Next()
	}
}
