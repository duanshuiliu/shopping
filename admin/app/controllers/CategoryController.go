package controllers

import (
	"github.com/gin-gonic/gin"
	// "fmt"
	
	pError    "shopping/pkg/error"
	pResponse "shopping/pkg/response"
	aService  "shopping/admin/app/services"
)

type CategoryController struct {
	BaseController
}

func (this *CategoryController) List(c *gin.Context) {
	pResponse.ResponseSuccess(c, 1, 200, "OK")
	return
}

func (this *CategoryController) Show(c *gin.Context) {
	category := &aService.Category{}

	// 获取参数
	data, err := category.ValidateOfShow(c)

	if err != nil {
		panic(&pError.MessageError{Message: err.Error()})
	}

	result, err := category.Show(data)
	if err != nil { panic(err) }

	if result == nil {
		pResponse.ResponseError(c, "not found data", 500, nil)
		return
	}

	pResponse.ResponseSuccess(c, result, 200, "success")
	return
}

func (this *CategoryController) Create(c *gin.Context) {
	category := &aService.Category{}

	// 获取参数
	data, err := category.ValidateOfCreate(c)

	if err != nil {
		panic(&pError.MessageError{Message: err.Error()})
	}

	result, err := category.Create(data)

	if err != nil {
		panic(&pError.MessageError{Message: err.Error()})	
	}

	pResponse.ResponseSuccess(c, result, 200, "success")
	return
}

func (this *CategoryController) Update(c *gin.Context) {
	category := &aService.Category{}

	// 获取参数
	data, err := category.ValidateOfUpdate(c)

	if err != nil {
		panic(&pError.MessageError{Message: err.Error()})
	}

	result, err := category.Update(data)

	if err != nil {
		panic(&pError.MessageError{Message: err.Error()})
	}

	pResponse.ResponseSuccess(c, result, 200, "success")
	return
}

func (this *CategoryController) Delete(c *gin.Context) {
	category := &aService.Category{}

	// 获取参数
	data, err := category.ValidateOfDelete(c)

	if err != nil {
		panic(&pError.MessageError{Message: err.Error()})
	}

	result, err := category.Delete(data)
	if err != nil { panic(err) }

	pResponse.ResponseSuccess(c, result, 200, "success")
	return
}