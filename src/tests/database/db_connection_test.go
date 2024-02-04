package database

import (
	"Redis/redis-impl-go/src/app/model/domain"
	"Redis/redis-impl-go/src/app/utils/database"
	"Redis/redis-impl-go/src/app/utils/errors"
	"Redis/redis-impl-go/src/app/utils/response"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var (
	dbConnect   = database.MysqlConnect()
	redisClient = database.RedisClient()
	ctx         = context.Background()
)

func TestMysqlConnection(t *testing.T) {
	assert.NotNil(t, dbConnect)
}

func TestRedisConnection(t *testing.T) {
	assert.NotNil(t, redisClient)

	pong := redisClient.Ping(ctx).Val()
	log.Println(pong)
	assert.Equal(t, "PONG", pong)

	err := redisClient.Set(ctx, "isSave", true, 0).Err()
	errors.PanicIfError(err)

	isSave := redisClient.Get(ctx, "isSave").Val()
	log.Println(isSave)
	assert.Equal(t, "1", isSave)
	//assert.Equal(t, "true", isSave)
}

type ExampleRedis struct {
	Id   int    `redis:"idEx"`
	Name string `redis:"name"`
}

func TestJSONRedis(t *testing.T) {

	courses := []domain.Course{}
	err := dbConnect.Limit(2).Find(&courses).Error
	errors.PanicIfError(err)

	marshal, _ := json.Marshal(response.ToCourseResponses(courses))
	log.Println(string(marshal))

	err = redisClient.Set(ctx, "courses", string(marshal), 0).Err()
	errors.PanicIfError(err)

	val := redisClient.Get(ctx, "courses").Val()
	log.Println(val)
}
