package services

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"

	pError    "shopping/pkg/error"
	mShopping "shopping/pkg/models/shopping"
	orm       "shopping/pkg/orm"
)

type Category struct {
	BaseService
}

type CategorySearch struct {
	Pid uint 
}

func (this *Category) ValidateOfList(c *gin.Context) (*CategorySearch, error) {
	var search CategorySearch
	
	if err := c.ShouldBindJSON(&search); err != nil {
		return &search, err
	}

	// TODO 参数验证
	return &search, nil
}

type CategoryResponseList struct {
	ID   uint   `json:"id"`
	Pid  uint   `json:"pid"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (this *Category) List(data *CategorySearch) (interface{}, error) {
	category := &mShopping.Category{}

	var res    []CategoryResponseList
	var tables []mShopping.Category

	condition := make(map[string]interface{})
	condition[orm.SearchAll]    = &tables
	condition[orm.SearchFields] = []string{"id", "name", "pid"}
	condition[orm.SearchReturn] = &res

	result, err := category.Search(category, condition)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (this *Category) ValidateOfShow(c *gin.Context) (uint, error) {
	id := c.Param("id")

	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return 0, err
	}

	return uint(i), nil
}

type CategoryResponseDetail struct {
	ID   uint   `json:"id"`
	Pid  uint   `json:"pid"`
	Name string `json:"name"`
	Desc string `json:"desc"` 
}

func (this *Category) Show(id uint) (interface{}, error) {
	category := &mShopping.Category{}

	condition := make(map[string]interface{})
	condition["id"]             = id
	condition[orm.SearchOne]    = 1
	condition[orm.SearchFields] = []string{"id", "name", "pid"}
	condition[orm.SearchReturn] = &CategoryResponseDetail{}

	result, err := category.Search(category, condition)

	if err != nil {
		return result, err
	}

	if result == nil {
		return result, &pError.MessageError{Message: "not found data"}
	}

	fmt.Println("%T", result)
	return result, nil
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

func (this *Category) Update(data *CategoryUpdate) (rows int64, err error) {
	category := &mShopping.Category{
		Name: data.Name,
		Desc: data.Desc,
	}

	category.ID = data.ID

	rows, err = category.Update(category)
	if err != nil { return }

	if rows == 0 {
		err = &pError.MessageError{Message: "not found data"}
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

func (this *Category) Delete(id uint) (rows int64, err error) {
	category := &mShopping.Category{}
	category.ID = id

	rows, err = category.Delete(category)
	if err != nil { return }
	
	if rows == 0 {
		err = &pError.MessageError{Message: "not found data"}
		return
	}

	return
}
