package models

import "encoding/json"

//type Product struct {
//	tableName      struct{} `pg:"products"`
//	Id             int64    `pg:"id,pk"`
//	VendorCode     string   `pg:"vendor_code"`
//	Name           string   `pg:"name"`
//	Category       string   `pg:"category"`
//	Type           string
//	Description    string   `pg:"description"`
//	StockBalance   int32    `pg:"stock_balance"`
//	Specifications json.RawMessage
//}

type Product struct {
	tableName      struct{} `pg:"products"`
	Id             int      `pg:"id,pk"`
	VendorCode     string
	Name           string
	Category       *Category
	Type           *Type
	Description    string
	StockBalance   int
	Specifications json.RawMessage
}

type ProductCreateRequest struct {
	VendorCode     string `json:"vendor_code"`
	Name           string `json:"name_product"`
	Category       *Category
	Type           *Type
	StockBalance   int `json:"stock_balance"`
	Description    string
	Specifications json.RawMessage
	//Specifications *Specifications
}
