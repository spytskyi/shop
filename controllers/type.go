package controllers

import (
	"errors"
	"shop/models"
	"shop/repositories"

	"github.com/go-pg/pg/v10"
	"github.com/kataras/iris/v12"
)

func InsertNewType(ctx iris.Context, db *pg.DB) {

	var request models.TypeRequest
	err := ctx.ReadBody(&request)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Triangle creation failure").DetailErr(err))
		return
	}

	if request.Name == "" || request.CategoryId == 0 {
		err = errors.New("Request data is empty")
	}
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Request data is empty").DetailErr(err))
		return
	}

	err = repositories.CreateType(db, &request)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Request data is empty").DetailErr(err))
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
