package orm

import (
	"github.com/jinzhu/gorm"
	"errors"
	//"fmt"
)

var ErrStruct = errors.New("invalid struct")

type Model struct {
	gorm.Model
}

type ModelMaker interface {
	Connection() string
	GetDb(ModelMaker) (*gorm.DB, error) 
	TableName() string 
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

	if result.Error != nil {
		return modelmaker, result.Error
	}

	modelmaker, ok := result.Value.(ModelMaker)

	if !ok {
		return modelmaker, ErrStruct
	}

	return modelmaker, nil
}

func (this *Model) Delete(modelmaker ModelMaker) (ModelMaker, error) {
	db, err := modelmaker.GetDb(modelmaker)

	if err != nil {
		return modelmaker, err
	}

	result := db.Delete(modelmaker)

	modelmaker, ok := result.Value.(ModelMaker)

	if !ok {
		return modelmaker, ErrStruct
	}

	return modelmaker, nil
}

// func (this *Model) SearchAll(modelmaker ModelMaker) {
// 	db, err := modelmaker.GetDb(modelmaker)

// 	if err != nil {
// 		return modelmaker, err
// 	}

// 	result := db.Where(modelmaker).	
// }

// func (this *Model) SearchOne(modelmaker ModelMaker) (ModelMaker, error) {
// 	db, err := modelmaker.GetDb(modelmaker)

// 	if err != nil {
// 		return modelmaker, err
// 	}

// 	result := db.Where(modelmaker).First(modelmaker)

// 	modelmaker, ok := result.Value.(ModelMaker)

// 	if !ok {
// 		return modelmaker, ErrStruct
// 	}

// 	return modelmaker, nil
// } 
