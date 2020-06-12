package admin

import (
	"github.com/gin-gonic/gin"
)

type SiteController struct {
	BaseController
}

func (this *SiteController) Index(c *gin.Context) {
	panic("sdfsdfsdf")
	c.JSON(200, gin.H{
		"message": "I am sitecontroller index",
	})
}
