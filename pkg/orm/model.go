package orm

import (
	"github.com/jinzhu/gorm"
	"errors"
	// "fmt"
)

var (
	SearchAll    string = "kelvin_sp_all"
	SearchOne    string = "kelvin_sp_one"
	SearchFields string = "kelvin_sp_fields"
	SearchReturn string = "kelvin_sp_return"
	SearchTables string = "kelvin_sp_tables"
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

	return modelmaker, nil
}

func (this *Model) Update(modelmaker ModelMaker) (int64, error) {
	db, err := modelmaker.GetDb(modelmaker);
	
	if err != nil {
		return 0, err	
	}

	result := db.Model(modelmaker).Update(modelmaker)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
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

	tableName := modelmaker.TableName()
	db = db.Table(tableName)
	db = modelmaker.Condition(data, db)

	if value, ok := data[SearchFields]; ok {
		db = db.Select(value)
	}

	if tables, ok := data[SearchAll]; ok {
		// fmt.Printf("%T", tables)
		// db = db.Find(tables)
		
		if returnStruct, ok2 := data[SearchReturn]; ok2 {
			db = db.Scan(returnStruct)

			if db.Error != nil {
				if db.Error == gorm.ErrRecordNotFound {
					return nil, nil
				}
	
				return nil, db.Error
			}

			return returnStruct, nil
		} else {
			db = db.Find(tables)

			if db.Error != nil {
				if db.Error == gorm.ErrRecordNotFound {
					return nil, nil
				}
	
				return nil, db.Error
			}

			return tables, nil
		}
	}

	if _, ok := data[SearchOne]; ok {
		if returnStruct, ok2 := data[SearchReturn]; ok2 {
			db = db.Limit(1).Scan(returnStruct)

			if db.Error != nil {
				if db.Error == gorm.ErrRecordNotFound {
					return nil, nil
				}
	
				return nil, db.Error
			}

			return returnStruct, nil
		}else {
			db = db.First(modelmaker)

			if db.Error != nil {
				if db.Error == gorm.ErrRecordNotFound {
					return nil, nil
				}
	
				return nil, db.Error
			}

			return db.Value, nil
		}
	}

	return nil, nil
}