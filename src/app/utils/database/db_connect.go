package database

import (
	"Redis/redis-impl-go/src/app/utils/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/redis?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	errors.PanicIfError(err)

	return db
}
