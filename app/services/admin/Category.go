package admin

import (
	"github.com/gin-gonic/gin"
	//"fmt"
	"errors"

	//orm    "shopping/utils/orm"
	sModel "shopping/app/models/shopping"
)

type Category struct {
	BaseService
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

func (this *Category) Create(data *CategoryCreate) (category *sModel.Category, err error) {
	category = &sModel.Category{
		Type: data.Type,
		Pid : data.Pid,
		Name: data.Name,
		Desc: data.Desc,
	}

	model, err := category.Create(category)

	if err != nil {
		return
	}

	category, ok := model.(*sModel.Category);
		
	if !ok {
		err = errors.New("error struct type")
		return
	}

	return
}
