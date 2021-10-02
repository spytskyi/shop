package delivery

import (
	"shop/controllers"

	"github.com/go-pg/pg/v10"
	"github.com/kataras/iris/v12"
)

func NewRouter(app *iris.Application, db *pg.DB) *iris.Application {
	productsAPI := app.Party("/admin")
	{
		// middleware
		productsAPI.Use(iris.Compression)

		productsAPI.Post("/add-new-category", func(ctx iris.Context) {
			controllers.InsertNewCategory(ctx, db)
		})

		productsAPI.Get("/get-square/all-triangles", func(ctx iris.Context) {

		})

		productsAPI.Post("/add-square", func(ctx iris.Context) {

		})
	}

	return app
}
