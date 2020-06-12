package router

import (
	"github.com/gin-gonic/gin"
	// capi   "shopping/controller/api"
	cadmin "shopping/app/controller/admin"
	madmin "shopping/app/middleware/admin"
)

func AddRouter(r *gin.Engine) {

	admin := r.Group("/admin").Use(madmin.Logger(), madmin.Recovery())
	{
		site := &cadmin.SiteController{}	
		admin.GET("/site", site.Index)
	}
}
