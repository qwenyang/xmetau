package tables

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	globalDB *gorm.DB = nil
)

func openDB() (*gorm.DB, error) {
	if globalDB != nil {
		return globalDB, nil
	}
	var mysqlDSN string = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", MainConfig.DataBaseUserName, MainConfig.DataBasePassword, MainConfig.DataBaseIP, MainConfig.DataBasePort, MainConfig.DataBaseName)
	log.Printf("created db %o", globalDB)
	db, err := gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{})
	log.Printf("created db %o", globalDB, err)
	sqlDB, err := db.DB()
	if err != nil {
		sqlDB.SetMaxIdleConns(25)  //空闲连接数
		sqlDB.SetMaxOpenConns(600) //最大连接数
		sqlDB.SetConnMaxLifetime(30)
	}
	globalDB = db
	return globalDB, err
}
