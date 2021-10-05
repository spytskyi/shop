package repositories

import (
	"shop/models"

	"github.com/go-pg/pg/v10"
)

func CreateType(db *pg.DB, reqData *models.TypeRequest) (err error) {

	newType := models.Type{
		Name:       reqData.Name,
		CategoryId: reqData.CategoryId,
	}

	_, err = db.Model(&newType).Insert()
	return err
}
