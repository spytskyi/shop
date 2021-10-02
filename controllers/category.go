package controllers

import (
	"fmt"
	"shop/models"
	"shop/repositories"

	"github.com/go-pg/pg/v10"
	"github.com/kataras/iris/v12"
)

func InsertNewCategory(ctx iris.Context, db *pg.DB) {

	var request models.CategoryRequest
	fmt.Println(ctx)
	err := ctx.ReadBody(&request)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Triangle creation failure").DetailErr(err))
		return
	}

	err = repositories.CreateCategory(db, &request)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Triangle creation failure").DetailErr(err))
		return
	}
	ctx.StatusCode(iris.StatusCreated)
	_, err = ctx.JSON(iris.Map{"message": "OK"})
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Triangle creation failure").DetailErr(err))
		return
	}
	return
}
