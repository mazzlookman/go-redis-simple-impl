package database

import (
	"Redis/redis-impl-go/src/app/utils/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMysqlConnection(t *testing.T) {
	dbConnect := database.MysqlConnect()
	assert.NotNil(t, dbConnect)
}

func TestRedisConnection(t *testing.T) {
	redisClient := database.RedisClient()
	assert.NotNil(t, redisClient)
}
