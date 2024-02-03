package database

import (
	"Redis/redis-impl-go/src/app/utils/errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConnect() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/redis?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	errors.PanicIfError(err)

	return db
}

func RedisClient() *redis.Client {
	rclient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	return rclient
}
