package models

//Category struct ...
type Category struct {
	tableName struct{} `pg:"categories"`
	Id        int      `pg:",pk"`
	Name      string
}

//CategoryRequest struct
type CategoryRequest struct {
	Name string `json:"name"`
}
