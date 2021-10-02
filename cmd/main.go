package main

import (
	"context"
	"shop/conf"
	"shop/delivery"
	"shop/postgres"

	"github.com/kataras/iris/v12"
)

func main() {
	conf.InitConfig()

	db := postgres.InitPostgres()

	defer db.Close()

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	app := iris.Default()

	app = delivery.NewRouter(app, db)
	// This handler will match /user/john but will not match /user/ or /user

	app.Listen(":8080")
}
