package router

import (
	"github.com/gin-gonic/gin"
	// capi   "shopping/controller/api"
	cadmin "shopping/app/controllers/admin"
	//madmin "shopping/app/middleware/admin"
)

func AddRoute(r *gin.Engine) {

	admin := r.Group("/admin").Use()
	{
		category := &cadmin.CategoryController{}	
		admin.GET("/category", category.List)
		admin.GET("/category/{}", category.Show)
		admin.POST("/category", category.Create)
		admin.PUT("/category/{}", category.Update)
		admin.DELETE("/category/{}", category.Delete)
	}
}
