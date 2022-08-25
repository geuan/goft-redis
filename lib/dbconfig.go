package lib

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"log"
)
var Gorm *gorm.DB
func init(){
	Gorm=gormDB()
}
func gormDB() *gorm.DB {
	dsn:="root:xc456789110@tcp(192.168.19.138:3306)/mygo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB,err:=db.DB()
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)
	return db
}