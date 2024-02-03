package database

import (
	"Redis/redis-impl-go/src/app/utils/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnection(t *testing.T) {
	dbConnect := database.DBConnect()
	assert.NotNil(t, dbConnect)
}
