package main

import (
	"Redis/redis-impl-go/src/app/utils/api"
	"Redis/redis-impl-go/src/app/utils/errors"
)

func main() {
	app := api.NewRouter()

	err := app.Listen("localhost:3000")
	errors.PanicIfError(err)
}
