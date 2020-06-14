package admin

import (
	"github.com/gin-gonic/gin"
	"fmt"
	
	aService "shopping/app/services/admin"
)

type CategoryController struct {
	BaseController
}

func (this *CategoryController) List(c *gin.Context) {
	this.ResponseSuccess(c, 1, 200, "OK")
}

func (this *CategoryController) Show(c *gin.Context) {
	// 
}

func (this *CategoryController) Create(c *gin.Context) {
	category := &aService.Category{}

	var err error

	// 获取参数
	data, err := category.ValidateOfCreate(c)

	if err != nil {
		this.ResponseError(c, err.Error(), 500, nil)
		return
	}

	result, err := category.Create(data)

	if err != nil {
		this.ResponseError(c, err.Error(), 500, nil)
		return	
	}

	fmt.Println("返回数据", result)
	this.ResponseSuccess(c, result.ID, 200, "success")
}

func (this *CategoryController) Update(c *gin.Context) {
	// 
}

func (this *CategoryController) Delete(c *gin.Context) {
	// 
}
