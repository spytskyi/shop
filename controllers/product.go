package controllers

import (
	"shop/models"
	"shop/repositories"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/kataras/iris/v12"
)

func InsertNewProduct(ctx iris.Context, db *pg.DB) {
	var product models.ProductCreateRequest
	err := ctx.ReadBody(&product)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Product creation failure").DetailErr(err))
		return
	}

	product.Name = strings.Join(strings.Fields(product.Name), " ")

	isExists, err := repositories.RepositoryCheckIsExistsByCode(db, product.VendorCode)
	if isExists {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("This vendor code already exists").DetailErr(err))
		return
	}

	if err != nil && isExists {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title(" creation failure").DetailErr(err))
		return
	}
	err = repositories.RepositoryInsertProduct(db, &product)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Request data is empty").DetailErr(err))
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	_, err = ctx.JSON(iris.Map{"message": "OK"})
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Type product creation failure").DetailErr(err))
		return
	}
	return
}
