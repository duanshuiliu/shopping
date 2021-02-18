package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"io"

	mShopping "shopping/admin/pkg/models/shopping"
	orm       "shopping/admin/pkg/orm"
	helper    "shopping/admin/pkg/exception"
)

type Category struct {
	BaseService
}

type CategoryListParams struct {
	Pid uint 
}

func (this *Category) ValidateOfList(c *gin.Context) (*CategoryListParams, error) {
	var search CategoryListParams
	
	if err := c.ShouldBindJSON(&search); err != nil {
		if err == io.EOF {
			return &search, nil
		}
		
		return &search, err
	}

	// TODO 参数验证
	return &search, nil
}

type CategoryListRes struct {
	ID          uint   `json:"id"`
	Pid         uint   `json:"pid"`
	Name        string `json:"name"`
	Description string `json:"desc"`
}

func (this *Category) List(data *CategoryListParams) (interface{}, error) {
	var res []CategoryListRes
	category := &mShopping.Category{}

	condition := make(map[string]interface{})
	condition[orm.SearchAll]    = 1
	condition[orm.SearchFields] = []string{"id", "name", "description", "pid"}
	condition[orm.SearchRes]    = &res

	result, err := category.Search(category, condition)
	if err != nil { return result, err }

	return result, nil
}

func (this *Category) ValidateOfShow(c *gin.Context) (uint, error) {
	id := c.Param("id")

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil { return 0, err }

	return uint(i), nil
}

type CategoryDetailRes struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
}

func (this *Category) Show(id uint) (interface{}, error) {
	category := &mShopping.Category{}

	condition := make(map[string]interface{})
	condition["id"]             = id
	condition[orm.SearchOne]    = 1
	condition[orm.SearchFields] = []string{"id", "name", "description"}
	condition[orm.SearchRes]    = &CategoryDetailRes{}

	result, err := category.Search(category, condition)
	if err != nil { return result, err }

	if result == nil { return result, &helper.MsgException{Message: "not found data"} }
	return result, nil
}

type CategoryCreateParams struct {
	Type        int8    `json:"type"`
	Pid         uint    `json:"pid"`
	Name        string  `json:"name"`
	Description string  `json:"desc"`
}

func (this *Category) ValidateOfCreate(c *gin.Context) (*CategoryCreateParams, error) {
	var create CategoryCreateParams

	if err := c.ShouldBindJSON(&create); err != nil {
		return &create, err
	}

	// TODO 参数验证
	return &create, nil
}

func (this *Category) Create(data *CategoryCreateParams) (*mShopping.Category, error) {
	category := &mShopping.Category{
		Type       : data.Type,
		Pid        : data.Pid,
		Name       : data.Name,
		Description: data.Description,
	}

	result, err := category.Create(category)
	if err != nil { return nil, err }

	m, ok := result.(*mShopping.Category)
	if ok { return m, nil }
	return nil, errors.New("insert error")
}

type CategoryUpdateParams struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"desc"`
}

func (this *Category) ValidateOfUpdate(c *gin.Context) (*CategoryUpdateParams, error) {
	var update CategoryUpdateParams

	if err := c.ShouldBindJSON(&update); err != nil {
		return &update, err
	}

	id     := c.Param("id")
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil { return &update, err }

	update.ID = uint(i)

	// TODO 参数验证
	return &update, nil
}

func (this *Category) Update(data *CategoryUpdateParams) (int64, error) {
	category := &mShopping.Category{
		Name       : data.Name,
		Description: data.Description,
	}
	category.ID = data.ID

	rows, err := category.Update(category)
	if err != nil { return 0, err }
	if rows == 0 { return 0, &helper.MsgException{Message: "not found data"} }

	return rows, nil
}

func (this *Category) ValidateOfDelete(c *gin.Context) (uint, error) {
	id     := c.Param("id")
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil { return 0, err }

	return uint(i), nil
}

func (this *Category) Delete(id uint) (int64, error) {
	category := &mShopping.Category{}
	category.ID = id

	rows, err := category.Delete(category)
	if err != nil { return 0, err }
	if rows == 0 { return 0, &helper.MsgException{Message: "not found data"} }

	return rows, nil
}
