package controllers

import (
	"shop/models"
	"shop/repositories"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/kataras/iris/v12"
)

func InsertNewType(ctx iris.Context, db *pg.DB) {
	correctMarks := func(r rune) bool {
		return (r < 'A' || r > 'z') && r != ' '
	}

	var tmpType models.TypeRequest
	err := ctx.ReadBody(&tmpType)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Categories creation failure").DetailErr(err))
		return
	}

	tmpType.Name = strings.Join(strings.Fields(tmpType.Name), " ")

	if strings.IndexFunc(tmpType.Name, correctMarks) != -1 {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Categories creation failure, Found special char").DetailErr(err))
		return
	}

	_, err = repositories.RepositorySelectTypeByName(db, tmpType.Name)
	if err == nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("This type name already exists").DetailErr(err))
		return
	}

	err = repositories.RepositoryInsertType(db, &tmpType)
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

func SelectAllTypes(ctx iris.Context, db *pg.DB) {
	_, err := repositories.RepositorySelectAllTypes(db)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Types selecting failure"))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_, err = ctx.JSON(iris.Map{"message": "OK"})
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Types selecting failure"))
		return
	}
	return
}

func SelectTypeById(ctx iris.Context, db *pg.DB) {
	typeId, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Type("The request failed"))
		return
	}

	_, err = repositories.RepositorySelectTypeById(db, typeId)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("There is no type with the specified id"))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_, err = ctx.JSON(iris.Map{"message": "OK"})
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Types select failure"))
		return
	}
	return
}

func UpdateType(ctx iris.Context, db *pg.DB) {
	correctMarks := func(r rune) bool {
		return r < 'A' || r > 'z'
	}

	var tmpType models.Type

	err := ctx.ReadBody(&tmpType)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Categories update failure").DetailErr(err))
		return
	}

	if strings.IndexFunc(tmpType.Name, correctMarks) != -1 {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Categories update failure, Found special char").DetailErr(err))
		return
	}

	_, err = repositories.RepositorySelectTypeByName(db, tmpType.Name)
	if err == nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("This category name already exists").DetailErr(err))
		return
	}

	_, err = repositories.RepositorySelectTypeById(db, tmpType.Id)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("There is no such category").DetailErr(err))
		return
	}

	err = repositories.RepositoryUpdateType(db, tmpType)
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

//DeleteType ...
func DeleteType(ctx iris.Context, db *pg.DB) {
	var tmpType models.Type
	err := ctx.ReadBody(&tmpType)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Categories remove failure").DetailErr(err))
		return
	}

	//arrTypes, err := repositories.RepositorySelectAllProductsByTypeId(db, tmpType.Id)
	//if err != nil {
	//	ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
	//		Title("There is no such category").DetailErr(err))
	//	return
	//}

	//err = repositories.RepositoryDeleteAllProductsByTypeId(db, tmpType)
	//if err != nil {
	//	ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
	//		Title("Categories remove failure").DetailErr(err))
	//	return
	//}

	//err = repositories.RepositoryDeleteType(db, tmpType)
	//if err != nil {
	//	ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
	//		Title("Categories remove failure").DetailErr(err))
	//	return
	//}

	ctx.StatusCode(iris.StatusOK)
	_, err = ctx.JSON(iris.Map{"message": "OK"})
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Category remove failure").DetailErr(err))
		return
	}
	return
}
