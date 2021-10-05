package models

type Type struct {
	tableName  struct{} `pg:"types"`
	Id         int      `pg:",pk"`
	Name       string
	CategoryId int `pg:"on_delete:RESTRICT"`
}

type TypeRequest struct {
	Name       string `json:"name"`
	CategoryId int    `json:"category_id"`
}
