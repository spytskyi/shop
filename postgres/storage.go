package postgres

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
)

func InitPostgres() *pg.DB {

	db := pg.Connect(&pg.Options{
		Addr:     viper.GetString("db.host") + viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.pass"),
		Database: viper.GetString("db.name"),
	})
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}
	return db
}

