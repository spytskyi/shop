package repositories

import (
	"shop/models"

	"github.com/go-pg/pg/v10"
)

func CreateCategory(db *pg.DB, reqData *models.CategoryRequest) (err error) {
	newCategory := models.Category{
		Name: reqData.Name,
	}

	_, err = db.Model(&newCategory).Insert()
	if err != nil {
		return
	}

	return nil
}
