package repositories

import (
	"shop/models"
	"strings"

	"github.com/go-pg/pg/v10"
)

//RepositoryInsertCategory ...
func RepositoryInsertCategory(db *pg.DB, reqData *models.CategoryRequest) (err error) {
	newCategory := models.Category{
		Name: strings.ToLower(reqData.Name),
	}

	_, err = db.Model(&newCategory).Insert()
	if err != nil {
		return
	}

	return nil
}

//RepositorySelectAllCategories ...
func RepositorySelectAllCategories(db *pg.DB) (arrCategories []models.Category, err error) {

	err = db.Model(&arrCategories).Select()

	if err != nil {
		return nil, err
	}

	return arrCategories, nil
}

//RepositorySelectCategoryById ...
func RepositorySelectCategoryById(db *pg.DB, categoryId int) (category models.Category, err error) {

	category.Id = categoryId

	err = db.Model(&category).WherePK().Select()

	if err != nil {
		return category, err
	}

	return category, err
}

func RepositoryCheckIsExistsByName(db *pg.DB, categoryName string) (isExists bool, err error) {
	isExists, err = db.Model(&models.Category{}).Where("name = ?", strings.ToLower(categoryName)).Exists()
	if err != nil {
		return isExists, err
	}
	return isExists, err
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

//RepositoryDeleteCategory ...
func RepositoryDeleteCategory(db *pg.DB, category models.Category) (err error) {
	_, err = db.Model(&category).WherePK().Delete()
	if err != nil {
		return err
	}
	return err
}
