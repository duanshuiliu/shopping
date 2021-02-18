package orm

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"errors"
	"time"

	config "shopping/admin/pkg/conf"
)

var (
	Instance *DbConns

	ErrEmptyConn = errors.New("not found connection")
)

type DbConns struct {
	ConnMap map[string]*gorm.DB
}

type DbConfig struct {
	Driver       string
	Host         string
	Port         int
	Dbname       string
	Username     string
	Password     string
	MaxIdleConns int
	MaxOpenConns int
}

func Register() error {
	dbsConf, err := getDbConfig()
	if err != nil { return err }

	Instance = &DbConns{
		ConnMap: make(map[string]*gorm.DB),
	}

	for dbName := range dbsConf {
		dbConf := dbsConf[dbName]
		dsn := dbConf.Username+":"+dbConf.Password+"@tcp("+dbConf.Host+":"+strconv.Itoa(dbConf.Port)+")/"+dbConf.Dbname+"?parseTime=true"
		conn, err := sql.Open(dbConf.Driver, dsn)
		if err != nil { return err }

		conn.SetConnMaxLifetime(time.Minute*3)
		conn.SetMaxOpenConns(dbConf.MaxOpenConns)
		conn.SetMaxIdleConns(dbConf.MaxIdleConns)

		db, err := gorm.Open(mysql.New(mysql.Config{
			Conn: conn,
		}), &gorm.Config{})

		if err != nil { return err }
		Instance.ConnMap[dbName] = db
	}

	return nil
}

func New(connName string) (*gorm.DB, error) {
	if _, ok := Instance.ConnMap[connName]; !ok {
		return nil, ErrEmptyConn
	}

	return Instance.ConnMap[connName], nil
}

func getDbConfig() (map[string]*DbConfig, error) {
	dbConf, err := config.New("database")
	if err != nil { return nil, err }

	c := make(map[string]*DbConfig)
	err = dbConf.Unmarshal(&c)
	if err != nil { return nil, err }
	return c, nil
}
