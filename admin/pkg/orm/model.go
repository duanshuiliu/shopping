package orm

import (
	"gorm.io/gorm"
)

var (
	SearchAll    string = "kelvin_sp_all"
	SearchOne    string = "kelvin_sp_one"
	SearchFields string = "kelvin_sp_fields"
	SearchRes    string = "kelvin_sp_res"
)

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

func (this *Model) TableName() string {
	return "models"
}

func (this *Model) GetDb(m ModelMaker) (*gorm.DB, error) {
	conn := m.Connection()
	return New(conn)
}

func (this *Model) Condition(data map[string]interface{}, db *gorm.DB) *gorm.DB {
	return db
}

func (this *Model) Create(m ModelMaker) (ModelMaker, error) {
	db, err := m.GetDb(m)
	if err != nil { return nil, err }

	result := db.Create(m)
	if result.Error != nil { return nil, result.Error }
	return m, nil
}

func (this *Model) Update(m ModelMaker) (int64, error) {
	db, err := m.GetDb(m)
	if err != nil { return 0, err }

	result := db.Model(m).Updates(m)
	if result.Error != nil { return 0, result.Error }
	return result.RowsAffected, nil
}

func (this *Model) Delete(m ModelMaker) (int64, error) {
	db, err := m.GetDb(m)
	if err != nil { return 0, err }

	result := db.Delete(m)
	if result.Error != nil { return 0, result.Error }
	return result.RowsAffected, nil
}

func (this *Model) Search(m ModelMaker, data map[string]interface{}) (interface{}, error) {
	db, err := m.GetDb(m)
	if err != nil { return nil, err }

	tableName := m.TableName()
	db = db.Table(tableName)
	db = m.Condition(data, db)

	// 查询具体的字段
	if value, ok := data[SearchFields]; ok {
		db = db.Select(value)
	}

	// 查询所有
	if _, ok := data[SearchAll]; ok {
		if s, ok2 := data[SearchRes]; ok2 {
			db = db.Find(s)

			if db.Error != nil {
				if db.Error == gorm.ErrRecordNotFound { return nil, nil }
				return nil, db.Error
			}

			return s, nil
		} else {
			var res []map[string]interface{}
			db = db.Find(&res)

			if db.Error != nil {
				if db.Error == gorm.ErrRecordNotFound { return nil, nil }
				return nil, db.Error
			}

			return res, nil
		}
	}

	// 查询单条
	if _, ok := data[SearchOne]; ok {
		if s, ok2 := data[SearchRes]; ok2 {
			db = db.First(s)

			if db.Error != nil {
				if db.Error == gorm.ErrRecordNotFound { return nil, nil }
				return nil, db.Error
			}

			return s, nil
		} else {
			res := make(map[string]interface{})
			db = db.Take(&res)

			if db.Error != nil {
				if db.Error == gorm.ErrRecordNotFound { return nil, nil }
				return nil, db.Error
			}

			return res, nil
		}
	}

	return nil, nil
}