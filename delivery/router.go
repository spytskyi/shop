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

		productsAPI.Get("/get-all-categories", func(ctx iris.Context) {
			controllers.SelectAllCategories(ctx, db)
		})

		productsAPI.Get("/get-category-id/{id:int}", func(ctx iris.Context) {
			controllers.SelectCategoryById(ctx, db)
		})

		productsAPI.Delete("/delete-category", func(ctx iris.Context) {
			controllers.DeleteCategory(ctx, db)
		})

		productsAPI.Patch("/update-category", func(ctx iris.Context) {
			controllers.UpdateCategory(ctx, db)
		})

		productsAPI.Post("/add-new-type", func(ctx iris.Context) {
			controllers.InsertNewType(ctx, db)
		})

		productsAPI.Post("/add-square", func(ctx iris.Context) {

		})
	}

	return app
}
