package repositories

import (
	"fmt"
	"shop/models"
	"strings"

	"github.com/go-pg/pg/v10"
)

//RepositoryCreateCategory ...
func RepositoryCreateCategory(db *pg.DB, reqData *models.CategoryRequest) (err error) {
	newCategory := models.Category{
		Name: strings.ToLower(reqData.Name),
	}

	_, err = db.Model(&newCategory).Insert()
	if err != nil {
		return
	}

	return nil
}

//RepositoryGetAllCategories ...
func RepositoryGetAllCategories(db *pg.DB) (arrCategories []models.Category, err error) {

	err = db.Model(&arrCategories).Select()

	if err != nil {
		return nil, err
	}

	return arrCategories, nil
}

//RepositoryGetCategoryById ...
func RepositoryGetCategoryById(db *pg.DB, categoryId int) (category models.Category, err error) {

	category.Id = categoryId

	err = db.Model(&category).WherePK().Select()

	if err != nil {
		return category, err
	}

	return category, err
}

func RepositoryGetCategoryByName(db *pg.DB, categoryName string) (category models.Category, err error) {
	category.Name = strings.ToLower(categoryName)
	err = db.Model(&category).Where("name = ?", category.Name).Select()
	fmt.Println(err)
	if err != nil {
		return category, err
	}
	return category, err
}

//RepositoryUpdateCategory ...
func RepositoryUpdateCategory(db *pg.DB, category models.Category) (err error) {
	category.Name = strings.ToLower(category.Name)
	_, err = db.Model(&category).WherePK().Update()
	if err != nil {
		return err
	}

	return err
}

//RepositoryRemoveCategory ...
func RepositoryRemoveCategory(db *pg.DB, category models.Category) (err error) {
	res, err := db.Model(&category).WherePK().Delete()
	if err != nil {
		return err
	}
	fmt.Println("deleted", res.RowsAffected())
	return err
}
