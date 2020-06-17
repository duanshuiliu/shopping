package orm

import (
	"github.com/jinzhu/gorm"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"errors"
	//"fmt"

	config "shopping/utils/conf"
)

type Orm struct {
	ConnMap map[string]*gorm.DB
}

type dbConfig struct {
	name         string
	driver       string
	host         string
	port         int
	dbname       string
	username     string
	password     string
	//charset      string
	maxIdleConns int
	maxOpenConns int	
}

var Instance *Orm

func Register() (err error) {
	dbsConf, err := getDbConfig()

	if err != nil {
		return
	}

	Instance = &Orm{
		ConnMap: make(map[string]*gorm.DB),
	}

	for _, dbconf := range dbsConf {
		dsn := dbconf.username+":"+dbconf.password+"@tcp("+dbconf.host+":"+strconv.Itoa(dbconf.port)+")/"+dbconf.dbname+"?parseTime=true"
		db, err1 := gorm.Open(dbconf.driver, dsn)

		if err1 != nil {
			err = err1
			return
		}

		db.DB().SetMaxIdleConns(dbconf.maxIdleConns)
		db.DB().SetMaxOpenConns(dbconf.maxOpenConns)
		
		Instance.ConnMap[dbconf.name] = db
	}

	return 
}

func New(connection string) (db *gorm.DB, err error) {
	if _, ok := Instance.ConnMap[connection]; !ok {
		err = errors.New("未找到该数据库连接，请确认")
		return
	}

	db = Instance.ConnMap[connection]
	return
}

func getDbConfig() (data []dbConfig, err error) {	
	dbconf, err := config.New("database")

	if err != nil {
		return
	}

	allconfs := dbconf.All()

	if m, ok := allconfs.(map[string]map[string]string); ok {
		for section := range m {
			if section != "DEFAULT" {
				driver       := m[section]["driver"]
				host         := m[section]["host"]
				port         := m[section]["port"]
				dbname       := m[section]["dbname"]
				username     := m[section]["username"]
				password     := m[section]["password"]
				//charset      := m[section]["charset"]
				maxIdleConns := m[section]["maxIdleConns"]
				maxOpenConns := m[section]["maxOpenConns"]
				//fmt.Println(section, driver, host, port, dbname, username, password, charset, maxIdleConns, maxOpenConns)	

				portInt,_         := strconv.Atoi(port)
				maxIdleConnsInt,_ := strconv.Atoi(maxIdleConns)
				maxOpenConnsInt,_ := strconv.Atoi(maxOpenConns)

				data = append(data, dbConfig{
					name         : section,
					driver       : driver,
					host         : host,
					port         : portInt,
					dbname       : dbname,
					username     : username,
					password     : password,
					//charset      : charset,
					maxIdleConns : maxIdleConnsInt,
					maxOpenConns : maxOpenConnsInt,
				})
			}
		}
	}

	return
}
