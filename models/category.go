package models

type Category struct {
	tableName struct{} `pg:"categories"`
	Id        int
	Name      string
}

type CategoryRequest struct {
	Name string `json:"name"`
}
