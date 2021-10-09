package repositories

import (
	"shop/models"
	"strings"

	"github.com/go-pg/pg/v10"
)

func RepositoryInsertType(db *pg.DB, reqData *models.TypeRequest) (err error) {

	newType := models.Type{
		Name:       strings.ToLower(reqData.Name),
		CategoryId: reqData.CategoryId,
	}

	_, err = db.Model(&newType).Insert()
	if err != nil {
		return err
	}
	return nil
}

func RepositorySelectTypeByName(db *pg.DB, typeName string) (typeRet models.Type, err error) {
	typeRet.Name = strings.ToLower(typeName)
	err = db.Model(&typeRet).Where("name=?", typeRet.Name).Select()
	if err != nil {
		return typeRet, err
	}
	return typeRet, nil
}

func RepositorySelectAllTypes(db *pg.DB) (arrTypes *[]models.Type, err error) {
	err = db.Model(&arrTypes).Select()
	if err != nil {
		return nil, err
	}
	return arrTypes, nil
}

func RepositorySelectTypeById(db *pg.DB, typeId int) (typeRet models.Type, err error) {
	typeRet.Id = typeId
	err = db.Model(&typeRet).WherePK().Select()
	if err != nil {
		return typeRet, err
	}
	return typeRet, nil
}

func RepositoryUpdateType(db *pg.DB, typeReq models.Type) (err error) {
	typeReq.Name = strings.ToLower(typeReq.Name)
	_, err = db.Model(&typeReq).WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}

func RepositoryDeleteType(db *pg.DB, typeReq models.Type) (err error) {
	_, err = db.Model(&typeReq).WherePK().Delete()
	if err != nil {
		return err
	}
	return nil
}

func RepositorySelectAllTypesByCategoryId(db *pg.DB, categoryId int) (arrTypes []models.Type, err error) {

	err = db.Model(&arrTypes).
		Where("category_id = ?", categoryId).
		Select()
	if err != nil {
		return nil, err
	}
	return arrTypes, nil
}

func RepositoryDeleteAllTypesByCategoryId(db *pg.DB, arrTypes []models.Type) (err error) {
	_, err = db.Model(&arrTypes).Delete()
	if err != nil {
		return err
	}
	return nil
}
