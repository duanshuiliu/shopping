package orm

import (
	"github.com/jinzhu/gorm"
	"errors"
	//"fmt"
)

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
	db, err := modelmaker.GetDb(modelmaker);
	if err != nil {
		return modelmaker, err
	}

	result := db.Create(modelmaker)

	if result.Error != nil {
		return modelmaker, result.Error
	}

	modelmaker, ok := result.Value.(ModelMaker)

	if !ok {
		return modelmaker, errors.New("error struct type")
	}

	//fmt.Println("orm result: ", result.Value, result.Error, result.RowsAffected)
	return modelmaker, nil
}

//func (this *Model) Update(modelmaker ModelMaker) (modelmaker ModelMaker, err error) {
//	if db, err := modelmaker.GetDb(); err != nil {
//		return
//	}
//
//	db.Save(model)
//	return
//}

//func (this *Model) Delete(model Modeler) (model Modeler, err error) {
//	if db, err := this.GetDb(); err != nil {
//		return
//	}
//
//	db.Delete(model)
//	return
//}
//
//func (this *Model) Search(model Modeler) {
//	// 
//} 
