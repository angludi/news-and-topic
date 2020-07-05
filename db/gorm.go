package gorm

import (
	"bareksa/config"
	"fmt"

	"github.com/jinzhu/gorm"

	// Register Gorm Mysql Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// Register Go Sql Driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConfig  = config.Config.DB
	mysqlConn *gorm.DB
	err       error
)

func init() {
	if dbConfig.Driver == "mysql" {
		setupMySQLConn()
	}
}

func setupMySQLConn() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	mysqlConn, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	err = mysqlConn.DB().Ping()
	if err != nil {
		panic(err)
	}

	logStatus := true

	mysqlConn.LogMode(logStatus)
}

func MysqlConn() *gorm.DB {
	return mysqlConn
}
