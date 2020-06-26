package controllers

import (
	"github.com/gin-gonic/gin"
	"fmt"
	
	aService "shopping/admin/app/services"
)

type CategoryController struct {
	BaseController
}

func (this *CategoryController) List(c *gin.Context) {
	this.ResponseSuccess(c, 1, 200, "OK")
}

func (this *CategoryController) Show(c *gin.Context) {
	category := &aService.Category{}

	// 获取参数
	data, err := category.ValidateOfShow(c)

	if err != nil {
		this.ResponseError(c, err.Error(), 500, nil)
	}

	result, err := category.Show(data)

	if err != nil {
		this.ResponseError(c, err.Error(), 500, nil)
	}

	this.ResponseSuccess(c, result, 200, "success")
}

func (this *CategoryController) Create(c *gin.Context) {
	category := &aService.Category{}

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
	return
}

func (this *CategoryController) Update(c *gin.Context) {
	category := &aService.Category{}

	// 获取参数
	data, err := category.ValidateOfUpdate(c)

	if err != nil {
		this.ResponseError(c, err.Error(), 500, nil)
		return
	}

	result, err := category.Update(data)

	if err != nil {
		this.ResponseError(c, err.Error(), 500, nil)
		return
	}

	this.ResponseSuccess(c, result, 200, "success")
	return
}

func (this *CategoryController) Delete(c *gin.Context) {
	category := &aService.Category{}

	// 获取参数
	data, err := category.ValidateOfDelete(c)

	if err != nil {
		this.ResponseError(c, err.Error(), 500, nil)
		return
	} 

	_, err = category.Delete(data)

	if err != nil {
		this.ResponseError(c, err.Error(), 500, nil)
		return
	}

	this.ResponseSuccess(c, nil, 200, "success")
	return
}
