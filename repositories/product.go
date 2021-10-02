package repositories

import (
	"encoding/json"
	"fmt"
	"shop/models"

	"github.com/go-pg/pg/v10"
)

func CreateProduct(db *pg.DB, reqData *models.ProductCreateRequest) (err error) {

	specifications, err := json.Marshal(reqData.Specifications)

	if err != nil {
		fmt.Println("Error %s", err.Error())
	}
	product := models.Product{
		VendorCode:     reqData.VendorCode,
		Name:           reqData.Name,
		Category:       reqData.Category,
		Type:           reqData.Type,
		Description:    reqData.Description,
		Specifications: specifications,
	}

	if err != nil {
		return
	}

	_, err = db.Model(&product).Insert()
	if err != nil {
		return
	}

	return nil
}
