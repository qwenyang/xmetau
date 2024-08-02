package tables

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mysqlDSN string   = fmt.Sprintf("xmetau:%s@(%s:%d)/xmetau?charset=utf8mb4&parseTime=True&loc=Local", DataBasePassword, DataBaseIP, DataBasePort)
	globalDB *gorm.DB = nil
)

func openDB() (*gorm.DB, error) {
	if globalDB != nil {
		return globalDB, nil
	}
	log.Printf("created db %o", globalDB)
	db, err := gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{})
	sqlDB, err := db.DB()
	if err != nil {
		sqlDB.SetMaxIdleConns(25)  //空闲连接数
		sqlDB.SetMaxOpenConns(600) //最大连接数
		sqlDB.SetConnMaxLifetime(30)
	}
	globalDB = db
	return globalDB, err
}
