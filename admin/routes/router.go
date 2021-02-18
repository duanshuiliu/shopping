package routes

import (
	"github.com/gin-gonic/gin"
	
	controller "shopping/admin/app/controllers"
	middleware "shopping/admin/app/middleware"
)

func AddRoute(r *gin.Engine) {

	admin := r.Group("/").Use(middleware.Recovery(), middleware.Logger())
	{
		// 类别
		category := &controller.CategoryController{}	
		admin.GET("/category", category.List)
		admin.GET("/category/:id", category.Show)
		admin.POST("/category", category.Create)
		admin.PUT("/category/:id", category.Update)
		admin.DELETE("/category/:id", category.Delete)
	}
}
