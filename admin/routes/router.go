package router

import (
	"github.com/gin-gonic/gin"
	
	cadmin "shopping/admin/app/controllers/admin"
	//madmin "shopping/app/middleware/admin"
)

func AddRoute(r *gin.Engine) {

	admin := r.Group("/").Use()
	{
		// 类别
		category := &cadmin.CategoryController{}	
		admin.GET("/category", category.List)
		admin.GET("/category/:id", category.Show)
		admin.POST("/category", category.Create)
		admin.PUT("/category/:id", category.Update)
		admin.DELETE("/category/:id", category.Delete)
	}
}
