package controllers

import (
	"shop/models"
	"shop/repositories"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/kataras/iris/v12"
)

//InsertNewCategory ...
func InsertNewCategory(ctx iris.Context, db *pg.DB) {
	correctMarks := func(r rune) bool {
		return r < 'A' || r > 'z'
	}

	var category models.CategoryRequest
	err := ctx.ReadBody(&category)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Categories creation failure").DetailErr(err))
		return
	}

	if strings.IndexFunc(category.Name, correctMarks) != -1 {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Categories creation failure, Found special char").DetailErr(err))
		return
	}

	isExists, err := repositories.RepositoryCheckIsExistsByName(db, category.Name)
	if isExists {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("This category name already exists").DetailErr(err))
		return
	}

	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Categories creation failure").DetailErr(err))
		return
	}

	err = repositories.RepositoryInsertCategory(db, &category)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Categories creation failure").DetailErr(err))
		return
	}
	ctx.StatusCode(iris.StatusCreated)
	_, err = ctx.JSON(iris.Map{"message": "OK"})
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Categories creation failure").DetailErr(err))
		return
	}
	return
}

//SelectAllCategories ...
func SelectAllCategories(ctx iris.Context, db *pg.DB) {
	_, err := repositories.RepositorySelectAllCategories(db)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Categories selecting failure").DetailErr(err))
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	_, err = ctx.JSON(iris.Map{"message": "OK"})
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Categories selecting failure").DetailErr(err))
		return
	}
	return
}

//SelectCategoryById ...
func SelectCategoryById(ctx iris.Context, db *pg.DB) {
	categoryId, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Category select failure").DetailErr(err))
		return
	}

	_, err = repositories.RepositorySelectCategoryById(db, categoryId)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Category select failure").DetailErr(err))
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	_, err = ctx.JSON(iris.Map{"message": "OK"})
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Category select failure").DetailErr(err))
		return
	}

	return
}

//UpdateCategory ...
func UpdateCategory(ctx iris.Context, db *pg.DB) {
	correctMarks := func(r rune) bool {
		return r < 'A' || r > 'z'
	}

	var category models.Category

	err := ctx.ReadBody(&category)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Categories update failure").DetailErr(err))
		return
	}

	if strings.IndexFunc(category.Name, correctMarks) != -1 {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Categories update failure, Found special char").DetailErr(err))
		return
	}

	isExists, err := repositories.RepositoryCheckIsExistsByName(db, category.Name)
	if isExists {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("This category name already exists").DetailErr(err))
		return
	}

	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Categories creation failure").DetailErr(err))
		return
	}

	_, err = repositories.RepositorySelectCategoryById(db, category.Id)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("There is no such category").DetailErr(err))
		return
	}

	err = repositories.RepositoryUpdateCategory(db, category)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Categories update failure").DetailErr(err))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_, err = ctx.JSON(iris.Map{"message": "OK"})
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Category update failure").DetailErr(err))
		return
	}

	return
}

//DeleteCategory ...
func DeleteCategory(ctx iris.Context, db *pg.DB) {
	var category models.Category
	err := ctx.ReadBody(&category)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Categories remove failure").DetailErr(err))
		return
	}

	arrTypes, err := repositories.RepositorySelectAllTypesByCategoryId(db, category.Id)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("There is no such category").DetailErr(err))
		return
	}

	err = repositories.RepositoryDeleteAllTypesByCategoryId(db, arrTypes)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Categories remove failure").DetailErr(err))
		return
	}

	err = repositories.RepositoryDeleteCategory(db, category)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Categories remove failure").DetailErr(err))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_, err = ctx.JSON(iris.Map{"message": "OK"})
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Category remove failure").DetailErr(err))
		return
	}
	return
}
