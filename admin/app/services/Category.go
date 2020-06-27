package services

import (
	"github.com/gin-gonic/gin"
	//"fmt"
	"strconv"

	mShopping "shopping/pkg/models/shopping"
)

type Category struct {
	BaseService
}

func (this *Category) ValidateOfShow(c *gin.Context) (uint, error) {
	id := c.Param("id")

	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return 0, err
	}

	return uint(i), nil
}

func (this *Category) Show(id uint) (category *mShopping.Category, err error) {
	category = &mShopping.Category{}
	// category.ID = id

	model, err := category.SearchOne(category)

	if err != nil { return }

	category, ok := model.(*mShopping.Category)

	if !ok {
		err = ErrStruct
		return
	}

	return
}

type CategoryCreate struct {
	Type int8    `json:"type"`
	Pid  uint    `json:"pid"`
	Name string  `json:"name"`
	Desc string  `json:"desc"`
}

func (this *Category) ValidateOfCreate(c *gin.Context) (*CategoryCreate, error) {
	var create CategoryCreate

	if err := c.ShouldBindJSON(&create); err != nil {
		//fmt.Println("validate error: ", err)
		return &create, err
	}

	// TODO 参数验证
	return &create, nil
}

func (this *Category) Create(data *CategoryCreate) (category *mShopping.Category, err error) {
	category = &mShopping.Category{
		Type: data.Type,
		Pid : data.Pid,
		Name: data.Name,
		Desc: data.Desc,
	}

	model, err := category.Create(category)

	if err != nil {
		return
	}

	category, ok := model.(*mShopping.Category);
		
	if !ok {
		err = ErrStruct
		return
	}

	return
}

type CategoryUpdate struct {
	ID   uint    `json:"id"` 
	Name string  `json:"name"`
	Desc string  `json:"desc"`
}

func (this *Category) ValidateOfUpdate(c *gin.Context) (*CategoryUpdate, error) {
	var update CategoryUpdate

	if err := c.ShouldBindJSON(&update); err != nil {
		return &update, err
	}

	id := c.Param("id")

	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return &update, err
	}

	update.ID = uint(i)

	// TODO 参数验证
	return &update, nil
}

func (this *Category) Update(data *CategoryUpdate) (category *mShopping.Category, err error) {
	category = &mShopping.Category{
		Name: data.Name,
		Desc: data.Desc,
	}

	category.ID = data.ID

	model, err := category.Update(category)

	if err != nil { return }

	category, ok := model.(*mShopping.Category);

	if !ok {
		err = ErrStruct
		return
	}

	return
}

func (this *Category) ValidateOfDelete(c *gin.Context) (uint, error) {
	id := c.Param("id")

	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return 0, err
	}

	return uint(i), nil
}

func (this *Category) Delete(id uint) (category *mShopping.Category, err error) {
	category = &mShopping.Category{}
	category.ID = id

	model, err := category.Delete(category)

	if err != nil { return }

	category, ok := model.(*mShopping.Category);

	if !ok {
		err = ErrStruct
		return
	}

	return
}
