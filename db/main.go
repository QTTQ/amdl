package db

import (
	"amdl/config"
	"errors"
	"fmt"

	// aaaaa
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

var (
	// DB aaaa
	DB *gorm.DB
	// ErrDBNull aaa
	ErrDBNull = errors.New("db is null,please connect first")
)

//ConnectAndInit aaaa
func ConnectAndInit(conf config.ConfStruct, models ...interface{}) {
	dbUser := conf.DB.User
	dbPass := conf.DB.Password
	dbHost := conf.DB.Host
	dbName := conf.DB.Name
	dbPort := fmt.Sprintf("%d", conf.DB.Port)
	// dbPrefix := conf.DB.Prefix
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		// if defaultTableName == "meta" {
		// 	defaultTableName = "metas"
		// }
		// return dbPrefix + defaultTableName
		return defaultTableName
	}
	Connect(dbUser, dbPass, dbHost, dbPort, dbName)
	Init(models...)
}

// Init aaa
func Init(models ...interface{}) error {
	if DB == nil {
		return ErrDBNull
	}
	DB.AutoMigrate(models...)
	return nil
}

// Connect aaa
func Connect(user, password, host string, dbPort, dbName string) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbPort, dbName))
	if err != nil {
		panic("connect database failed" + err.Error())
	}
	db.SingularTable(true) //取消建表为复数形式
	DB = db
}
