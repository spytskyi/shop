package repositories

import (
	"encoding/json"
	"fmt"
	"shop/models"

	"github.com/go-pg/pg/v10"
)

func RepositoryInsertProduct(db *pg.DB, reqData *models.ProductCreateRequest) (err error) {

	specifications, err := json.Marshal(reqData.Specifications)
	fmt.Println(reqData.Name)
	product := models.Product{
		Name:           reqData.Name,
		VendorCode:     reqData.VendorCode,
		Price:          reqData.Price,
		StockBalance:   reqData.StockBalance,
		Description:    reqData.Description,
		Specifications: specifications,
		CategoryId:     reqData.CategoryId,
		TypeId:         reqData.TypeId,
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

func RepositoryCheckIsExistsByCode(db *pg.DB, vendorCode float64) (isExists bool, err error) {
	isExists, err = db.Model(&models.Product{}).Where("vendor_code = ?", vendorCode).Exists()
	if err != nil {
		return isExists, err
	}
	return isExists, err
}
