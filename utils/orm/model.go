package orm

import (
	"github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model
}

type Modeler interface {
	// 
}

func (this *Model) Connection() string {
	return "default"
}

func (this *Model) GetDb() (db *gorm.DB, err error) {
	// 获取连接名
	conn := this.Connection()
	// 获取数据库连接信息
	return New(conn)
} 

func (this *Model) Create(model Modeler) (model Modeler, err error) {
	if db, err := this.GetDb(); err != nil {
		return
	}

	// 插入数据
	db.Create(model)
	return
}

func (this *Model) Update(model Modeler) (model Modeler, err error) {
	if db, err := this.GetDb(); err != nil {
		return
	}

	db.Save(model)
	return
}

func (this *Model) Delete(model Modeler) (model Modeler, err error) {
	if db, err := this.GetDb(); err != nil {
		return
	}

	db.Delete(model)
	return
}

func (this *Model) Search(model Modeler) {
	// 
} 
