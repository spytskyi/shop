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

		productsAPI.Post("/add-new-product", func(ctx iris.Context) {
			controllers.InsertNewProduct(ctx, db)
		})

		//InsertNewCategory ...
		productsAPI.Post("/add-new-category", func(ctx iris.Context) {
			controllers.InsertNewCategory(ctx, db)
		})

		//InsertNewType ...
		productsAPI.Post("/add-new-type", func(ctx iris.Context) {
			controllers.InsertNewType(ctx, db)
		})

		//SelectAllCategories ...
		productsAPI.Get("/get-all-categories", func(ctx iris.Context) {
			controllers.SelectAllCategories(ctx, db)
		})

		//SelectAllTypes ...
		productsAPI.Get("/get-all-types", func(ctx iris.Context) {
			controllers.SelectAllTypes(ctx, db)
		})

		//SelectCategoryById
		productsAPI.Get("/get-category-id/{id:int}", func(ctx iris.Context) {
			controllers.SelectCategoryById(ctx, db)
		})

		//SelectTypeById ...
		productsAPI.Get("/get-type-id/{id:int}", func(ctx iris.Context) {
			controllers.SelectTypeById(ctx, db)
		})
		//UpdateCategory ...
		productsAPI.Patch("/update-category", func(ctx iris.Context) {
			controllers.UpdateCategory(ctx, db)
		})

		//UpdateType ...
		productsAPI.Patch("/update-type", func(ctx iris.Context) {
			controllers.UpdateType(ctx, db)
		})

		//DeleteCategory ...
		productsAPI.Delete("/delete-category", func(ctx iris.Context) {
			controllers.DeleteCategory(ctx, db)
		})

		//DeleteType ...
		productsAPI.Delete("/delete-type", func(ctx iris.Context) {
			controllers.DeleteType(ctx, db)
		})

	}

	return app
}
