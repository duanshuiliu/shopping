package orm

import (
	"github.com/jinzhu/gorm"
	"errors"
	"fmt"
)

var (
	SearchAll    string = "kelvin_sp_all"
	SearchOne    string = "kelvin_sp_one"
	SelectFields string = "kelvin_sp_fields"
)

var ErrStruct = errors.New("invalid struct")

type Model struct {
	gorm.Model
}

type ModelMaker interface {
	Connection() string
	GetDb(ModelMaker) (*gorm.DB, error) 
	TableName() string 
	Condition(map[string]interface{}, *gorm.DB) *gorm.DB
}

func (this *Model) Connection() string {
	return "default"
}

func (this *Model) GetDb(modelmaker ModelMaker) (db *gorm.DB, err error) {
	conn := modelmaker.Connection()
	return New(conn)
}

func (this *Model) TableName() string {
	return "models"
}

func (this *Model) Condition(data map[string]interface{}, db *gorm.DB) *gorm.DB {
	if value, ok := data["id"]; ok {
		db = db.Where("id = ?", value)
	}

	return db
}

func (this *Model) Create(modelmaker ModelMaker) (ModelMaker, error) {
	db, err := modelmaker.GetDb(modelmaker)

	if err != nil {
		return modelmaker, err
	}

	result := db.Create(modelmaker)

	if result.Error != nil {
		return modelmaker, result.Error
	}

	modelmaker, ok := result.Value.(ModelMaker)

	if !ok {
		return modelmaker, ErrStruct
	}

	//fmt.Println("orm result: ", result.Value, result.Error, result.RowsAffected)
	return modelmaker, nil
}

func (this *Model) Update(modelmaker ModelMaker) (ModelMaker, error) {
	db, err := modelmaker.GetDb(modelmaker);
	
	if err != nil {
		return modelmaker, err	
	}

	result := db.Model(modelmaker).Update(modelmaker)

	fmt.Println("enter update func", modelmaker)

	if result.Error != nil {
		return modelmaker, result.Error
	}

	modelmaker, ok := result.Value.(ModelMaker)

	if !ok {
		return modelmaker, ErrStruct
	}

	return modelmaker, nil
}

func (this *Model) Delete(modelmaker ModelMaker) (int64, error) {
	db, err := modelmaker.GetDb(modelmaker)

	if err != nil {
		return 0, err
	}

	result := db.Delete(modelmaker)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (this *Model) Search(modelmaker ModelMaker, data map[string]interface{}) (interface{}, error) {
	db, err := modelmaker.GetDb(modelmaker)

	if err != nil {
		return nil, err
	}

	db = modelmaker.Condition(data, db)

	if value, ok := data[SelectFields]; ok {
		db = db.Select(value)
	}

	if _, ok := data[SearchAll]; ok {
		result := db.Find(modelmaker)

		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				return nil, nil
			}

			return nil, result.Error
		}

		return result.Value, nil
	}

	if _, ok := data[SearchOne]; ok {
		result := db.First(modelmaker)

		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				return nil, nil
			}

			return nil, result.Error
		}
		
		return result.Value, nil
	}

	return nil, nil
}