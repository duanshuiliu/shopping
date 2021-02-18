package controllers

import (
	"github.com/gin-gonic/gin"

	"shopping/admin/app/services"
	"shopping/admin/pkg/helper"
)

type CategoryController struct {
	BaseController
}

func (this *CategoryController) List(c *gin.Context) {
	category := &services.Category{}

	// 获取参数
	data, err := category.ValidateOfList(c)
	if err != nil { panic(err) }

	result, err := category.List(data)
	if err != nil { panic(err) }

	helper.ResponseSuccess(c, result, 200, "success")
	return
}

func (this *CategoryController) Show(c *gin.Context) {
	category := &services.Category{}

	// 获取参数
	data, err := category.ValidateOfShow(c)
	if err != nil { panic(err) }

	result, err := category.Show(data)
	if err != nil { panic(err) }

	helper.ResponseSuccess(c, result, 200, "success")
	return
}

func (this *CategoryController) Create(c *gin.Context) {
	category := &services.Category{}

	// 获取参数
	data, err := category.ValidateOfCreate(c)
	if err != nil { panic(err) }

	m, err := category.Create(data)
	if err != nil { panic(err) }

	helper.ResponseSuccess(c, m.ID, 200, "success")
	return
}

func (this *CategoryController) Update(c *gin.Context) {
	category := &services.Category{}

	// 获取参数
	data, err := category.ValidateOfUpdate(c)
	if err != nil { panic(err) }

	_, err = category.Update(data)
	if err != nil { panic(err) }

	helper.ResponseSuccess(c, nil, 200, "success")
	return
}

func (this *CategoryController) Delete(c *gin.Context) {
	category := &services.Category{}

	// 获取参数
	data, err := category.ValidateOfDelete(c)
	if err != nil { panic(err) }

	result, err := category.Delete(data)
	if err != nil { panic(err) }

	helper.ResponseSuccess(c, result, 200, "success")
	return
}